<template>
  <div class="library">
    <!-- sidebar -->
    <ul>
      <li
        v-if="listGames"
        v-for="g in listGames"
        @click="
          () => {
            selected = g;
            importPanelOpen = false;
          }
        "
        :class="`${selected === g && 'active'} ${
          g.mostRecentVersion ? 'most-recent' : 'old-version'
        }`"
      >
        <img
          :src="`http://localhost:3003/${directory}/assets/${g.rootHash}/cover.png`"
          alt=""
        />
        <p>
          {{ g.title }}
        </p>
        <strong>v{{ g.version }}</strong>
      </li>

      <p class="empty" v-else>Nothing to show here 🥲</p>
    </ul>

    <!-- game details -->
    <div class="details-wrapper">
      <div class="nav">
        <p>🎮 Your games:</p>
        <button @click="() => (importPanelOpen = !importPanelOpen)">
          {{ importPanelOpen ? "❌ close" : "🌍 import" }}
        </button>
        <button @click="checkForUpdates" :disabled="checkingForUpdates">
          ♻️ Check for updates
        </button>
        <div class="checkbox">
          <input type="checkbox" name="" id="" v-model="showOldVersions" />
          <p>Show old versions</p>
        </div>
      </div>

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
import { ref, watch, onMounted, computed } from "vue";
import { useRoute, useRouter } from "vue-router";
import GameEntry from "../components/library/GameEntry.vue";
import {
  IsDownloading,
  UninstallGame,
  GetDirectory,
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

const directory = ref("");

// check for updates to your library
const checkingForUpdates = ref(false);
async function checkForUpdates() {
  checkForUpdates.value = true;
  await games.checkForUpdates();
  checkForUpdates.value = false;
}

// show old versions of game in list or not
const showOldVersions = ref(false);

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
  await games.refreshOwnedGames();
  directory.value = await GetDirectory();

  const gameHash = route.query.game;
  if (!gameHash) {
    if (games.ownedGames.length === 0) return;

    selected.value = games.ownedGames[0];
    return;
  }

  selected.value = games.ownedGames.find((g) => gameHash === g.rootHash);
});

const listGames = computed(() => {
  let gs = [];
  games.ownedGames
    .sort((a, b) => a.release > b.release)
    .forEach((g) => {
      if (gs.find((_g) => _g.title === g.title && _g.uploader === g.uploader)) {
        g.mostRecentVersion = false;
        if (showOldVersions.value) gs.push(g);
      } else {
        g.mostRecentVersion = true;
        gs.push(g);
      }
    });
  return gs;
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
    padding-top: 0.5rem;
    max-width: 1500px;
    justify-self: center;
    overflow-y: auto;
    position: relative;
    width: 100%;

    > .nav {
      margin: 0.5rem 1.5rem;
      display: flex;
      gap: 0.5rem;
      align-items: center;

      > p {
        font-weight: bold;
        font-size: 1.25rem;
      }

      > button:nth-child(2) {
        margin-left: auto;
      }

      > button {
        background-color: lighten(#131313, 10%);
        border: none;
        padding: 5px 8px;
        font-weight: bold;
        border-radius: 3px;
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

      > .checkbox {
        background-color: lighten(#131313, 10%);
        display: flex;
        gap: 5px;
        align-items: center;
        justify-content: center;
        padding: 5px 8px;
        font-weight: bold;
        border-radius: 3px;
        color: rgb(0, 174, 255);
        font-size: 0.7rem;

        > input {
          margin-top: 2px;
        }
      }
    }
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

    > li {
      cursor: pointer;
      transition: 100ms;
      height: fit-content;
      font-weight: bold;
      transition: 150ms;
      display: flex;
      align-items: center;

      &.most-recent {
        font-size: 1.15rem;

        > p {
          padding: 10px 5px;
        }

        > img {
          width: 50px;
          height: 50px;
        }

        > strong {
          font-size: 0.95rem;
          padding: 10px 5px;
        }
      }

      &.old-version {
        font-size: 0.75rem;

        > p {
          padding: 0px 5px;
        }

        > img {
          width: 25px;
          height: 25px;
        }

        > strong {
          padding: 0px 5px;
          font-size: 0.75rem;

          color: rgb(209, 17, 17);
        }
      }

      > img {
        object-fit: cover;
      }

      > strong {
        margin-left: auto;
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
