// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import {IAllocationManager} from "@eigenlayer-contracts/src/contracts/interfaces/IAllocationManager.sol";
import {IKeyRegistrar} from "@eigenlayer-contracts/src/contracts/interfaces/IKeyRegistrar.sol";

import {TaskAVSRegistrarBase} from "@hourglass-monorepo/src/avs/TaskAVSRegistrarBase.sol";

contract TaskAVSRegistrar is TaskAVSRegistrarBase {
    constructor(
        address avs,
        IAllocationManager allocationManager,
        IKeyRegistrar keyRegistrar,
        address owner,
        AvsConfig memory initialConfig
    ) TaskAVSRegistrarBase(avs, allocationManager, keyRegistrar) {}

    // TODO: Implement
}
