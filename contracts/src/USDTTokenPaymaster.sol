// SPDX-License-Identifier: MIT
pragma solidity ^0.8.26;

import {ECDSA} from "@openzeppelin/contracts/utils/cryptography/ECDSA.sol";
import {MessageHashUtils} from "@openzeppelin/contracts/utils/cryptography/MessageHashUtils.sol";
import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import {Ownable} from "@openzeppelin/contracts/access/Ownable.sol";
import {IEntryPoint} from "./interfaces/IEntryPoint.sol";
import {IPaymaster} from "./interfaces/IPaymaster.sol";
import {UserOperation} from "./interfaces/UserOperation.sol";

/// @notice paymasterAndData 编码:
/// abi.encodePacked(
///   address(paymaster),
///   address(gasToken),
///   uint48(validUntil),
///   uint256(tokenAmount),
///   bytes(paymasterQuoteSignature65)
/// )
contract USDTTokenPaymaster is IPaymaster, Ownable {
    uint256 public constant PAYMASTER_DATA_LENGTH = 20 + 20 + 6 + 32 + 65;

    struct TokenConfig {
        bool enabled;
        uint256 tokenPerNative; // 1e18 native 对应 token 最小单位
    }

    IEntryPoint public immutable entryPoint;
    mapping(address => TokenConfig) public tokenConfigs;

    address public quoteSigner;
    uint256 public settlementMarkupBps; // 例如 500 => 5%

    error NotFromEntryPoint();
    error InvalidPaymasterDataLength(uint256 actual);
    error UnsupportedGasToken(address token);
    error InvalidTokenPerNative();
    error UnderQuotedTokenAmount(uint256 provided, uint256 minimumRequired);
    error InvalidPaymasterQuoteSignature();
    error TokenTransferFailed();

    event QuoteSignerUpdated(address indexed oldSigner, address indexed newSigner);
    event TokenConfigUpdated(address indexed token, bool enabled, uint256 tokenPerNative);
    event MarkupUpdated(uint256 oldMarkup, uint256 newMarkup);
    event Settled(address indexed user, address indexed token, uint256 tokenCharged, uint256 actualGasCost);

    constructor(
        IEntryPoint _entryPoint,
        address _quoteSigner,
        uint256 _settlementMarkupBps
    ) Ownable(msg.sender) {
        entryPoint = _entryPoint;
        quoteSigner = _quoteSigner;
        settlementMarkupBps = _settlementMarkupBps;
    }

    modifier onlyEntryPoint() {
        if (msg.sender != address(entryPoint)) {
            revert NotFromEntryPoint();
        }
        _;
    }

    receive() external payable {}

    function setQuoteSigner(address newSigner) external onlyOwner {
        emit QuoteSignerUpdated(quoteSigner, newSigner);
        quoteSigner = newSigner;
    }

    function setTokenConfig(address token, bool enabled, uint256 tokenPerNative) external onlyOwner {
        if (enabled && tokenPerNative == 0) {
            revert InvalidTokenPerNative();
        }
        tokenConfigs[token] = TokenConfig({enabled: enabled, tokenPerNative: tokenPerNative});
        emit TokenConfigUpdated(token, enabled, tokenPerNative);
    }

    function setSettlementMarkupBps(uint256 newMarkup) external onlyOwner {
        require(newMarkup <= 2_000, "markup too high");
        emit MarkupUpdated(settlementMarkupBps, newMarkup);
        settlementMarkupBps = newMarkup;
    }

    function depositToEntryPoint() external payable onlyOwner {
        entryPoint.depositTo{value: msg.value}(address(this));
    }

    function addStakeToEntryPoint(uint32 unstakeDelaySec) external payable onlyOwner {
        entryPoint.addStake{value: msg.value}(unstakeDelaySec);
    }

    function withdrawDeposit(address payable to, uint256 amount) external onlyOwner {
        entryPoint.withdrawTo(to, amount);
    }

    function withdrawToken(address token, address to, uint256 amount) external onlyOwner {
        bool ok = IERC20(token).transfer(to, amount);
        if (!ok) {
            revert TokenTransferFailed();
        }
    }

    function validatePaymasterUserOp(
        UserOperation calldata userOp,
        bytes32,
        uint256 maxCost
    ) external override onlyEntryPoint returns (bytes memory context, uint256 validationData) {
        if (userOp.paymasterAndData.length != PAYMASTER_DATA_LENGTH) {
            revert InvalidPaymasterDataLength(userOp.paymasterAndData.length);
        }

        bytes calldata paymasterData = userOp.paymasterAndData[20:];

        address gasToken = _readAddress(paymasterData, 0);
        TokenConfig memory cfg = tokenConfigs[gasToken];
        if (!cfg.enabled) {
            revert UnsupportedGasToken(gasToken);
        }

        uint48 validUntil = _readUint48(paymasterData, 20);
        // 禁止在 validatePaymasterUserOp 中读取 block.timestamp（部分 bundler 会拒绝 TIMESTAMP opcode）。
        // 过期约束交给 validationData.validUntil 由 EntryPoint 统一处理。

        uint256 tokenAmount = _readUint256(paymasterData, 26);
        bytes calldata quoteSig = paymasterData[58:123];

        uint256 minimumRequired = _nativeToToken(maxCost, cfg.tokenPerNative);
        if (tokenAmount < minimumRequired) {
            revert UnderQuotedTokenAmount(tokenAmount, minimumRequired);
        }

        bytes32 quoteHash = _quoteHash(userOp, gasToken, tokenAmount, validUntil);
        bytes memory quoteSigMem = quoteSig;
        address recoveredSigner = ECDSA.recover(
            MessageHashUtils.toEthSignedMessageHash(quoteHash),
            quoteSigMem
        );
        if (recoveredSigner != quoteSigner) {
            revert InvalidPaymasterQuoteSignature();
        }

        context = abi.encode(userOp.sender, gasToken, tokenAmount);
        validationData = _packValidationData(validUntil, 0);
    }

    function postOp(PostOpMode, bytes calldata context, uint256 actualGasCost) external override onlyEntryPoint {
        (address user, address gasToken, uint256 maxTokenCharge) = abi.decode(context, (address, address, uint256));
        TokenConfig memory cfg = tokenConfigs[gasToken];
        if (!cfg.enabled) {
            revert UnsupportedGasToken(gasToken);
        }

        uint256 settlement = _nativeToToken(actualGasCost, cfg.tokenPerNative);
        settlement = (settlement * (10_000 + settlementMarkupBps)) / 10_000;
        if (settlement > maxTokenCharge) {
            settlement = maxTokenCharge;
        }

        bool ok = IERC20(gasToken).transferFrom(user, address(this), settlement);
        if (!ok) {
            revert TokenTransferFailed();
        }

        emit Settled(user, gasToken, settlement, actualGasCost);
    }

    function _readAddress(bytes calldata data, uint256 offset) internal pure returns (address value) {
        assembly {
            value := shr(96, calldataload(add(data.offset, offset)))
        }
    }

    function _readUint48(bytes calldata data, uint256 offset) internal pure returns (uint48 value) {
        assembly {
            value := shr(208, calldataload(add(data.offset, offset)))
        }
    }

    function _readUint256(bytes calldata data, uint256 offset) internal pure returns (uint256 value) {
        assembly {
            value := calldataload(add(data.offset, offset))
        }
    }

    function _nativeToToken(uint256 nativeWeiAmount, uint256 tokenPerNative) internal pure returns (uint256) {
        return (nativeWeiAmount * tokenPerNative) / 1e18;
    }

    function _quoteHash(
        UserOperation calldata userOp,
        address gasToken,
        uint256 tokenAmount,
        uint48 validUntil
    ) internal view returns (bytes32) {
        return keccak256(
            abi.encode(
                block.chainid,
                address(this),
                userOp.sender,
                userOp.nonce,
                keccak256(userOp.callData),
                userOp.callGasLimit,
                userOp.verificationGasLimit,
                userOp.preVerificationGas,
                userOp.maxFeePerGas,
                userOp.maxPriorityFeePerGas,
                gasToken,
                tokenAmount,
                validUntil
            )
        );
    }

    function _packValidationData(uint48 validUntil, uint48 validAfter) internal pure returns (uint256) {
        return (uint256(validUntil) << 160) | (uint256(validAfter) << 208);
    }
}
