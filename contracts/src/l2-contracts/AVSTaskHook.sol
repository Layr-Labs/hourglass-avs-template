// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import {IAVSTaskHook} from "@eigenlayer-contracts/src/contracts/interfaces/IAVSTaskHook.sol";
import {ITaskMailboxTypes} from "@eigenlayer-contracts/src/contracts/interfaces/ITaskMailbox.sol";

contract AVSTaskHook is IAVSTaskHook {
    function validatePreTaskCreation(
        address, /*caller*/
        ITaskMailboxTypes.TaskParams memory /*taskParams*/
    ) external view {
        //TODO: Implement
    }

    function handlePostTaskCreation(
        bytes32 /*taskHash*/
    ) external {
        //TODO: Implement
    }

    function validatePreTaskResultSubmission(
        address, /*caller*/
        bytes32, /*taskHash*/
        bytes memory, /*cert*/
        bytes memory /*result*/
    ) external view {
        //TODO: Implement
    }

    function handlePostTaskResultSubmission(
        address, /*caller*/
        bytes32 /*taskHash*/
    ) external {
        //TODO: Implement
    }

    function calculateTaskFee(
        ITaskMailboxTypes.TaskParams memory /*taskParams*/
    ) external view returns (uint96) {
        //TODO: Implement
    }
}
