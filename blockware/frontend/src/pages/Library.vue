<template>
  <div class="library">
    <!-- sidebar -->
    <ul>
      <div class="header">
        <p>🎮 Your games:</p>
        <button @click="() => (importPanelOpen = !importPanelOpen)">
          {{ importPanelOpen ? "❌ close" : "🌍 import" }}
        </button>
      </div>

      <li
        v-if="games.ownedGames"
        v-for="g in games.ownedGames"
        @click="
          () => {
            selected = g;
            importPanelOpen = false;
          }
        "
        :class="`${selected === g && 'active'}`"
      >
        <p>
          {{ g.title }}
        </p>
        <strong>{{ g.version }}</strong>
      </li>

      <p class="empty" v-else>Nothing to show here 🥲</p>
    </ul>

    <!-- game details -->
    <div class="details-wrapper">
      <form
        @submit.prevent="
          () => {
            if (importGameHash.length !== 64) return;
            games.importGame(importGameHash);
            importGameHash = '';
            importPanelOpen = false;
          }
        "
        class="import-game"
        :class="importPanelOpen ? 'open' : 'shut'"
      >
        <div class="text">
          <h5>Import one of your owned games:</h5>
          <p>
            Enter the root hash of one of your owned games to add it to your
            local library,
          </p>
        </div>

        <div class="input">
          <input
            type="text"
            name=""
            id=""
            placeholder="The game's hash"
            v-model="importGameHash"
          />
          <button type="submit">import</button>
        </div>
      </form>

      <Error />

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
import GameEntry from "../components/library/GameEntry.vue";
import {
  IsDownloading,
  UninstallGame,
} from "../../wailsjs/go/controller/Controller";
import { useGamesStore } from "../stores/games";
import Error from "../components/Error.vue";

/*

Show the user's owned games and allow them to manage downloads

*/

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

// import new game
const importPanelOpen = ref(false);
const importGameHash = ref("");

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

// load game mentioned in query parameter if it exists
onMounted(async () => {
  const gameHash = route.query.game;
  if (!gameHash) {
    if (games.ownedGames.length === 0) return;

    selected.value = games.ownedGames[0];
    return;
  }

  selected.value = games.ownedGames.find((g) => gameHash === g.rootHash);
});

/*
start a download for an owned game
*/
function createDownload() {
  games.createDownload(selected.value.rootHash);
  selectedIsDownloading.value = 2;
}

/*
uninstall a downloaded game
*/
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
    position: relative;
    width: 100%;
  }

  .import-game {
    z-index: 100;
    top: 1.2rem;
    position: absolute;
    background-color: lighten(#131313, 7%);
    box-shadow: 0 4px 8px 0 rgba(0, 0, 0, 0.2), 0 6px 20px 0 rgba(0, 0, 0, 0.19);
    border-radius: 5px;
    padding: 0.75rem 1rem;
    max-width: 325px;
    display: flex;
    flex-direction: column;
    gap: 1rem;
    transition: 250ms;

    &.open {
      left: 1.5rem;
    }

    &.shut {
      left: -25rem;
    }

    > .text {
      > p {
        font-style: italic;
        font-size: 0.85rem;
        color: darken(white, 12%);
      }
    }

    > .input {
      display: flex;
      align-items: center;

      input {
        border-radius: 1px 0 0 1px;
        padding: 2.5px 3px;
        border: none;
        height: 25px;
        width: 250px;
        background-color: darken(white, 5%);
      }

      button {
        border-radius: 0 1px 1px 0;
        border: none;
        height: 30px;
        width: 50px;
        background-color: rgb(0, 131, 253);
        font-weight: bold;
        color: white;
      }
    }
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
    min-width: 230px;

    > .empty {
      color: darken(white, 20%);
      text-align: left;
    }

    > .header {
      display: flex;
      padding: 12px 15px;
      align-items: center;

      > p {
        font-weight: bold;
      }

      > button {
        margin-left: auto;
        background-color: lighten(#131313, 20%);
        border: none;
        padding: 5px 8px;
        font-weight: bold;
        border-radius: 6px;
        margin-top: 4px;
        cursor: pointer;
        transition: 150ms;
        color: rgb(0, 174, 255);
        font-size: 0.7rem;

        &:hover {
          background-color: lighten(#131313, 25%);
        }

        &.active {
          background-color: lighten(#131313, 20%);
        }
      }
    }

    > li {
      cursor: pointer;
      transition: 100ms;
      padding: 0.5rem 0.75rem;
      font-weight: bold;
      font-size: 1.15rem;
      border-bottom: 1px solid rgb(85, 85, 85);
      transition: 150ms;
      display: flex;
      align-items: center;

      > strong {
        margin-left: auto;
        font-size: 0.95rem;
      }

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
