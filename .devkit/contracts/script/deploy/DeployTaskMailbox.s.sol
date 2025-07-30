// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import {Script, console} from "forge-std/Script.sol";
import {ProxyAdmin} from "@openzeppelin/contracts/proxy/transparent/ProxyAdmin.sol";
import {TransparentUpgradeableProxy} from "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";
import {ITransparentUpgradeableProxy} from "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";

import {IKeyRegistrarTypes} from "@eigenlayer-contracts/src/contracts/interfaces/IKeyRegistrar.sol";
import {TaskMailbox} from "@eigenlayer-contracts/src/contracts/avs/task/TaskMailbox.sol";

contract DeployTaskMailbox is Script {
    function run(
        string memory environment,
        address bn254CertificateVerifier,
        address ecdsaCertificateVerifier
    ) public {
        // Load the private key from the environment variable
        uint256 deployerPrivateKey = vm.envUint("PRIVATE_KEY_DEPLOYER");
        address deployer = vm.addr(deployerPrivateKey);

        // Deploy the TaskMailbox contract with proxy pattern
        vm.startBroadcast(deployerPrivateKey);
        console.log("Deployer address:", deployer);

        // Deploy ProxyAdmin
        ProxyAdmin proxyAdmin = new ProxyAdmin();
        console.log("ProxyAdmin deployed to:", address(proxyAdmin));

        // Deploy implementation
        TaskMailbox taskMailboxImpl = new TaskMailbox(bn254CertificateVerifier, ecdsaCertificateVerifier, "0.0.1");
        console.log("TaskMailbox implementation deployed to:", address(taskMailboxImpl));

        // Deploy proxy with initialization
        TransparentUpgradeableProxy proxy = new TransparentUpgradeableProxy(
            address(taskMailboxImpl),
            address(proxyAdmin),
            abi.encodeWithSelector(TaskMailbox.initialize.selector, deployer)
        );
        console.log("TaskMailbox proxy deployed to:", address(proxy));

        // Transfer ProxyAdmin ownership to deployer (or a multisig in production)
        proxyAdmin.transferOwnership(deployer);

        vm.stopBroadcast();

        // Write deployment info to output file
        _writeOutputToJson(environment, address(proxy), address(taskMailboxImpl), address(proxyAdmin));
    }

    function _writeOutputToJson(string memory environment, address taskMailboxProxy, address taskMailboxImpl, address proxyAdmin) internal {
        // Add the addresses object
        string memory addresses = "addresses";
        vm.serializeAddress(addresses, "taskMailbox", taskMailboxProxy);
        vm.serializeAddress(addresses, "taskMailboxImpl", taskMailboxImpl);
        addresses = vm.serializeAddress(addresses, "l2ProxyAdmin", proxyAdmin);

        // Add the chainInfo object
        string memory chainInfo = "chainInfo";
        chainInfo = vm.serializeUint(chainInfo, "chainId", block.chainid);

        // Finalize the JSON
        string memory finalJson = "final";
        vm.serializeString(finalJson, "addresses", addresses);
        finalJson = vm.serializeString(finalJson, "chainInfo", chainInfo);

        // Write to output file
        string memory outputFile = string.concat("script/", environment, "/output/deploy_hourglass_core_output.json");
        vm.writeJson(finalJson, outputFile);
    }
}
