// SPDX-License-Identifier: MIT
pragma solidity ^0.8.26;

import {Script} from "forge-std/Script.sol";
import {console2} from "forge-std/console2.sol";
import {MockUSDT} from "../src/MockUSDT.sol";
import {USDTTokenPaymaster} from "../src/USDTTokenPaymaster.sol";

contract DeployTokensAndConfigurePaymasterScript is Script {
    function run() external {
        uint256 deployerPk = vm.envUint("DEPLOYER_PRIVATE_KEY");
        address payable paymasterAddress = payable(vm.envAddress("PAYMASTER_ADDRESS"));
        address initialUser = vm.envOr("INITIAL_USER", address(0));
        uint256 initialMint = vm.envOr("INITIAL_MINT", uint256(1_000_000e6));
        uint256 tusdtPerNative = vm.envOr("TUSDT_PER_NATIVE", uint256(600e6));
        uint256 tusdcPerNative = vm.envOr("TUSDC_PER_NATIVE", uint256(600e6));

        vm.startBroadcast(deployerPk);

        MockUSDT tusdt = new MockUSDT("Test USDT", "tUSDT");
        MockUSDT tusdc = new MockUSDT("Test USDC", "tUSDC");

        USDTTokenPaymaster paymaster = USDTTokenPaymaster(paymasterAddress);
        paymaster.setTokenConfig(address(tusdt), true, tusdtPerNative);
        paymaster.setTokenConfig(address(tusdc), true, tusdcPerNative);

        if (initialUser != address(0)) {
            tusdt.mint(initialUser, initialMint);
            tusdc.mint(initialUser, initialMint);
        }

        vm.stopBroadcast();

        console2.log("Configured paymaster:", paymasterAddress);
        console2.log("tUSDT:", address(tusdt));
        console2.log("tUSDC:", address(tusdc));
    }
}
