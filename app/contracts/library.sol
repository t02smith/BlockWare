// SPDX-License-Identifier: SEE LICENSE IN LICENSE
pragma solidity ^0.8.18;

contract Library {
    // triggered every time a game is uploaded
    event NewGame(bytes32 hash, GameEntry game);

    // game root hash => entry
    mapping(bytes32 => GameEntry) public games;

    // a list of all hashes
    bytes32[] public gameHashes;

    // details about each game
    struct GameEntry {
        // game meta data
        string title;
        string version;
        string releaseDate;
        string developer;
        bytes32 rootHash;
        // address to download hash data from IPFS
        string ipfsAddress;
    }

    constructor() {}

    /**
     * Upload a new game to the network
     * @param _game the details about the game
     */
    function uploadGame(GameEntry memory _game) external {
        games[_game.rootHash] = _game;
        gameHashes.push(_game.rootHash);
        emit NewGame(_game.rootHash, _game);
    }

    /**
     *
     */
    function libSize() public view returns (uint) {
        return gameHashes.length;
    }
}
