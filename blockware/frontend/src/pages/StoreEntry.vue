<template>
  <div class="wrapper">
    <GameEntry v-if="game" :game="game">
      <div class="purchase">
        <strong>{{ game.price }} ETH</strong>
        <button @click="purchase">🛒 Purchase</button>
      </div>
    </GameEntry>

    <p v-else>game not found</p>
  </div>
</template>
<script setup>
import { ref, onMounted } from "vue";
import { useRoute } from "vue-router";
import { useGamesStore } from "../stores/games";
import GameEntry from "../components/store/GameEntry.vue";

const route = useRoute();
const games = useGamesStore();

//
const game = ref(null);

onMounted(async () => {
  await games.getStoreGames();

  const gameHash = route.query.game;
  if (!gameHash) return;

  game.value = games.storeGames.find((g) => gameHash === g.rootHash);
});

async function purchase() {
  if (!game.value) return;

  await games.purchase(game.value.rootHash);
}
</script>
<style scoped lang="scss">
.wrapper {
  display: flex;
  justify-content: center;
  overflow-x: hidden;

  > * {
    max-width: 80vw;
  }
}

.purchase {
  display: flex;
  flex-direction: column;
  align-items: flex-end;
  gap: 0.5rem;

  > strong {
    font-size: 2rem;
  }

  > button {
    border-radius: 4px;
    padding: 0.25rem 1rem;
    font-weight: bold;
    background-color: rgb(142, 7, 7);
    cursor: pointer;
    transition: 75ms;

    &:active {
      scale: 0.99;
    }

    &:hover {
      scale: 1.01;
    }
  }
}
</style>
