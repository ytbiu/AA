// SPDX-License-Identifier: MIT
pragma solidity ^0.8.26;

import {ECDSA} from "@openzeppelin/contracts/utils/cryptography/ECDSA.sol";
import {MessageHashUtils} from "@openzeppelin/contracts/utils/cryptography/MessageHashUtils.sol";
import {IAccount} from "./interfaces/IAccount.sol";
import {IEntryPoint} from "./interfaces/IEntryPoint.sol";
import {UserOperation} from "./interfaces/UserOperation.sol";

/// @notice 作为 EIP-7702 delegation target 的账户逻辑。
/// 在 7702 升级后，EOA 会委托执行到本合约，签名者仍然是原 EOA 本身。
contract EIP7702DelegateAccount is IAccount {
    uint256 internal constant SIG_VALIDATION_FAILED = 1;

    IEntryPoint public immutable entryPoint;
    uint256 private _nonce;

    error NotFromEntryPoint();
    error NotFromEntryPointOrSelf();
    error CallFailed(bytes reason);

    event Executed(address indexed target, uint256 value, bytes data);

    constructor(IEntryPoint _entryPoint) {
        entryPoint = _entryPoint;
    }

    receive() external payable {}

    modifier onlyEntryPointOrSelf() {
        if (msg.sender != address(entryPoint) && msg.sender != address(this)) {
            revert NotFromEntryPointOrSelf();
        }
        _;
    }

    function nonce() external view returns (uint256) {
        return _nonce;
    }

    function execute(address target, uint256 value, bytes calldata data) external onlyEntryPointOrSelf {
        (bool ok, bytes memory ret) = target.call{value: value}(data);
        if (!ok) {
            revert CallFailed(ret);
        }
        emit Executed(target, value, data);
    }

    function executeBatch(
        address[] calldata targets,
        uint256[] calldata values,
        bytes[] calldata data
    ) external onlyEntryPointOrSelf {
        uint256 length = targets.length;
        require(length == values.length && length == data.length, "length mismatch");
        for (uint256 i = 0; i < length; i++) {
            (bool ok, bytes memory ret) = targets[i].call{value: values[i]}(data[i]);
            if (!ok) {
                revert CallFailed(ret);
            }
            emit Executed(targets[i], values[i], data[i]);
        }
    }

    function validateUserOp(
        UserOperation calldata userOp,
        bytes32 userOpHash,
        uint256 missingAccountFunds
    ) external override returns (uint256 validationData) {
        if (msg.sender != address(entryPoint)) {
            revert NotFromEntryPoint();
        }

        if (userOp.nonce != _nonce) {
            return SIG_VALIDATION_FAILED;
        }

        address recovered = ECDSA.recover(
            MessageHashUtils.toEthSignedMessageHash(userOpHash),
            userOp.signature
        );

        // 委托执行场景下 address(this) 就是 EOA 地址。
        if (recovered != address(this)) {
            return SIG_VALIDATION_FAILED;
        }

        _nonce++;

        if (missingAccountFunds > 0) {
            (bool sent, ) = payable(msg.sender).call{value: missingAccountFunds}("");
            sent;
        }

        return 0;
    }
}
