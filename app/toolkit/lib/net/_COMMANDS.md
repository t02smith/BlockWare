# Client-server commands

Each peer in the network will:

- Run a server socket to allow other peers to connect to and
  send messages to, and
- Make connections with other peer's server sockets

## **1. Client => Server**

### *Library 0x00*

The server sends a transcript of their game library by sending
a series of hashes that represent each game.

```asciidoc
= Response =
CLIENT: 0x0
SERVER: <number of games: uint16>:<hash #1: [32]byte>:<hash #2: [32]byte>:...;
```

### *Shards 0x01*

The server will send a list of blocks that it owns of the given game
if it has it at all

```asciidoc
= Parameters =
rootHash => The root hash of a game 
```

```asciidoc
= Response =
CLIENT: 0x1:<rootHash>;
SERVER: <number of blocks: uint32>:
```

## **2. Server => Client**
