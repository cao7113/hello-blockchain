// SPDX-License-Identifier: GPL-3.0
pragma solidity >=0.8.13 <0.9.0;

contract HelloWorld {
    event log_string(bytes32 log);

    fallback() external {
        emit log_string("Hello World1!");
    }
}
