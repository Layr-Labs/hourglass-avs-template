// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import {Script, console} from "forge-std/Script.sol";
import {stdJson} from "forge-std/StdJson.sol";

import {IKeyRegistrarTypes} from "@eigenlayer-contracts/src/contracts/interfaces/IKeyRegistrar.sol";
import {OperatorSet, OperatorSetLib} from "@eigenlayer-contracts/src/contracts/libraries/OperatorSetLib.sol";
import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import {ITaskMailbox, ITaskMailboxTypes} from "@eigenlayer-contracts/src/contracts/interfaces/ITaskMailbox.sol";
import {IAVSTaskHook} from "@eigenlayer-contracts/src/contracts/interfaces/IAVSTaskHook.sol";

contract SetupAVSTaskMailboxConfig is Script {
    using stdJson for string;

    function run(string memory environment, uint32 executorOperatorSetId, uint96 taskSLA, uint8 curveType) public {
        // Read addresses from config files
        address taskMailbox = _readHourglassConfigAddress(environment, "taskMailbox");
        console.log("Task Mailbox:", taskMailbox);

        // Read AVS L2 contract addresses
        address taskHook = _readAVSL2ConfigAddress(environment, "avsTaskHook");
        console.log("AVS Task Hook:", taskHook);

        // Load the private key from the environment variable
        uint256 avsPrivateKey = vm.envUint("PRIVATE_KEY_AVS");
        address avs = vm.addr(avsPrivateKey);

        vm.startBroadcast(avsPrivateKey);
        console.log("AVS address:", avs);

        // Set the Executor Operator Set Task Config
        ITaskMailboxTypes.ExecutorOperatorSetTaskConfig memory executorOperatorSetTaskConfig = ITaskMailboxTypes
            .ExecutorOperatorSetTaskConfig({
            taskHook: IAVSTaskHook(taskHook),
            taskSLA: taskSLA,
            feeToken: IERC20(address(0)),
            curveType: IKeyRegistrarTypes.CurveType(curveType),
            feeCollector: address(0),
            consensus: ITaskMailboxTypes.Consensus({
                consensusType: ITaskMailboxTypes.ConsensusType.STAKE_PROPORTION_THRESHOLD,
                value: abi.encode(10_000)
            }),
            taskMetadata: bytes("")
        });
        ITaskMailbox(taskMailbox).setExecutorOperatorSetTaskConfig(
            OperatorSet(avs, executorOperatorSetId), executorOperatorSetTaskConfig
        );
        ITaskMailboxTypes.ExecutorOperatorSetTaskConfig memory executorOperatorSetTaskConfigStored =
            ITaskMailbox(taskMailbox).getExecutorOperatorSetTaskConfig(OperatorSet(avs, executorOperatorSetId));
        console.log(
            "Executor Operator Set Task Config set with curve type:",
            uint8(executorOperatorSetTaskConfigStored.curveType),
            address(executorOperatorSetTaskConfigStored.taskHook)
        );

        vm.stopBroadcast();
    }

    function _readHourglassConfigAddress(
        string memory environment,
        string memory key
    ) internal view returns (address) {
        // Load the Hourglass config file
        string memory hourglassConfigFile =
            string.concat("script/", environment, "/output/deploy_hourglass_core_output.json");
        string memory hourglassConfig = vm.readFile(hourglassConfigFile);

        // Parse and return the address
        return stdJson.readAddress(hourglassConfig, string.concat(".addresses.", key));
    }

    function _readAVSL2ConfigAddress(string memory environment, string memory key) internal view returns (address) {
        // Load the AVS L2 config file
        string memory avsL2ConfigFile = string.concat("script/", environment, "/output/deploy_avs_l2_output.json");
        string memory avsL2Config = vm.readFile(avsL2ConfigFile);

        // Parse and return the address
        return stdJson.readAddress(avsL2Config, string.concat(".addresses.", key));
    }
}
