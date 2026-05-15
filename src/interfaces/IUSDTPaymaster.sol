// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

interface IUSDTPaymaster {
    struct UserOperation {
        address user;
        Call[] calls;
    }

    struct Call {
        address to;
        bytes data;
    }

    function executeBatch(UserOperation calldata userOp, bytes calldata signature) external;
    function addRelayer(address relayer) external;
    function removeRelayer(address relayer) external;
    function setFeeRate(uint256 rate) external;
    function setFeeRecipient(address recipient) external;
    function setOracle(address oracle) external;
    function isRelayer(address relayer) external view returns (bool);
    function feeRate() external view returns (uint256);
    function feeRecipient() external view returns (address);
    function usdtToken() external view returns (address);
    function oracle() external view returns (address);
}
