// SPDX-License-Identifier: SEE LICENSE IN LICENSE
pragma solidity ^0.8.18;

contract Library {
    // triggered every time a game is uploaded
    event NewGame(bytes32 hash, GameEntry game);

    // game root hash => entry
    mapping(bytes32 => GameEntry) public games;

    // game root hash => address => purchased
    // addresses are only included when a user has purchased
    mapping(bytes32 => mapping(address => bool)) purchases;

    // a list of all hashes
    bytes32[] public gameHashes;

    // metadata about each game
    struct GameEntry {

        // game meta data
        string title;
        string version;
        string releaseDate;
        string developer;
        bytes32 rootHash;
        bytes32 previousVersion;

        // purchasing
        uint price;
        address payable uploader;

        // address to download hash data from IPFS
        string ipfsAddress;
    }

    constructor() {}

    /**
     * Upload a new game to the network
     * @param _game the details about the game
     */
    function uploadGame(GameEntry memory _game) external {
        // check input data
        require(_game.rootHash.length > 0, "no root hash given");
        require(bytes(_game.ipfsAddress).length > 0, "no IPFS address given for hash treee");

        // look for previous version
        if (_game.previousVersion != 0) {
          require(bytes(games[_game.previousVersion].title).length > 0, "previous version of game not found");

          GameEntry memory g = games[_game.previousVersion];
          require(g.uploader == msg.sender, "only the original uploader can update their game");
          purchases[_game.rootHash][msg.sender] = true;
        }

        // upload game
        _game.uploader = payable(msg.sender);
        games[_game.rootHash] = _game;
        gameHashes.push(_game.rootHash);
        emit NewGame(_game.rootHash, _game);
    }

    /**
     * Purchase a new game
     * @param _game the root hash of the game
     */
    function purchaseGame(bytes32 _game) public payable {
      require(bytes(games[_game].title).length > 0, "game not found");
      
      GameEntry storage game = games[_game];
      require(msg.value >= game.price, "user cannot afford game");

      require(purchases[_game][msg.sender], "user already owns game");

      game.uploader.transfer(game.price);
      purchases[_game][msg.sender] = true;
    }

    /**
     * How many games exist in the current library
     */
    function libSize() public view returns (uint) {
        return gameHashes.length;
    }

    /**
     * Has a given user purchased a game
     * @param _game The root hash of the chosen game
     * @param _addr The address of the person to check
     */
    function hasPurchased(bytes32 _game, address _addr) public view returns (bool) {
      return purchases[_game][_addr];
    }
}
