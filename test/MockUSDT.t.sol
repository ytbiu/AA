// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "forge-std/Test.sol";
import "../src/MockUSDT.sol";

contract MockUSDTTest is Test {
    MockUSDT usdt;
    address deployer;
    address user;

    function setUp() public {
        deployer = address(this);
        user = address(0x1);
        usdt = new MockUSDT();
    }

    function test_NameAndSymbol() public view {
        assertEq(usdt.name(), "Mock USDT");
        assertEq(usdt.symbol(), "USDT");
    }

    function test_Decimals() public view {
        assertEq(usdt.decimals(), 18);
    }

    function test_InitialSupply() public view {
        assertEq(usdt.balanceOf(deployer), 1000000 * 10 ** 18);
    }

    function test_Faucet() public {
        vm.prank(user);
        usdt.faucet();
        assertEq(usdt.balanceOf(user), 100 * 10 ** 18);
    }

    function test_FaucetMultipleTimes() public {
        vm.prank(user);
        usdt.faucet();
        vm.prank(user);
        usdt.faucet();
        assertEq(usdt.balanceOf(user), 200 * 10 ** 18);
    }

    function test_FaucetAmountConstant() public view {
        assertEq(usdt.FAUCET_AMOUNT(), 100 * 10 ** 18);
    }
}
