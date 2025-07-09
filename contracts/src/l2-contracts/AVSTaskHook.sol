// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import {IBN254CertificateVerifierTypes} from
    "@eigenlayer-contracts/src/contracts/interfaces/IBN254CertificateVerifier.sol";
import {OperatorSet} from "@eigenlayer-contracts/src/contracts/libraries/OperatorSetLib.sol";

import {IAVSTaskHook} from "@hourglass-monorepo/src/interfaces/avs/l2/IAVSTaskHook.sol";

import {ITaskMailboxTypes} from "@hourglass-monorepo/src/interfaces/core/ITaskMailbox.sol";

contract AVSTaskHook is IAVSTaskHook {
    function validatePreTaskCreation(
        address caller,
        ITaskMailboxTypes.TaskParams memory taskParams
    ) external view override {
        // TODO: Implement
    }

    function handlePostTaskCreation(
        bytes32 taskHash
    ) external override {
        // TODO: Implement
    }

    function validatePreTaskResultSubmission(
        address caller,
        bytes32 taskHash,
        bytes memory cert,
        bytes memory result
    ) external view override {
        // TODO: Implement
    }

    function handlePostTaskResultSubmission(
        bytes32 taskHash
    ) external override {
        // TODO: Implement
    }

    function calculateTaskFee(
        OperatorSet memory operatorSet,
        bytes memory payload
    ) external view override returns (uint96) {
        // TODO: Implement
        return 0;
    }
}
