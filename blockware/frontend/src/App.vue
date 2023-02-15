<script setup>
import Navbar from "./components/Navbar.vue";
</script>

<template>
  <div class="main">
    <Navbar />

    <div class="router-view">
      <router-view />
    </div>
  </div>
</template>

<style scoped lang="scss">
.main {
  display: flex;
  flex-direction: column;

  > .router-view {
    padding: 1rem;
  }
}
</style>
