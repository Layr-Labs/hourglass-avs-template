// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import {Script, console} from "forge-std/Script.sol";
import {stdJson} from "forge-std/StdJson.sol";
import {OperatorSet} from "@eigenlayer-contracts/src/contracts/libraries/OperatorSetLib.sol";
import {ITaskMailbox, ITaskMailboxTypes} from "@eigenlayer-contracts/src/contracts/interfaces/ITaskMailbox.sol";

contract CreateTask is Script {
    using stdJson for string;

    function run(
        string memory environment,
        address taskMailbox,
        address avs,
        uint32 executorOperatorSetId,
        bytes memory payload
    ) public {
        // TaskMailbox address from args
        console.log("Task Mailbox:", taskMailbox);

        // Load the private key from the environment variable
        uint256 appPrivateKey = vm.envUint("PRIVATE_KEY_APP");
        address app = vm.addr(appPrivateKey);

        vm.startBroadcast(appPrivateKey);
        console.log("App address:", app);

        // Call createTask
        OperatorSet memory executorOperatorSet = OperatorSet({avs: avs, id: executorOperatorSetId});
        ITaskMailboxTypes.TaskParams memory taskParams = ITaskMailboxTypes.TaskParams({
            refundCollector: address(0),
            executorOperatorSet: executorOperatorSet,
            payload: payload
        });
        bytes32 taskHash = ITaskMailbox(taskMailbox).createTask(taskParams);
        console.log("Created task with hash:");
        console.logBytes32(taskHash);
        ITaskMailboxTypes.Task memory task = ITaskMailbox(taskMailbox).getTaskInfo(taskHash);
        console.log("Task status:", uint8(task.status));
        console.logBytes(task.payload);

        vm.stopBroadcast();
    }
}
