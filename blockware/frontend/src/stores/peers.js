import { defineStore } from "pinia";
import { ref, onMounted } from "vue";
import {
  GetPeerInformation,
  ConnectToPeer,
  ConnectToManyPeers,
  Disconnect,
  ResendValidation,
  ConnectFromFile,
} from "../../wailsjs/go/controller/Controller";
import { EventsOn } from "../../wailsjs/runtime/runtime";

/**

Peer Store:
- store state about connected peers
- have interface functions to the Controller functions

*/
export const usePeerStore = defineStore("peers", () => {
  // list of connected peers
  const peers = ref([]);

  /**
   * Refresh the current list of peers
   */
  async function refreshPeers() {
    peers.value = await GetPeerInformation();
  }

  // wrapper functions

  /**
   * Disconnect to an existing peer
   * @param {String} hostname their hostname
   * @param {Number} port their port number
   */
  function disconnect(hostname, port) {
    Disconnect(hostname, port);
  }

  /**
   * Connect to a new peer
   * @param {String} hostname their hostname
   * @param {Number} port their port number
   */
  function connect(hostname, port) {
    ConnectToPeer(hostname, port);
  }

  /**
   * Connect to many peers at once
   * Each peer should be of the form "{hostname}:{port}"
   * @param {Array} peerLs
   */
  function connectToAll(peerLs) {
    ConnectToManyPeers(peerLs);
  }

  function connectFromFile(filepath) {
    ConnectFromFile(filepath);
  }

  /**
   * Resend a validation method to a given peer
   * A validation message is used to verify a peer's Eth address
   * @param {String} hostname their hostname
   * @param {Number} port their port number
   */
  function resendValidation(hostname, port) {
    ResendValidation(hostname, poort);
  }

  return {
    peers,
    connect,
    refreshPeers,
    connectToAll,
    disconnect,
    resendValidation,
    connectFromFile,
  };
});
