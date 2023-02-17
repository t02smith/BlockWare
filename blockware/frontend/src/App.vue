<template>
  <div class="main">
    <Navbar v-if="route !== '/'" />

    <div class="router-view">
      <router-view />
    </div>
  </div>
</template>
<script setup>
import { computed } from "vue";
import { useRoute } from "vue-router";
import Navbar from "./components/Navbar.vue";

const router = useRoute();

const route = computed(() => router.path);
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
