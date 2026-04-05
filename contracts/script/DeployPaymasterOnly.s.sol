// SPDX-License-Identifier: MIT
pragma solidity ^0.8.26;

import {Script} from "forge-std/Script.sol";
import {console2} from "forge-std/console2.sol";
import {IEntryPoint} from "../src/interfaces/IEntryPoint.sol";
import {USDTTokenPaymaster} from "../src/USDTTokenPaymaster.sol";

contract DeployPaymasterOnlyScript is Script {
    function run() external {
        uint256 deployerPk = vm.envUint("DEPLOYER_PRIVATE_KEY");
        address entryPointAddress = vm.envAddress("ENTRYPOINT_ADDRESS");
        address tusdtAddress = vm.envAddress("TUSDT_ADDRESS");
        address tusdcAddress = vm.envAddress("TUSDC_ADDRESS");
        address quoteSigner = vm.envAddress("QUOTE_SIGNER");

        uint256 tusdtPerNative = vm.envOr("TUSDT_PER_NATIVE", uint256(600e6));
        uint256 tusdcPerNative = vm.envOr("TUSDC_PER_NATIVE", uint256(600e6));
        uint256 settlementMarkupBps = vm.envOr("SETTLEMENT_MARKUP_BPS", uint256(500));
        uint256 paymasterDepositWei = vm.envOr("PAYMASTER_DEPOSIT_WEI", uint256(0.2 ether));
        uint256 paymasterStakeWei = vm.envOr("PAYMASTER_STAKE_WEI", uint256(0.1 ether));
        uint32 unstakeDelaySec = uint32(vm.envOr("PAYMASTER_UNSTAKE_DELAY_SEC", uint256(86400)));

        vm.startBroadcast(deployerPk);

        USDTTokenPaymaster paymaster = new USDTTokenPaymaster(
            IEntryPoint(entryPointAddress),
            quoteSigner,
            settlementMarkupBps
        );
        paymaster.setTokenConfig(tusdtAddress, true, tusdtPerNative);
        paymaster.setTokenConfig(tusdcAddress, true, tusdcPerNative);

        if (paymasterDepositWei > 0) {
            paymaster.depositToEntryPoint{value: paymasterDepositWei}();
        }

        if (paymasterStakeWei > 0) {
            paymaster.addStakeToEntryPoint{value: paymasterStakeWei}(unstakeDelaySec);
        }

        vm.stopBroadcast();

        console2.log("USDTTokenPaymaster(new):", address(paymaster));
        console2.log("EntryPoint:", entryPointAddress);
        console2.log("tUSDT:", tusdtAddress);
        console2.log("tUSDC:", tusdcAddress);
        console2.log("quoteSigner:", quoteSigner);
    }
}

