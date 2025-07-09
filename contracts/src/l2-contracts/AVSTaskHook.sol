// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import {OperatorSet} from "@eigenlayer-contracts/src/contracts/libraries/OperatorSetLib.sol";
import {IAVSTaskHook} from "@hourglass-monorepo/src/interfaces/avs/l2/IAVSTaskHook.sol";
import {ITaskMailboxTypes} from "@hourglass-monorepo/src/interfaces/core/ITaskMailbox.sol";

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
        bytes32 /*taskHash*/
    ) external {
        //TODO: Implement
    }

    function calculateTaskFee(
        OperatorSet memory, /*operatorSet*/
        bytes memory /*payload*/
    ) external view returns (uint96) {
        //TODO: Implement
    }
}
