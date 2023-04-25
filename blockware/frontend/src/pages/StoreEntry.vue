<template>
  <div class="wrapper">
    <Error />
    <GameEntry v-if="game" :game="game">
      <div class="purchase">
        <strong>{{ game.price }} Wei</strong>

        <router-link
          :to="`/library?game=${game.rootHash}`"
          v-if="ownsGame"
          class="owned"
          >📖 View in Library</router-link
        >
        <button @click="purchase" v-else>Purchase</button>
      </div>
    </GameEntry>

    <div class="not-found" v-else>
      <h2>🥲 Game not found</h2>
      <p>
        You searched for <strong>{{ route.query.game }}</strong>
      </p>
      <br />
      <router-link to="/store">Click here to return to the store!</router-link>
    </div>
  </div>
</template>
<script setup>
import { ref, onMounted, computed } from "vue";
import { useRoute } from "vue-router";
import { useGamesStore } from "../stores/games";
import GameEntry from "../components/library/GameEntry.vue";
import Error from "../components/Error.vue";

/*

A store page for a given game that will show users the description
and allow them to purchase it

*/

// pinia stores
const games = useGamesStore();

// selected game
const game = ref(null);
const ownsGame = computed(
  () =>
    game.value &&
    games.ownedGames.find((g) => g.rootHash === game.value.rootHash)
);

//

// The game is identified by its root hash in the URL
const route = useRoute();

// attempt to fetch that game to display it
onMounted(async () => {
  const gameHash = route.query.game;
  if (!gameHash) return;

  game.value = games.storeGames.find((g) => gameHash === g.rootHash);
});

const submitted = ref(false);

/*
purchase the currently on display game if there is oen
*/
async function purchase() {
  if (!game.value) return;

  submitted.value = true;
  await games.purchase(game.value.rootHash);
  await games.refreshOwnedGames();
  submitted.value = false;
}
</script>
<style scoped lang="scss">
.wrapper {
  display: flex;
  justify-content: center;
  height: 100%;
  overflow: hidden;
  background-color: lighten(#131313, 2%);

  > * {
    max-width: 80vw;
  }
}

.not-found {
  background-color: lighten(#131313, 6%);
  height: fit-content;
  padding: 1.5rem;
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  box-shadow: 0 4px 8px 0 rgba(0, 0, 0, 0.2), 0 6px 20px 0 rgba(0, 0, 0, 0.19);

  border-radius: 5px;

  > h2 {
    font-size: 3rem;
  }

  > p {
    font-style: italic;
  }

  a {
    color: rgb(0, 162, 255);
    text-decoration: none;
    cursor: pointer;

    &:hover {
      text-decoration: underline;
    }
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

  .owned {
    color: white;
    text-decoration: none;
    background-color: green;
  }

  > button,
  .owned {
    border-radius: 4px;
    padding: 0.25rem 1rem;
    font-weight: bold;
    background-color: rgb(219, 12, 12);
    cursor: pointer;
    transition: 75ms;
    border: none;

    &:active {
      scale: 0.99;
    }

    &:hover {
      scale: 1.01;
    }
  }
}
</style>
