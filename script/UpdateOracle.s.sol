// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "forge-std/Script.sol";
import "../src/SimplePriceOracle.sol";
import "../src/USDTPaymasterNonProxy.sol";

contract UpdateOracleScript is Script {
    function run() external {
        uint256 deployerPrivateKey = vm.envUint("DEPLOYER_PRIVATE_KEY");
        address deployer = 0x84D98c4faa590cD7cA746E18AcF3bcE8AD61E1b2;
        address paymasterAddress = 0xA61D461AF55029B58d4846C9EA818De9cBC711D3;

        vm.startBroadcast(deployerPrivateKey);

        // 1. 部署 SimplePriceOracle
        SimplePriceOracle oracle = new SimplePriceOracle();
        console.log("SimplePriceOracle:", address(oracle));

        // 2. 更新 Paymaster 的 Oracle
        USDTPaymasterNonProxy paymaster = USDTPaymasterNonProxy(paymasterAddress);
        paymaster.setOracle(address(oracle));
        console.log("Updated Paymaster oracle to:", address(oracle));

        vm.stopBroadcast();

        console.log("\n=== Update Summary ===");
        console.log("SimplePriceOracle (NEW):", address(oracle));
        console.log("USDTPaymasterNonProxy:", paymasterAddress);
        console.log("\n=== No config update needed ===");
    }
}