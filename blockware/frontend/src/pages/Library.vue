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
    <div class="details-wrapper">
      <GameEntry :game="selected" v-if="selected">
        <button
          @click="createDownload"
          v-if="selectedIsDownloading === 0"
          class="download new"
        >
          <h3>💡 Download</h3>
        </button>

        <div
          v-else-if="selectedIsDownloading === 1"
          class="download finished"
          @click="uninstall"
        >
          <h3>🗑️ Uninstall</h3>
        </div>

        <router-link
          :to="`/downloads`"
          v-else-if="selectedIsDownloading === 2"
          class="download downloading"
        >
          <h3>⏸️ Pause</h3>
        </router-link>
      </GameEntry>
    </div>
  </div>
</template>
<script setup>
import { ref, watch, onMounted } from "vue";
import { useRoute, useRouter } from "vue-router";
import GameEntry from "../components/store/GameEntry.vue";
import {
  IsDownloading,
  UninstallGame,
} from "../../wailsjs/go/controller/Controller";
import { useGamesStore } from "../stores/games";

const props = defineProps({
  store: {
    type: Boolean,
    default: false,
  },
});

// setup
const games = useGamesStore();
const router = useRouter();
const route = useRoute();

// The game currently being viewed
const selected = ref(null);
const selectedIsDownloading = ref(0);

watch(selected, async () => {
  if (!selected.value) return;

  // update route
  router.replace({
    path: route.path,
    query: { game: selected.value.rootHash },
  });

  // check download status
  selectedIsDownloading.value = await IsDownloading(selected.value.rootHash);
});

onMounted(async () => {
  const gameHash = route.query.game;
  if (!gameHash) {
    if (games.ownedGames.length === 0) return;

    selected.value = games.ownedGames[0];
    return;
  }

  selected.value = games.ownedGames.find((g) => gameHash === g.rootHash);
});

function createDownload() {
  games.createDownload(selected.value.rootHash);
  selectedIsDownloading.value = 2;
}

async function uninstall() {
  await UninstallGame(selected.value.rootHash);
  selectedIsDownloading.value = await IsDownloading(selected.value.rootHash);
}
</script>
<style scoped lang="scss">
.library {
  position: relative;
  display: grid;
  grid-template-columns: 1fr 5fr;
  gap: 1rem;
  height: 100%;
  flex-grow: 1;
  overflow: hidden;

  .details-wrapper {
    padding: 1.5rem;
    max-width: 1500px;
    justify-self: center;
    overflow-y: auto;
  }

  ul {
    list-style: none;
    background-color: lighten(#131313, 5%);
    box-shadow: 0 4px 8px 0 rgba(0, 0, 0, 0.2), 0 6px 20px 0 rgba(0, 0, 0, 0.19);
    border-radius: 0 10px 0;
    border: solid 2px rgba(0, 132, 255, 0.24);
    border-bottom: none;
    border-left: none;
    max-width: 250px;
    overflow-y: auto;

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

.download {
  padding: 0.5rem 1rem;
  font-weight: bold;
  color: white;
  border-radius: 5px;
  box-shadow: 0 4px 8px 0 rgba(0, 0, 0, 0.2), 0 6px 20px 0 rgba(0, 0, 0, 0.19);
  cursor: pointer;
  transform-origin: 150ms;
  text-decoration: none;
  font-size: 0.9rem;

  &:hover {
    scale: 1.02;
  }

  &:active {
    scale: 0.99;
  }

  &.new {
    background-color: rgb(1, 129, 189);
  }

  &.downloading {
    background-color: orangered;
  }

  &.finished {
    background-color: green;
  }
}
</style>
