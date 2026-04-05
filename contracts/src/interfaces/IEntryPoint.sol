// SPDX-License-Identifier: MIT
pragma solidity ^0.8.26;

import {UserOperation} from "./UserOperation.sol";

interface IEntryPoint {
    function depositTo(address account) external payable;
    function addStake(uint32 unstakeDelaySec) external payable;

    function withdrawTo(address payable withdrawAddress, uint256 amount) external;

    function balanceOf(address account) external view returns (uint256);

    function getUserOpHash(UserOperation calldata userOp) external view returns (bytes32);

    function handleOps(UserOperation[] calldata ops, address payable beneficiary) external;
}
