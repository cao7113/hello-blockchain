// SPDX-License-Identifier: GPL-3.0
pragma solidity >=0.8.13 <0.9.0;

contract HelloEvent {
    event log_string(bytes32 log);

    function Hello() public returns (string memory) {
        emit log_string("Hello World!");
        return "Hello World";
    }

    function Greet(string memory str) public pure returns (string memory) {
        return str;
    }
}
