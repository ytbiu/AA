// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "forge-std/Test.sol";
import "../src/MockUSDT.sol";
import "../src/PriceOracle.sol";

contract MockPancakeRouter {
    address public WBNB;
    address public USDT;

    constructor(address _wbnb, address _usdt) {
        WBNB = _wbnb;
        USDT = _usdt;
    }

    function getAmountsOut(uint256 amountIn, address[] calldata path) external view returns (uint256[] memory amounts) {
        amounts = new uint256[](path.length);
        // 模拟 1 BNB = 300 USDT (USDT 也是 18 decimals)
        amounts[0] = amountIn;
        if (path.length == 2 && path[1] == USDT) {
            amounts[1] = amountIn * 300; // 300 USDT per BNB
        } else {
            amounts[1] = amountIn;
        }
    }
}

contract PriceOracleTest is Test {
    PriceOracle oracle;
    MockUSDT usdt;
    MockPancakeRouter pancakeRouter;
    address deployer;
    address user;

    function setUp() public {
        deployer = address(this);
        user = address(0x1);

        usdt = new MockUSDT();
        // BSC 测试网 WBNB 地址
        address wbnb = 0xae13d989dAC2F0DeBFF9dcA3EB5e0B1fD735F2D7;
        pancakeRouter = new MockPancakeRouter(wbnb, address(usdt));
        oracle = new PriceOracle(address(pancakeRouter), wbnb, address(usdt));
    }

    function test_RouterSetCorrectly() public view {
        assertEq(oracle.router(), address(pancakeRouter));
    }

    function test_OwnerIsDeployer() public view {
        assertEq(oracle.owner(), deployer);
    }

    function test_OwnerCanUpdateRouter() public {
        oracle.setRouter(user);
        assertEq(oracle.router(), user);
    }

    function test_NonOwnerCannotUpdateRouter() public {
        vm.prank(user);
        vm.expectRevert();
        oracle.setRouter(user);
    }

    function test_USDTAddressSetCorrectly() public view {
        assertEq(oracle.USDT(), address(usdt));
    }

    function test_WBNBAddressSetCorrectly() public view {
        assertEq(oracle.WBNB(), 0xae13d989dAC2F0DeBFF9dcA3EB5e0B1fD735F2D7);
    }

    function test_GetBNBPriceInUSDT() public view {
        // 模拟 1 BNB = 300 USDT
        uint256 price = oracle.getBNBPriceInUSDT();
        assertEq(price, 300e18);
    }
}
