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

    <div class="search">
      <h3>Search for a game:</h3>
      <input
        disabled
        type="text"
        name=""
        id=""
        placeholder="The game's root hash"
      />
    </div>

    <div class="featured">
      <CustomLibrary :games="games.storeGames" name="Featured" />
    </div>
  </div>
</template>
<script setup>
import { onMounted } from "vue";
import CustomLibrary from "../components/library/CustomLibrary.vue";
import { useGamesStore } from "../stores/games";

const games = useGamesStore();

onMounted(async () => {
  games.getStoreGames();
});
</script>
<style scoped lang="scss">
.store {
  display: flex;
  flex-direction: column;
  padding: 2rem;

  > .search {
    align-self: center;
    display: grid;
    place-items: center;

    > input {
      padding: 5px 10px;
      width: 500px;
      border-radius: 4px;
      outline: none;
    }
  }

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

  > .featured {
    > h3 {
      color: orangered;
      font-size: 1.4rem;
    }
  }
}
</style>
