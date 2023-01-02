// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.9;

contract Game {
    // information about the game
    string public title;
    string public version;
    int public price;

    Developer public developer;

    struct Developer {
        string name;
        address payable addr;
        // used for fetching SSL certificate
        string domain;
    }

    constructor(
        string memory _title,
        string memory _version,
        int _price,
        Developer memory _developer
    ) {
        title = _title;
        version = _version;
        price = _price;
        developer = _developer;
    }
}
