<template>
  <div class="library">
    <header>
      <h2>Your Library</h2>

      <div class="right">
        <h2 class="game-count">
          🕹️{{ games.length }} {{ games.length === 1 ? "Game" : "Games" }}
        </h2>

        <button class="refresh" @click="refreshLibrary">
          <h2>♻️Refresh</h2>
        </button>
      </div>
    </header>
    <div class="games">
      <AddGames />
      <Game
        v-for="g in games"
        :title="g.title"
        :version="g.version"
        :dev="g.dev"
      />
    </div>
  </div>
</template>
<script setup>
import { onMounted, ref } from "vue";
import Game from "./Game.vue";
import { GetOwnedGames } from "../../../wailsjs/go/main/App.js";
import AddGames from "./AddGames.vue";

const games = ref([]);

onMounted(() => refreshLibrary());

async function refreshLibrary() {
  games.value = await GetOwnedGames();
}
</script>
<style scoped lang="scss">
.library {
  background-color: lighten(#131313, 5%);
  border-radius: 10px;
  margin: 1rem;

  > .games {
    display: flex;
    align-items: center;
  }

  > header {
    display: flex;
    align-items: center;
    padding: 0.5rem;
    background-color: lighten(#131313, 15%);
    border-radius: 10px 10px 0 0;

    > .right {
      margin-left: auto;
      display: flex;
      align-items: center;
      gap: 1rem;
      margin-right: 1rem;

      .game-count {
        color: rgb(9, 109, 240);
      }

      .refresh {
        background-color: transparent;
        border: none;
        outline: none;
        color: darken(white, 20%);
        transition: 150ms;
        cursor: pointer;

        &:hover {
          color: rgb(12, 187, 12);
        }
      }
    }
  }
}
</style>
