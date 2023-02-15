<template>
  <div class="library">
    <h2>Your Library</h2>
    <div class="games">
      {{ games }}
      <!-- <Game
        v-for="g in games"
        :title="g.T"
        version="1.4.6"
        dev="CD Project Red"
        release="12/10/2014"
      />
      <Game
        title="The Witcher 3"
        version="1.4.6"
        dev="CD Project Red"
        release="12/10/2014"
      /> -->
    </div>
  </div>
</template>
<script setup>
import { onMounted, ref } from "vue";
import Game from "./Game.vue";
import { GetOwnedGames } from "../../wailsjs/go/main/App.js";

const games = ref([]);

onMounted(async () => {
  games.value = await GetOwnedGames();
});
</script>
<style scoped lang="scss">
.library {
  > .games {
    display: flex;
    gap: 1rem;
  }
}
</style>
