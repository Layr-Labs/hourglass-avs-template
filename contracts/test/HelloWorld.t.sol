// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import {Test, console} from "forge-std/Test.sol";

import {HelloWorld} from "@project/HelloWorld.sol";

contract HelloWorldTest is Test {
    HelloWorld public helloWorld;

    function setUp() public {
        // Deploy the HelloWorld contract
        helloWorld = new HelloWorld();
    }

    function testSetMessage() public {
        helloWorld.setMessage("Hello World");
        assertEq(helloWorld.getMessage(), "Hello World");
    }
}
