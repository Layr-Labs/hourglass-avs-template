// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import {Test, console} from "forge-std/Test.sol";

import {HelloWorldL2} from "@project/l2-contracts/HelloWorldL2.sol";

contract HelloWorldL2Test is Test {
    HelloWorldL2 public helloWorldL2;

    function setUp() public {
        // Deploy the HelloWorldL2 contract
        helloWorldL2 = new HelloWorldL2();
    }

    function testInitialMessage() view public {
        assertEq(helloWorldL2.getMessage(), "Hello World from L2");
    }

    function testSetMessage() public {
        helloWorldL2.setMessage("New Message from L2");
        assertEq(helloWorldL2.getMessage(), "New Message from L2");
    }
}
