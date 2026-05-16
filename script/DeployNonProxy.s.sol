// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "forge-std/Script.sol";
import "../src/MockUSDT.sol";
import "../src/PriceOracle.sol";
import "../src/USDTPaymasterNonProxy.sol";
import "../src/Simple7702Account.sol";

contract DeployNonProxyScript is Script {
    function run() external {
        // 使用默认的部署者私钥（Relayer 的私钥）
        uint256 deployerPrivateKey = vm.envUint("DEPLOYER_PRIVATE_KEY");
        address deployer = 0x84D98c4faa590cD7cA746E18AcF3bcE8AD61E1b2;

        vm.startBroadcast(deployerPrivateKey);

        // 1. 部署 Simple7702Account (不使用代理，直接部署)
        Simple7702Account accountImpl = new Simple7702Account();
        console.log("Simple7702Account (implementation):", address(accountImpl));

        // 2. 使用现有的 MockUSDT 和 PriceOracle 地址
        address usdtAddress = 0x0cF1130E64744860cbA5f992008527485C88F3C8;
        address oracleAddress = 0x18CC7E9CF8f40dd32Aa0fafD5FfE938B88E455a4;

        // 直接部署 Paymaster（不使用 UUPS 代理）
        USDTPaymasterNonProxy paymaster = new USDTPaymasterNonProxy(
            usdtAddress, oracleAddress, deployer, deployer
        );
        console.log("USDTPaymasterNonProxy:", address(paymaster));

        // 3. 添加 Relayer（部署者自己）
        paymaster.addRelayer(deployer);
        console.log("Added relayer:", deployer);

        vm.stopBroadcast();

        console.log("\n=== Deployment Summary ===");
        console.log("MockUSDT (existing):", usdtAddress);
        console.log("PriceOracle (existing):", oracleAddress);
        console.log("USDTPaymasterNonProxy (NEW):", address(paymaster));
        console.log("Simple7702Account (NEW impl):", address(accountImpl));
        console.log("Deployer:", deployer);
        console.log("\n=== Update configs ===");
        console.log("backend/.env: CONTRACT_PAYMASTER=", address(paymaster));
        console.log("backend/.env: CONTRACT_7702_ACCOUNT=", address(accountImpl));
        console.log("frontend/.env.local: NEXT_PUBLIC_PAYMASTER_ADDRESS=", address(paymaster));
        console.log("frontend/.env.local: NEXT_PUBLIC_7702_ACCOUNT_ADDRESS=", address(accountImpl));
    }
}