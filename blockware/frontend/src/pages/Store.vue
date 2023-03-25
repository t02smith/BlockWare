<template>
  <div class="store">
    <div class="title">
      <h2>Welcome to the <strong>BlockWare Store</strong></h2>
      <p>
        Here you can view all the games uploaded to BlockWare and purchase them
        for yourself. Enjoy the selection and be sure to send any feedback to
        <a href="mailto:email@notreal.com">this email</a>.
      </p>
    </div>

    <div class="categories">
      <CustomLibrary
        gameLinkTo="store/entry"
        :games="games.storeGames"
        name="Featured"
      />

      <form class="search" @submit.prevent="searchForGame">
        <div class="header">
          <h3>Search for a Game</h3>
          <p>Fill in the details below to discover a new game</p>
        </div>

        <div class="fields">
          <div class="input">
            <p><strong>Have a game's root hash?</strong></p>
            <input
              v-model="search"
              type="text"
              name=""
              id=""
              placeholder="The game's root hash"
            />
            <button type="submit">Search</button>
          </div>
        </div>

        <div class="fields">
          <p><strong>Enter some info about the game:</strong></p>
          <marquee behavior="" direction=""> ⚠️ TODO ⚠️</marquee>
        </div>
      </form>
    </div>
  </div>
</template>
<script setup>
import { onMounted, ref } from "vue";
import CustomLibrary from "../components/library/CustomLibrary.vue";
import { useGamesStore } from "../stores/games";
import { useRouter } from "vue-router";

/*

Store page to allow users to find new games
Mostly unfinished for this project :()

*/

// hooks
const games = useGamesStore();
const router = useRouter();

// update games on mount
onMounted(() => {
  games.getStoreGames();
});

// root hash to search for a game
const search = ref("");

/*
Open the store page for a game using its root hash
*/
async function searchForGame() {
  if (search.value.length === 0) return;

  router.push({
    path: "/store/entry",
    query: {
      game: search.value,
    },
  });
  search.value = null;
}
</script>
<style scoped lang="scss">
.store {
  display: flex;
  flex-direction: column;
  padding: 2rem;
  height: 100%;

  > .title {
    display: flex;
    flex-direction: column;
    margin-bottom: 2rem;

    > * {
      max-width: 700px;
    }

    > h2 {
      font-size: 2rem;
    }

    > p {
      color: darken(white, 15%);
    }
  }

  > .categories {
    display: grid;
    grid-template-columns: 1fr 1fr;
    height: 100%;

    > .category {
      > h3 {
        color: orangered;
        font-size: 1.4rem;
      }
    }
  }

  .search {
    display: flex;
    flex-direction: column;
    margin: 1rem;
    gap: 1rem;

    > .header {
      display: flex;
      flex-direction: column;
    }

    > .fields {
      > .input {
        p {
          margin-bottom: 3px;
        }

        input {
          padding: 0.5rem 0.75rem;
          width: 300px;
          border-radius: 2px 0 0 2px;
          border: none;
        }

        button {
          border-radius: 0 2px 2px 0;
          border: none;
          padding: 0.5rem 0.75rem;
          background-color: rgb(0, 132, 255);
          font-weight: bold;
          color: white;
          cursor: pointer;
        }
      }
    }
  }
}
</style>
