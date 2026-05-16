// SPDX-License-Identifier: MIT
pragma solidity ^0.8.24;

import "@openzeppelin/contracts/utils/ReentrancyGuard.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts/utils/cryptography/ECDSA.sol";
import "@openzeppelin/contracts/utils/cryptography/MessageHashUtils.sol";
import "../interfaces/IUSDTPaymaster.sol";
import "../interfaces/IPriceOracle.sol";

abstract contract PaymasterBase is ReentrancyGuard, IUSDTPaymaster {
    using ECDSA for bytes32;
    using MessageHashUtils for bytes32;

    IERC20 internal _usdtToken;
    IPriceOracle internal _oracle;
    uint256 public feeRate;
    address public feeRecipient;
    mapping(address => bool) internal _relayers;

    event BatchExecuted(address indexed user, uint256 gasUsed, uint256 compensation);
    event RelayerAdded(address indexed relayer);
    event RelayerRemoved(address indexed relayer);
    event FeeRateUpdated(uint256 rate);
    event FeeRecipientUpdated(address recipient);
    event OracleUpdated(address oracle);
    event OperationExecuted(
        address indexed user,
        address indexed relayer,
        bytes32 operationHash,
        uint256 gasUsed,
        uint256 compensation,
        uint256 fee,
        bytes callsData
    );

    error NotRelayer();
    error InvalidSignature();
    error CallFailed();
    error TransferFailed();

    function _checkOwner() internal virtual;

    modifier onlyRelayer() {
        if (!_relayers[msg.sender]) revert NotRelayer();
        _;
    }

    modifier onlyAuthorized() {
        _checkOwner();
        _;
    }

    function executeBatch(UserOperation calldata userOp, bytes calldata signature)
        external
        onlyRelayer
        nonReentrant
    {
        bytes32 hash = keccak256(abi.encode(userOp));
        bytes32 ethSignedHash = hash.toEthSignedMessageHash();
        address signer = ethSignedHash.recover(signature);
        if (signer != userOp.user) revert InvalidSignature();

        uint256 gasBefore = gasleft();
        bytes memory executeData = abi.encodeWithSignature(
            "executeBatch((address,bytes)[])",
            userOp.calls
        );
        (bool success, bytes memory result) = userOp.user.call(executeData);
        if (!success) {
            if (result.length > 0) {
                assembly {
                    revert(add(result, 32), mload(result))
                }
            }
            revert CallFailed();
        }
        uint256 gasUsed = gasBefore - gasleft();

        uint256 bnbCost = gasUsed * tx.gasprice;
        uint256 bnbPriceInUsdt = _oracle.getBNBPriceInUSDT();
        uint256 compensation = (bnbCost * bnbPriceInUsdt) / 1e18;

        uint256 fee = (compensation * feeRate) / 10000;

        if (!_usdtToken.transferFrom(userOp.user, msg.sender, compensation))
            revert TransferFailed();
        if (fee > 0) {
            if (!_usdtToken.transferFrom(userOp.user, feeRecipient, fee))
                revert TransferFailed();
        }

        emit BatchExecuted(userOp.user, gasUsed, compensation + fee);

        bytes memory callsData = abi.encode(userOp.calls);
        emit OperationExecuted(
            userOp.user,
            msg.sender,
            keccak256(callsData),
            gasUsed,
            compensation,
            fee,
            callsData
        );
    }

    function addRelayer(address relayer) external onlyAuthorized {
        _relayers[relayer] = true;
        emit RelayerAdded(relayer);
    }

    function removeRelayer(address relayer) external onlyAuthorized {
        _relayers[relayer] = false;
        emit RelayerRemoved(relayer);
    }

    function setFeeRate(uint256 rate) external onlyAuthorized {
        feeRate = rate;
        emit FeeRateUpdated(rate);
    }

    function setFeeRecipient(address recipient) external onlyAuthorized {
        feeRecipient = recipient;
        emit FeeRecipientUpdated(recipient);
    }

    function setOracle(address _oracleAddr) external onlyAuthorized {
        _oracle = IPriceOracle(_oracleAddr);
        emit OracleUpdated(_oracleAddr);
    }

    function isRelayer(address relayer) external view returns (bool) {
        return _relayers[relayer];
    }

    function usdtToken() external view returns (address) {
        return address(_usdtToken);
    }

    function oracle() external view returns (address) {
        return address(_oracle);
    }
}