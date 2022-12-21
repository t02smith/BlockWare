// SPDX-License-Identifier: MIT
pragma solidity >=0.4.22 <0.9.0;

contract Game {
    address public author;
    uint256 public price;
    GameData data;

    struct GameData {
        string name;
        string version;
        GameDirectory rootDir;
    }

    struct GameDirectory {
        string name;
        GameFile[] files;
        GameDirectory[] subdirs;
    }

    struct GameFile {
        string name;
        bytes32[] hashes;
    }
}
