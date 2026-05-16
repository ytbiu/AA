// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "forge-std/Script.sol";
import "../src/Simple7702Account.sol";

contract Deploy7702 is Script {
    function run() external {
        uint256 deployerPrivateKey = vm.envUint("DEPLOYER_PRIVATE_KEY");

        vm.startBroadcast(deployerPrivateKey);

        // 只部署 Simple7702Account implementation（不通过 proxy）
        Simple7702Account accountImpl = new Simple7702Account();
        console.log("Simple7702Account Implementation:", address(accountImpl));

        vm.stopBroadcast();
    }
}