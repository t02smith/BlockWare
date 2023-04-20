// SPDX-License-Identifier: SEE LICENSE IN LICENSE
pragma solidity ^0.8.18;

/**
 * @title Library
 * @author tcs1g20 @ University of Southampton
 * @notice Manage a distributed marketplace for video games
 */
contract Library {

    /// @notice called when a game is uploaded
    /// @param hash the game's unique SHA-256 hash
    /// @param game the corresponding entry of the new game
    event NewGame(bytes32 hash, GameEntry game);

    /// @notice called when a game is purchased by a new user
    /// @param hash the game's unique SHA-256 hash
    /// @param buyer the address of the user buying the game
    /// @param data the data returned from the ether transfer 
    event Purchase(bytes32 hash, address buyer, bytes data);

    /// @notice maps game hashes to their metadata
    mapping(bytes32 => GameEntry) public games;

    /// @notice for each game maps what addresses have purchased it or not
    mapping(bytes32 => mapping(address => uint8)) purchases;

    /// @notice the list of all root hashes uploaded
    bytes32[] public gameHashes;

    /// @notice the metadata provided for each game
    struct GameEntry {

        // game meta data
        string title;
        string version;
        string releaseDate;
        string developer;
        bytes32 rootHash;

        // versioning data
        bytes32 previousVersion;
        bytes32 nextVersion;

        // purchasing
        uint256 price;
        address uploader;

        // address to download hash data from IPFS
        string hashTreeIPFSAddress;
        string assetsIPFSAddress;
    }

    /**
     * @notice upload a new game to the network
     * @param _game the details about the game
     */
    function uploadGame(GameEntry memory _game) external {
        // check input data
        require(bytes(_game.hashTreeIPFSAddress).length > 0, "no IPFS address given for hash treee");
        require(bytes(_game.assetsIPFSAddress).length > 0, "no IPFS address given for the assets ");

        // look for previous version
        bytes32 empty = bytes32(0x0);
        if (_game.previousVersion != empty) {
          require(bytes(games[_game.previousVersion].title).length > 0, "previous version of game not found");

          GameEntry storage g = games[_game.previousVersion];
          require(g.uploader == msg.sender, "only the original uploader can update their game");

          require(g.nextVersion == empty, "an update has already been released for this game");
          g.nextVersion = _game.rootHash;
          purchases[_game.rootHash][msg.sender] = 1;
        }

        // upload game
        _game.uploader = payable(msg.sender);
        games[_game.rootHash] = _game;
        gameHashes.push(_game.rootHash);
        emit NewGame(_game.rootHash, _game);
    }

    /**
     * @notice purchase a new game
     * @param _game the root hash of the game
     */
    function purchaseGame(bytes32 _game) external payable {
      require(bytes(games[_game].title).length > 0, "game not found");
      require(purchases[_game][msg.sender] == 0, "user already owns game");
      require(games[_game].price == msg.value, "unexpected price => value should equal the game's pricce");
      
      (bool sent, bytes memory data) = games[_game].uploader.call{value: msg.value}(bytes(games[_game].title));
      require(sent, "Failed to transfer Ether");

      purchases[_game][msg.sender] = 1;
      emit Purchase(_game, msg.sender, data);
    }

    /**
     * @notice How many games exist in the current library
     * @return uint number of games present in the library
     */
    function libSize() external view returns (uint) {
        return gameHashes.length;
    }

    /**
     * @notice Has a given user purchased a game
     * @param _game The root hash of the chosen game
     * @param _addr The address of the person to check
     * @return bool Whether the address has purchased the game
     */
    function hasPurchased(bytes32 _game, address _addr) external view returns (bool) {
      if (purchases[_game][_addr] == 1) {
        return true;
      }

      bytes32 empty = bytes32(0x0);
      GameEntry memory game = games[_game];
      while (game.previousVersion != empty) {
        game = games[game.previousVersion];
        if (purchases[game.rootHash][_addr] == 1) {
          return true;
        }
      }

      return false;
    }

    /// @notice get the data for the most recent version of a game
    /// @param _game the root hash of the requested game
    /// @return bytes32 the root hash of the most recent version
    function getMostRecentVersion(bytes32 _game) external view returns (bytes32) {
      require(bytes(games[_game].title).length > 0, "game not found");
      GameEntry memory game = games[_game];

      bytes32 empty = bytes32(0x0);
      while (game.nextVersion != empty) {
        game = games[game.nextVersion];
      }

      if (game.rootHash == _game) return empty;
      return game.rootHash;
    }
}
