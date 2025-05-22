// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import {Test, console} from "forge-std/Test.sol";

import {IAllocationManager} from "@eigenlayer-contracts/src/contracts/interfaces/IAllocationManager.sol";

import {TaskAVSRegistrar} from "@project/l1-contracts/TaskAVSRegistrar.sol";

contract TaskAVSRegistrarTest is Test {
    TaskAVSRegistrar public taskAVSRegistrar;

    function setUp() public {
        // Deploy the TaskAVSRegistrar contract
        taskAVSRegistrar = new TaskAVSRegistrar(address(0), IAllocationManager(address(0)));
    }

    function testDummy() public pure returns (bool) {
        return true;
    }
}
