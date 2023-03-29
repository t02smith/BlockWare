<template>
  <div class="main">
    <!-- Exclude the navbar on the login page -->
    <Navbar v-if="route !== '/'" />

    <!-- contents of each page -->
    <div class="router-view">
      <router-view />
    </div>
  </div>
</template>
<script setup>
import { computed, onMounted, ref } from "vue";
import { useRoute } from "vue-router";
import Error from "./components/Error.vue";
import Navbar from "./components/Navbar.vue";
import { useGamesStore } from "./stores/games";
import { usePeerStore } from "./stores/peers";
import { EventsOn } from "../wailsjs/runtime/runtime";
import { useErrStore } from "./stores/err";

/*

Root component
- content of page based upon router-view
- navbar included on most pages

*/

//
const router = useRoute();
const route = computed(() => router.path);

// hooks
const games = useGamesStore();
const peers = usePeerStore();
const err = useErrStore();

onMounted(() => {
  // listen to events emitted by the controller
  EventsOn("update-owned-games", () => games.refreshOwnedGames());
  EventsOn("update-downloads", () => games.refreshDownloads());
  EventsOn("new-peer", () => peers.refreshPeers());
  EventsOn("error", (newErr) => err.set(newErr));

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
