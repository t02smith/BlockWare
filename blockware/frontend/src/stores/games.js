import { defineStore } from "pinia";
import { ref, onMounted } from "vue";
import { GetOwnedGames, GetAllGames } from "../../wailsjs/go/main/App.js";

export const useGamesStore = defineStore("games", () => {
  const ownedGames = ref([]);

  onMounted(() => {
    refreshOwnedGames();
  });

  async function refreshOwnedGames() {
    ownedGames.value = await GetOwnedGames();
    await GetAllGames();
  }

  return { ownedGames, refreshOwnedGames };
});
