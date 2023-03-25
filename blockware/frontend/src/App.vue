<template>
  <div class="main">
    <!-- Exclude the navbar on the login page -->
    <Navbar v-if="route !== '/'" />

    <!-- show an error at the top of each page -->
    <Error />

    <!-- contents of each page -->
    <div class="router-view">
      <router-view />
    </div>
  </div>
</template>
<script setup>
import { computed, onMounted } from "vue";
import { useRoute } from "vue-router";
import Error from "./components/Error.vue";
import Navbar from "./components/Navbar.vue";
import { useGamesStore } from "./stores/games";
import { usePeerStore } from "./stores/peers";
import { EventsOn } from "../wailsjs/runtime/runtime";

/*

Root component
- content of page based upon router-view
- navbar included on most pages

*/

//
const router = useRoute();
const route = computed(() => router.path);

// instanti
const games = useGamesStore();
const peers = usePeerStore();

onMounted(() => {
  // listen to events emitted by the controller
  EventsOn("update-owned-games", () => games.refreshOwnedGames());
  EventsOn("update-downloads", () => games.refreshDownloads());
  EventsOn("new-peer", () => peers.refreshPeers());

  games.refreshOwnedGames();
});
</script>
<style scoped lang="scss">
.main {
  display: flex;
  flex-direction: column;
  height: 100%;
  max-width: 100vw;
  overflow: hidden;

  > .router-view {
    display: flex;
    flex-direction: column;
    height: 100%;
    overflow: hidden;
  }
}
</style>
