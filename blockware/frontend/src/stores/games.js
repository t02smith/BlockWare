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
  CheckForUpdates,
  CancelDownload,
  PauseDownload,
  ContinueDownload,
} from "../../wailsjs/go/controller/Controller";
import { EventsOn } from "../../wailsjs/runtime/runtime";

/**
 * Manage the collection of games known by the application.
 * Includes:
 * - games owned by the user
 * - games being downloaded
 * - games that are present on the store
 */
export const useGamesStore = defineStore("games", () => {
  // what games do they own
  const ownedGames = ref([]);

  // what games are being downloaded
  const downloads = ref([]);

  // games in the blockware store
  const storeGames = ref([]);

  // setup => refresh data
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
    await CreateDownload(gameHash);
    await refreshDownloads();
  }

  // pause an in progress download
  async function pauseDownload(gameHash) {
    await PauseDownload(gameHash);
    await refreshDownloads();
  }

  async function cancelDownload(gameHash) {
    await CancelDownload(gameHash);
    await refreshDownloads();
  }

  async function continueDownload(gameHash) {
    await ContinueDownload(gameHash);
    await refreshDownloads();
  }

  /**
   * Import a new game that the user owns but doesn't have locally
   * @param {String} gameHash its unique root hash
   */
  function importGame(gameHash) {
    FetchOwnedGame(gameHash);
  }

  /**
   * Fetch some games from the store
   */
  async function getStoreGames() {
    storeGames.value = await GetStoreGames();
  }

  /**
   * Purchase a game from the store
   * @param {String} gameHash its unique root hash
   */
  async function purchase(gameHash) {
    if (ownedGames.value.find((g) => g.rootHash === gameHash)) return;

    await PurchaseGame(gameHash);
  }

  /**
   * Load any requests that have bee postponed
   */
  function loadDeferredRequests() {
    LoadDeferredRequests();
  }

  /**
   * Check for any updates to existing games
   */
  async function checkForUpdates() {
    await CheckForUpdates();
    await refreshOwnedGames();
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
    checkForUpdates,
    pauseDownload,
    cancelDownload,
    continueDownload,
  };
});
