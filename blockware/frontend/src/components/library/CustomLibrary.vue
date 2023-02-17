<template>
  <div class="library">
    <header>
      <h2>{{ props.name }}</h2>

      <div class="right">
        <h2 class="game-count">
          🕹️{{ props.games.length }}
          {{ props.games.length === 1 ? "Game" : "Games" }}
        </h2>
      </div>
    </header>
    <div class="games">
      <Game
        v-if="props.games.length"
        v-for="g in props.games.slice(0, 4)"
        :title="g.title"
        :version="g.version"
        :dev="g.dev"
        :hash="toHexString(g.rootHash)"
      />
      <div v-else class="no-games">
        <h3>Ooops no games here...</h3>
      </div>
    </div>
  </div>
</template>
<script setup>
import Game from "./Game.vue";
import { toHexString } from "../../util/util";

const props = defineProps({
  games: {
    required: true,
    type: Array,
  },
  name: {
    type: String,
    required: true,
  },
});
</script>
<style scoped lang="scss">
.library {
  background-color: lighten(#131313, 5%);
  border-radius: 10px;
  margin: 1rem;

  > .games {
    display: flex;
    align-items: center;

    > .no-games {
      height: 200px;
      width: 400px;
      margin: 0.5rem;

      > h3 {
        margin: 1rem;
      }
    }
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
