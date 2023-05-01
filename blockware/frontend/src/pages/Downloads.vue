<template>
  <div class="downloads">
    <div class="header">
      <h1>Downloads</h1>

      <div class="options">
        <!-- <button @click="() => (pauseAll = !pauseAll)">
          {{ pauseAll ? "▶️" : "⏸️" }}
        </button> -->
        <div>
          <p>Download stuck?</p>
          <button @click="games.loadDeferredRequests">Click here!</button>
        </div>
      </div>
    </div>

    <Error />

    <div class="download-table" v-if="downloadGamePairs.length > 0">
      <div class="table-header">
        <h2><strong>Game</strong></h2>
        <h2>Status</h2>
        <h2>Timer</h2>
        <h2>Block's Left</h2>
        <h2>Progress</h2>
      </div>

      <router-link
        :to="`/library?game=${hash}`"
        class="table-row"
        :class="`${
          download.Finished ? 'finished' : download.Paused ? 'paused' : ''
        }`"
        v-for="[hash, download] in downloadGamePairs"
      >
        <p class="indicator">
          {{ download.Finished ? "✅" : download.Paused ? "⏸️" : "⌛" }}
        </p>
        <p>
          {{ download.Name }} <strong>v{{ download.Version }}</strong>
        </p>
        <p>{{ download.Stage }}</p>

        <p>
          {{
            download.Stage === "Setting up download"
              ? "00:00:00"
              : download.ElapsedTime
          }}
        </p>

        <p>
          {{ download.BlocksLeft }}
        </p>

        <p>
          <strong>
            {{
              download.Stage === "Setting up download"
                ? 0
                : download.Stage === "Finished"
                ? 100
                : Math.round(
                    (1 - download.BlocksLeft / download.TotalBlocks) * 1000
                  ) / 10
            }}
            %</strong
          >
        </p>
      </router-link>
    </div>

    <div v-else>
      <h2>🥲 You have no downloads currently in progress</h2>
      <p>Head to your library to start a new download</p>
    </div>
  </div>
</template>
<script setup>
import { computed, ref, onMounted, onUnmounted } from "vue";
import Error from "../components/Error.vue";
import { useGamesStore } from "../stores/games";

/*

Show the user's list of completed and in progress downloads

*/

// hooks
const games = useGamesStore();

// interval function to trigger download progress request
const refreshInterval = ref(null);

// start interval for download refresh
onMounted(() => {
  games.refreshDownloads();
  refreshInterval.value = setInterval(() => games.refreshDownloads(), 250);
});

// stop interval for download refresh
onUnmounted(() => {
  if (!refreshInterval.value) return;
  clearInterval(refreshInterval.value);
});

// For each download <hash, download data>
// hash is used for identifying game
// download date shows the data left to be downloaded
const downloadGamePairs = computed(() => {
  if (games.downloads.length === 0) return [];
  return Object.keys(games.downloads)
    .filter((hash) => games.downloads[hash] !== null)
    .map((hash) => {
      let dl = games.downloads[hash];

      const start = new Date(dl.StartTime);
      const diff = new Date() - start;

      dl.timer = `${Math.floor(diff / 1000 / 60)}:${Math.floor(
        (diff / 1000) % 60
      )}`;

      return [hash, dl];
    });
});

// TODO
const pauseAll = ref(false);
</script>
<style scoped lang="scss">
.downloads {
  width: 100%;
  display: flex;
  flex-direction: column;
  max-width: 2000px;
  align-self: center;
  overflow-y: auto;

  > * {
    margin: 0rem 1rem;
  }

  > .header {
    margin: 1rem 3rem;
    display: flex;
    justify-content: center;
    align-items: center;

    > h1 {
      font-size: 3rem;
    }

    > .options {
      margin-left: auto;
      display: flex;
      align-items: center;
      gap: 1rem;

      > button {
        background-color: transparent;
        border: none;
        font-size: 3rem;
        cursor: pointer;
        transition: 75ms;

        &:hover {
          scale: 1.01;
        }

        &:active {
          scale: 0.99;
        }
      }

      > div {
        display: flex;
        flex-direction: column;
        justify-content: center;
        align-items: center;
        background-color: lighten(#131313, 12%);
        padding: 10px;
        border-radius: 3px;
        box-shadow: 0 4px 8px 0 rgba(0, 0, 0, 0.2),
          0 6px 20px 0 rgba(0, 0, 0, 0.19);
        gap: 5px;

        > p {
          font-style: italic;
          font-weight: bold;
          color: darken(white, 20%);
        }

        > button {
          width: fit-content;
          padding: 5px 10px;
          font-weight: bold;
          font-size: 1.1rem;
          background-color: red;
          color: white;
          border-radius: 3px;
          transition: 150ms;

          &:hover {
            background-color: darken(red, 5%);
          }

          &:active {
            background-color: darken(red, 10%);
          }
        }
      }
    }
  }
}

.download-table {
  margin: 1rem 4rem;

  > * {
    display: grid;
    grid-template-columns: 3fr 1fr 1fr 1fr 1fr;
    gap: 5px;
    text-align: right;
  }

  > .table-header {
    color: darken(white, 25%);
    background-color: lighten(#131313, 10%);
    padding: 0.4rem 0.85rem;
    border-radius: 3px;
    margin-bottom: 1rem;
    box-shadow: 0 4px 8px 0 rgba(0, 0, 0, 0.2), 0 6px 20px 0 rgba(0, 0, 0, 0.19);
  }

  > .table-row {
    padding: 0.25rem 1rem;
    text-decoration: none;
    border-radius: 3px;
    color: white;
    transition: 150ms;
    position: relative;
    margin: 10px 0;

    > .indicator {
      position: absolute;
      left: -2.2rem;
      top: -1px;
      font-size: 1.5rem;
    }

    &.finished {
      background-color: rgba(0, 252, 0, 0.185);

      &:hover {
        background-color: lighten(rgba(0, 252, 0, 0.185), 10%);
      }
    }

    &.paused {
      background-color: rgba(255, 8, 0, 0.185);

      &:hover {
        background-color: lighten(rgba(255, 8, 0, 0.185), 10%);
      }
    }

    &:hover {
      background-color: lighten(#131313, 5%);
      box-shadow: 0 4px 8px 0 rgba(0, 0, 0, 0.2),
        0 6px 20px 0 rgba(0, 0, 0, 0.19);
    }
  }

  p {
    font-size: 1.15rem;

    &:nth-child(2) {
      font-weight: bold;
      text-align: left;
    }
  }

  h2:first-child {
    text-align: left;
  }
}

table {
  align-self: center;
  width: 90%;
  border-collapse: collapse;

  > thead {
    background-color: rgb(0, 110, 255);
  }

  td,
  th {
    border: solid 1px gray;
  }

  td {
    padding: 1rem 0.5rem;
  }

  tr {
    cursor: pointer;

    &:nth-child(even) {
      background-color: #181818;
    }

    &:hover {
      background-color: #202020;
    }
  }
}
</style>
