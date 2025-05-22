// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract HelloWorld {
    string private message = "Hello World";

    // Function to get the hello world message
    function getMessage() public view returns (string memory) {
        return message;
    }

    // Function to update the message (optional)
    function setMessage(string memory newMessage) public {
        message = newMessage;
    }
}
