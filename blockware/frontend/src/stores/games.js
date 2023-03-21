import { defineStore } from "pinia";
import { ref, onMounted } from "vue";
import {
  GetOwnedGames,
  GetStoreGames,
  GetDownloads,
  CreateDownload,
  PurchaseGame,
  ContinueAllDownloads,
  FetchOwnedGame,
  LoadDeferredRequests,
} from "../../wailsjs/go/controller/Controller";
import { EventsOn } from "../../wailsjs/runtime/runtime";

export const useGamesStore = defineStore("games", () => {
  // what games do they own
  const ownedGames = ref([]);

  // what games are being downloaded
  const downloads = ref([]);

  // games in the blockware store
  const storeGames = ref([]);

  // setup
  onMounted(() => {
    refreshOwnedGames();
    refreshDownloads();
    ContinueAllDownloads();
  });

  // get an updated list of owned games
  async function refreshOwnedGames() {
    ownedGames.value = await GetOwnedGames();
  }

  // get an updated list of downloads
  async function refreshDownloads() {
    downloads.value = await GetDownloads();
  }

  // create a new download for an existing game
  async function createDownload(gameHash) {
    const success = await CreateDownload(gameHash);
    if (!success) return;

    await refreshDownloads();
  }

  function importGame(gameHash) {
    FetchOwnedGame(gameHash);
  }

  async function getStoreGames() {
    storeGames.value = await GetStoreGames();
  }

  async function purchase(gameHash) {
    if (ownedGames.value.find((g) => g.rootHash === gameHash)) return;

    await PurchaseGame(gameHash);
  }

  function loadDeferredRequests() {
    LoadDeferredRequests();
  }

  return {
    ownedGames,
    downloads,
    storeGames,
    createDownload,
    refreshDownloads,
    refreshOwnedGames,
    getStoreGames,
    purchase,
    importGame,
    loadDeferredRequests,
  };
});
