<template>
  <div class="library">
    <!-- sidebar -->
    <ul>
      <li v-for="g in games.ownedGames" @click="() => (selected = g)">
        <p>{{ g.title }}</p>
      </li>
    </ul>

    <!-- game details -->
    <div class="details" v-if="selected">
      <div class="header">
        <img src="../assets/images/logo-universal.png" alt="" />
        <div class="header-text">
          <h2>{{ selected.title }}</h2>
          <h3>{{ selected.dev }}</h3>
        </div>
      </div>
    </div>
  </div>
</template>
<script setup>
import { ref } from "vue";
import { useGamesStore } from "../stores/games";

const games = useGamesStore();
const selected = ref(null);
</script>
<style scoped lang="scss">
.library {
  display: grid;
  grid-template-columns: 1fr 5fr;

  > ul {
    list-style: none;
    padding: 0.6rem 1rem;
    background-color: lighten(#131313, 5%);

    > li {
      cursor: pointer;
      transition: 100ms;
      padding: 3px 4px;

      &:hover {
        background-color: lighten(#131313, 10%);
      }
    }
  }

  .details {
    margin: 1rem;
    > .header {
      display: flex;
      align-items: flex-end;

      > img {
        width: 200px;
        height: 200px;
      }

      > .header-text {
        > h2 {
          font-size: 4rem;
        }

        > h3 {
          color: darken(white, 30%);
        }
      }
    }
  }
}
</style>
