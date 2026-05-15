// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "forge-std/Script.sol";
import "@openzeppelin/contracts/proxy/ERC1967/ERC1967Proxy.sol";
import "../src/MockUSDT.sol";
import "../src/PriceOracle.sol";
import "../src/USDTPaymaster.sol";
import "../src/Simple7702Account.sol";

contract DeployScript is Script {
    function run() external {
        uint256 deployerPrivateKey = vm.envUint("DEPLOYER_PRIVATE_KEY");
        address deployer = vm.addr(deployerPrivateKey);

        vm.startBroadcast(deployerPrivateKey);

        // 1. 部署 MockUSDT
        MockUSDT usdt = new MockUSDT();
        console.log("MockUSDT:", address(usdt));

        // 2. 部署 PriceOracle
        address pancakeRouter = vm.envOr("PANCAKE_ROUTER", 0xD99d1C33F9FC3447f8E83b589D6E4F07d8C496e6);
        // BSC 测试网 WBNB 地址
        address wbnb = vm.envOr("WBNB", 0xae13d989dAC2F0DeBFF9dcA3EB5e0B1fD735F2D7);
        PriceOracle oracle = new PriceOracle(pancakeRouter, wbnb, address(usdt));
        console.log("PriceOracle:", address(oracle));

        // 3. 部署 USDTPaymaster (UUPS Proxy)
        USDTPaymaster paymasterImpl = new USDTPaymaster();
        ERC1967Proxy paymasterProxy = new ERC1967Proxy(
            address(paymasterImpl),
            abi.encodeCall(USDTPaymaster.initialize, (address(usdt), address(oracle), deployer, deployer))
        );
        USDTPaymaster paymaster = USDTPaymaster(address(paymasterProxy));
        console.log("USDTPaymaster:", address(paymaster));

        // 4. 部署 Simple7702Account (UUPS Proxy)
        Simple7702Account accountImpl = new Simple7702Account();
        ERC1967Proxy accountProxy = new ERC1967Proxy(
            address(accountImpl),
            abi.encodeCall(Simple7702Account.initialize, (address(paymaster), deployer))
        );
        Simple7702Account account = Simple7702Account(address(accountProxy));
        console.log("Simple7702Account:", address(account));

        vm.stopBroadcast();

        console.log("\n=== Deployment Summary ===");
        console.log("MockUSDT:", address(usdt));
        console.log("PriceOracle:", address(oracle));
        console.log("USDTPaymaster:", address(paymaster));
        console.log("Simple7702Account:", address(account));
        console.log("Deployer:", deployer);
    }
}
