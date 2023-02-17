import { defineStore } from "pinia";
import { ref, onMounted } from "vue";
import { GetOwnedGames } from "../../wailsjs/go/main/App.js";

export const useGamesStore = defineStore("games", () => {
  const ownedGames = ref([]);

  onMounted(() => {
    refreshOwnedGames();
  });

  async function refreshOwnedGames() {
    ownedGames.value = await GetOwnedGames();
  }

  return { ownedGames, refreshOwnedGames };
});
