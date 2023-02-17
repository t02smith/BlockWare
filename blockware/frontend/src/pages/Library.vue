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

      <p>
        Lorem ipsum dolor, sit amet consectetur adipisicing elit. Deleniti
        dignissimos aliquid quod, quae perspiciatis incidunt, aperiam a illum
        facilis voluptate provident consequuntur iste fugit voluptatem? Ducimus
        et similique eius, excepturi quo maiores! Incidunt saepe magnam
        laudantium earum sed dolores dolor natus beatae, tempore labore sunt
        exercitationem quidem nam quis non.
      </p>
    </div>
  </div>
</template>
<script setup>
import { ref, watch, onMounted } from "vue";
import { useRoute, useRouter } from "vue-router";
import { useGamesStore } from "../stores/games";
import { toHexString } from "../util/util";

const games = useGamesStore();
const router = useRouter();
const route = useRoute();

const selected = ref(null);
watch(selected, () => {
  if (!selected.value) return;

  router.replace({
    path: route.path,
    query: { game: toHexString(selected.value.rootHash) },
  });
});

onMounted(() => {
  const gameHash = route.query.game;
  if (!gameHash) return;

  selected.value = games.ownedGames.find(
    (g) => gameHash === toHexString(g.rootHash)
  );
});
</script>
<style scoped lang="scss">
.library {
  display: grid;
  grid-template-columns: 1fr 5fr;
  gap: 1rem;

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
    display: flex;
    flex-direction: column;
    gap: 1rem;

    > .header {
      display: flex;
      align-items: flex-end;
      gap: 1rem;

      > img {
        width: 200px;
        height: 200px;
        background-color: lighten(#131313, 5%);
        border-radius: 10px;
        padding: 2px;
        box-shadow: 0 4px 8px 0 rgba(0, 0, 0, 0.2),
          0 6px 20px 0 rgba(0, 0, 0, 0.19);
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

    > p {
      color: darken(white, 20%);
    }
  }
}
</style>
