import { defineStore } from "pinia";
import { ref, onMounted } from "vue";
import {
  GetPeerInformation,
  ConnectToPeer,
  ConnectToManyPeers,
} from "../../wailsjs/go/controller/Controller";
import { EventsOn } from "../../wailsjs/runtime/runtime";

export const usePeerStore = defineStore("peers", () => {
  const peers = ref([]);

  onMounted(() => {
    refreshPeers();
  });

  async function refreshPeers() {
    peers.value = await GetPeerInformation();
  }

  function connect(hostname, port) {
    ConnectToPeer(hostname, port);
  }

  function connectToAll(peerLs) {
    ConnectToManyPeers(peerLs);
  }

  return { peers, connect, refreshPeers, connectToAll };
});
