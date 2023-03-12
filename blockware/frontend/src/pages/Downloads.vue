<template>
  <div class="downloads">
    <div class="header">
      <h1>Downloads</h1>

      <div class="options" v-if="games.downloads.length > 0">
        <button @click="() => (pauseAll = !pauseAll)">
          {{ pauseAll ? "▶️" : "⏸️" }}
        </button>
      </div>
    </div>

    <div class="download-table" v-if="downloadGamePairs.length > 0">
      <div class="table-header">
        <h2><strong>Game</strong></h2>
        <h2>File's Left</h2>
        <h2>Block's Left</h2>
        <h2>Progress</h2>
      </div>

      <router-link
        :to="`/library?game=${hash}`"
        class="table-row"
        :class="blocksLeft(download) === 0 && 'complete'"
        v-for="[hash, download] in downloadGamePairs"
      >
        <p>{{ download.Name }}</p>
        <p>{{ filesLeft(download) }}</p>
        <p>
          {{ blocksLeft(download) }}
        </p>
        <p>
          <strong>
            {{
              Math.round(
                (1 - blocksLeft(download) / download.TotalBlocks) * 1000
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
import { computed, ref } from "vue";
import { useGamesStore } from "../stores/games";

const games = useGamesStore();

const downloadGamePairs = computed(() => {
  if (games.downloads.length === 0) return [];
  return Object.keys(games.downloads)
    .filter((hash) => games.downloads[hash] !== null)
    .map((hash) => [hash, games.downloads[hash]]);
});

function blocksLeft(download) {
  if (!download) return 0;
  return Object.values(download.Progress).reduce(
    (acc, f) => acc + f.BlocksRemaining.length,
    0
  );
}

function filesLeft(download) {
  if (!download) return 0;
  return Object.values(download.Progress).filter(
    (f) => f.BlocksRemaining.length !== 0
  ).length;
}

const pauseAll = ref(false);
</script>
<style scoped lang="scss">
.downloads {
  width: 100%;
  display: flex;
  flex-direction: column;
  max-width: 2000px;
  align-self: center;

  > * {
    margin: 0rem 1rem;
  }

  > .header {
    margin: 1rem 3rem;
    display: flex;
    justify-content: center;

    > h1 {
      align-self: flex-start;
    }

    > .options {
      margin-left: auto;
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
    }
  }
}

.download-table {
  margin: 1rem 4rem;

  > * {
    display: grid;
    grid-template-columns: 3fr 1fr 1fr 1fr;
  }

  > .table-header {
    color: darken(white, 25%);
    background-color: lighten(#131313, 10%);
    padding: 0.4rem 0.85rem;
    border-radius: 10px;
    margin-bottom: 1rem;
    box-shadow: 0 4px 8px 0 rgba(0, 0, 0, 0.2), 0 6px 20px 0 rgba(0, 0, 0, 0.19);
  }

  > .table-row {
    padding: 0.25rem 1rem;
    text-decoration: none;
    border-radius: 10px;
    color: white;
    transition: 150ms;

    &.complete {
      background-color: rgba(0, 252, 0, 0.185);

      &:hover {
        background-color: lighten(rgba(0, 252, 0, 0.185), 10%);
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
