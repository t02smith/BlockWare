<template>
  <div class="library">
    <!-- sidebar -->
    <ul>
      <p>🎮 Your games:</p>
      <li
        v-for="g in games.ownedGames"
        @click="() => (selected = g)"
        :class="`${selected === g && 'active'}`"
      >
        <p>{{ g.title }}</p>
      </li>
    </ul>

    <!-- game details -->
    <GameEntry :game="selected" v-if="selected" />
  </div>
</template>
<script setup>
import { ref, watch, onMounted } from "vue";
import { useRoute, useRouter } from "vue-router";
import GameEntry from "../components/store/GameEntry.vue";
import { useGamesStore } from "../stores/games";
import { toHexString } from "../util/util";

const props = defineProps({
  store: {
    type: Boolean,
    default: false,
  },
});

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
  if (!gameHash) {
    if (games.ownedGames.length === 0) return;

    selected.value = games.ownedGames[0];
    return;
  }

  selected.value = games.ownedGames.find(
    (g) => gameHash === toHexString(g.rootHash)
  );
});
</script>
<style scoped lang="scss">
.library {
  position: relative;
  display: grid;
  grid-template-columns: 1fr 5fr;
  gap: 1rem;
  height: 100%;
  overflow-x: hidden;

  ul {
    list-style: none;
    background-color: lighten(#131313, 5%);
    box-shadow: 0 4px 8px 0 rgba(0, 0, 0, 0.2), 0 6px 20px 0 rgba(0, 0, 0, 0.19);
    border-radius: 0 10px 0;
    border: solid 2px rgba(0, 132, 255, 0.24);
    border-bottom: none;
    border-left: none;

    > p {
      margin: 0.4rem 0.5rem;
      font-weight: bold;
      text-align: center;
    }

    > li {
      cursor: pointer;
      transition: 100ms;
      padding: 0.5rem 0.75rem;
      font-weight: bold;
      font-size: 1.15rem;
      border-bottom: 1px solid rgb(85, 85, 85);
      transition: 150ms;

      &:last-child {
        border-color: transparent;
      }

      &:hover {
        background-color: lighten(#131313, 10%);
      }

      &.active {
        background-color: rgba(0, 132, 255, 0.24);
      }
    }
  }
}
</style>
