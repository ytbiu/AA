// SPDX-License-Identifier: MIT
pragma solidity ^0.8.26;

import {IEntryPoint} from "../interfaces/IEntryPoint.sol";
import {IAccount} from "../interfaces/IAccount.sol";
import {IPaymaster} from "../interfaces/IPaymaster.sol";
import {UserOperation} from "../interfaces/UserOperation.sol";

contract SimpleEntryPointMock is IEntryPoint {
    mapping(address => uint256) public deposits;

    event UserOperationHandled(address indexed sender, bool success);

    function depositTo(address account) external payable override {
        deposits[account] += msg.value;
    }

    function addStake(uint32) external payable override {
        deposits[msg.sender] += msg.value;
    }

    function withdrawTo(address payable withdrawAddress, uint256 amount) external override {
        uint256 bal = deposits[msg.sender];
        require(bal >= amount, "insufficient deposit");
        deposits[msg.sender] = bal - amount;
        (bool ok, ) = withdrawAddress.call{value: amount}("");
        require(ok, "withdraw failed");
    }

    function balanceOf(address account) external view override returns (uint256) {
        return deposits[account];
    }

    function getUserOpHash(UserOperation calldata userOp) external view override returns (bytes32) {
        return keccak256(abi.encode(block.chainid, address(this), _packUserOp(userOp)));
    }

    function handleOps(UserOperation[] calldata ops, address payable beneficiary) external override {
        beneficiary;
        for (uint256 i = 0; i < ops.length; i++) {
            UserOperation calldata op = ops[i];
            bytes32 userOpHash = keccak256(abi.encode(block.chainid, address(this), _packUserOp(op)));

            uint256 validationData = IAccount(op.sender).validateUserOp(op, userOpHash, 0);
            require(validationData == 0, "account validation failed");

            address paymaster = _parsePaymaster(op.paymasterAndData);
            bytes memory context;
            if (paymaster != address(0)) {
                uint256 maxCost = (op.callGasLimit + op.verificationGasLimit + op.preVerificationGas) * op.maxFeePerGas;
                (context, ) = IPaymaster(paymaster).validatePaymasterUserOp(op, userOpHash, maxCost);
            }

            (bool ok, ) = op.sender.call{gas: op.callGasLimit}(op.callData);

            if (paymaster != address(0)) {
                uint256 actualCost = ((op.callGasLimit / 2) + op.preVerificationGas + (op.verificationGasLimit / 2))
                    * op.maxFeePerGas;
                IPaymaster(paymaster).postOp(
                    ok ? IPaymaster.PostOpMode.opSucceeded : IPaymaster.PostOpMode.opReverted,
                    context,
                    actualCost
                );
            }

            emit UserOperationHandled(op.sender, ok);
        }
    }

    function _parsePaymaster(bytes calldata paymasterAndData) internal pure returns (address paymaster) {
        if (paymasterAndData.length < 20) {
            return address(0);
        }
        assembly {
            paymaster := shr(96, calldataload(paymasterAndData.offset))
        }
    }

    function _packUserOp(UserOperation calldata op) internal pure returns (bytes memory) {
        return abi.encode(
            op.sender,
            op.nonce,
            keccak256(op.initCode),
            keccak256(op.callData),
            op.callGasLimit,
            op.verificationGasLimit,
            op.preVerificationGas,
            op.maxFeePerGas,
            op.maxPriorityFeePerGas,
            keccak256(op.paymasterAndData)
        );
    }
}
