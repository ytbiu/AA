// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

interface IPriceOracle {
    function getBNBPriceInUSDT() external view returns (uint256);
    function setRouter(address _router) external;
    function router() external view returns (address);
}
