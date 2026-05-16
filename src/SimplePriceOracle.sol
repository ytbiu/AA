// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "./interfaces/IPriceOracle.sol";

contract SimplePriceOracle is IPriceOracle {
    // 固定价格: 1 BNB = 600 USDT (示例价格)
    // 返回值需要乘以 1e18 以匹配精度
    uint256 private constant BNB_PRICE_USDT = 600 * 1e18;

    address private _router;

    function getBNBPriceInUSDT() external pure override returns (uint256) {
        return BNB_PRICE_USDT;
    }

    function setRouter(address _routerAddr) external override {
        _router = _routerAddr;
    }

    function router() external view override returns (address) {
        return _router;
    }
}