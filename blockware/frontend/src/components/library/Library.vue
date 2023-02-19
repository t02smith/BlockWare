<template>
  <div class="library">
    <header>
      <h2>Your Library</h2>

      <div class="right">
        <h2 class="game-count">
          🕹️{{ games.ownedGames.length }}
          {{ games.ownedGames.length === 1 ? "Game" : "Games" }}
        </h2>

        <button class="refresh" @click="games.refreshOwnedGames">
          <h2>♻️Refresh</h2>
        </button>
      </div>
    </header>
    <div class="games">
      <AddGames />
      <Game
        v-for="g in games.ownedGames.slice(0, 4)"
        :title="g.title"
        :version="g.version"
        :dev="g.dev"
        :hash="g.rootHash"
      />
    </div>
  </div>
</template>
<script setup>
import Game from "./Game.vue";
import AddGames from "./AddGames.vue";
import { useGamesStore } from "../../stores/games";

const games = useGamesStore();
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
