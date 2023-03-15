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

    <form class="search" @submit.prevent="searchForGame">
      <h3>Find a Game</h3>

      <div class="input">
        <input
          v-model="search"
          type="text"
          name=""
          id=""
          placeholder="The game's root hash"
        />
        <button type="submit">Search</button>
      </div>
    </form>

    <div class="searched-game" v-if="searchedGame"></div>

    <div class="categories">
      <CustomLibrary
        gameLinkTo="store/entry"
        :games="games.storeGames"
        name="Featured"
      />

      <CustomLibrary
        gameLinkTo="store/entry"
        :games="games.storeGames"
        name="New Releases"
      />
    </div>
  </div>
</template>
<script setup>
import { onMounted, ref } from "vue";
import CustomLibrary from "../components/library/CustomLibrary.vue";
import { useGamesStore } from "../stores/games";
import { useRouter } from "vue-router";

const games = useGamesStore();
const router = useRouter();

onMounted(() => {
  games.getStoreGames();
});

// search
const search = ref("");

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

  > .search {
    display: flex;
    align-items: center;
    justify-content: center;
    gap: 1rem;

    > .input {
      input {
        padding: 0.5rem 0.75rem;
        width: 300px;
        border-radius: 5px 0 0 5px;
        border: none;
        outline: none;
      }

      button {
        border-radius: 0 5px 5px 0;
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
</style>
