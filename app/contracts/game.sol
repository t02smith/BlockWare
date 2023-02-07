// SPDX-License-Identifier: SEE LICENSE IN LICENSE
pragma solidity ^0.8.18;

contract Game {
    // triggered every time a game is uploaded
    event NewGame(bytes32 hash, GameEntry game);

    // game root hash => entry
    mapping(bytes32 => GameEntry) games;

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
        emit NewGame(_game.rootHash, _game);
    }
}
