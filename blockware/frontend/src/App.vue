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
import { computed } from "vue";
import { useRoute } from "vue-router";
import Error from "./components/Error.vue";
import Navbar from "./components/Navbar.vue";
import { useGamesStore } from "./stores/games";

const router = useRoute();
const route = computed(() => router.path);

useGamesStore();
</script>
<style scoped lang="scss">
.main {
  display: flex;
  flex-direction: column;
  height: 100vh;

  > .router-view {
    width: 100vw;
    height: 100%;
  }
}
</style>
