// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import {Script, console} from "forge-std/Script.sol";

import {AVSTaskHook} from "@project/l2-contracts/AVSTaskHook.sol";

contract DeployAVSL2Contracts is Script {
    function run(
        string memory environment
    ) public {
        // Load the private key from the environment variable
        uint256 deployerPrivateKey = vm.envUint("PRIVATE_KEY_DEPLOYER");
        address deployer = vm.addr(deployerPrivateKey);

        // Deploy the AVSTaskHook contract
        vm.startBroadcast(deployerPrivateKey);
        console.log("Deployer address:", deployer);

        AVSTaskHook avsTaskHook = new AVSTaskHook();
        console.log("AVSTaskHook deployed to:", address(avsTaskHook));

        vm.stopBroadcast();

        // Write deployment info to output file
        _writeOutputToJson(environment, address(avsTaskHook));
    }

    function _writeOutputToJson(string memory environment, address avsTaskHook) internal {
        // Add the addresses object
        string memory addresses = "addresses";
        addresses = vm.serializeAddress(addresses, "avsTaskHook", avsTaskHook);

        // Add the chainInfo object
        string memory chainInfo = "chainInfo";
        chainInfo = vm.serializeUint(chainInfo, "chainId", block.chainid);

        // Finalize the JSON
        string memory finalJson = "final";
        vm.serializeString(finalJson, "addresses", addresses);
        finalJson = vm.serializeString(finalJson, "chainInfo", chainInfo);

        // Write to output file
        string memory outputFile = string.concat("script/", environment, "/output/deploy_avs_l2_output.json");
        vm.writeJson(finalJson, outputFile);
    }
}
