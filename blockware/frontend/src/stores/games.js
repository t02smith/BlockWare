import { defineStore } from "pinia";
import { ref, onMounted } from "vue";
import {
  GetOwnedGames,
  GetStoreGames,
  GetDownloads,
  CreateDownload,
} from "../../wailsjs/go/main/App.js";

export const useGamesStore = defineStore("games", () => {
  // what games do they own
  const ownedGames = ref([]);

  // what games are being downloaded
  const downloads = ref([]);

  //
  const storeGames = ref([]);

  // setup
  onMounted(() => {
    refreshOwnedGames();
    refreshDownloads();
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

  // add listeners for the progress of downloads
  async function setupDownloadListeners() {}

  async function getStoreGames() {
    storeGames.value = await GetStoreGames();
  }

  return {
    ownedGames,
    downloads,
    storeGames,
    createDownload,
    refreshDownloads,
    refreshOwnedGames,
    getStoreGames,
  };
});
