<template>
  <div class="main">
    <Navbar v-if="route !== '/'" />

    <Error />

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

const router = useRoute();
const route = computed(() => router.path);

const games = useGamesStore();
const peers = usePeerStore();

onMounted(() => {
  EventsOn("update-owned-games", async () => await games.refreshOwnedGames());
  EventsOn("new-peer", async () => await peers.refreshPeers());
  EventsOn("update-downloads", async () => await games.refreshDownloads());

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
