// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "@openzeppelin/contracts/access/Ownable.sol";
import "./interfaces/IPriceOracle.sol";

interface IPancakeRouter {
    function getAmountsOut(uint256 amountIn, address[] calldata path) external view returns (uint256[] memory amounts);
    function WBNB() external view returns (address);
}

contract PriceOracle is Ownable, IPriceOracle {
    address public router;
    address public immutable WBNB;
    address public immutable USDT;

    constructor(address _router, address _usdt) Ownable(msg.sender) {
        router = _router;
        WBNB = IPancakeRouter(_router).WBNB();
        USDT = _usdt;
    }

    function getBNBPriceInUSDT() external view returns (uint256) {
        // 1 BNB 的 USDT 价格
        address[] memory path = new address[](2);
        path[0] = WBNB;
        path[1] = USDT;

        uint256[] memory amounts = IPancakeRouter(router).getAmountsOut(1e18, path);
        return amounts[1];
    }

    function setRouter(address _router) external onlyOwner {
        router = _router;
    }
}
