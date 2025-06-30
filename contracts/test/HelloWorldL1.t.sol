// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import {Test, console} from "forge-std/Test.sol";

import {HelloWorldL1} from "@project/l1-contracts/HelloWorldL1.sol";

contract HelloWorldL1Test is Test {
    HelloWorldL1 public helloWorldL1;

    function setUp() public {
        // Deploy the HelloWorldL1 contract
        helloWorldL1 = new HelloWorldL1();
    }

    function testInitialMessage() view public {
        assertEq(helloWorldL1.getMessage(), "Hello World from L1");
    }

    function testSetMessage() public {
        helloWorldL1.setMessage("New Message from L1");
        assertEq(helloWorldL1.getMessage(), "New Message from L1");
    }
}
