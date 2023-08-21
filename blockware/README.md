# Blockware

This is the source code for the Blockware application.

## How to Install

### Prerequisites

To install this application you will at minimum need the following:

- Go v1.19 [https://go.dev/dl/](https://go.dev/dl/)
- Node v16 [https://nodejs.org/en/blog/release/v16.16.0](https://nodejs.org/en/blog/release/v16.16.0)
- Kubo IPFS [https://docs.ipfs.tech/install/](https://docs.ipfs.tech/install/)
- Wails [https://wails.io/docs/gettingstarted/installation](https://wails.io/docs/gettingstarted/installation)

Make sure you have access to an Ethereum node. It is recommended that you use Alchemy [https://www.alchemy.com/](https://www.alchemy.com/) where you can get an RPC link.

### Build

You can build the application using the following command

```bash
wails build
```

This will create an executable in `/build/bin/blockware(.exe)?`. It is suggested you move the executable to a separate folder.

### How to Run

```bash
# (if this is your first time using IPFS)
ipfs init

# start the IPFS daemon
ipfs daemon

# run the application
./${path to application}/blockware.exe
```