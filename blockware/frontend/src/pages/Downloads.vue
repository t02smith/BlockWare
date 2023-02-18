<template>
  <div class="downloads">
    <div class="header">
      <h1>Downloads</h1>

      <div class="options">
        <button @click="() => (pauseAll = !pauseAll)">
          {{ pauseAll ? "▶️" : "⏸️" }}
        </button>
      </div>
    </div>

    <table>
      <thead>
        <th>Game</th>
        <th>Files Remaining</th>
        <th>Blocks Remaining</th>
      </thead>
      <tbody>
        <tr v-for="[hash, download] in downloadGamePairs">
          <td>{{ hash }}</td>
          <td>{{ Object.keys(download.Progress).length }}</td>
          <td>
            {{
              Object.values(download.Progress).reduce(
                (acc, f) => acc + f.BlocksRemaining.length,
                0
              )
            }}
          </td>
        </tr>
      </tbody>
    </table>
  </div>
</template>
<script setup>
import { computed, ref } from "vue";
import { useGamesStore } from "../stores/games";

const games = useGamesStore();

const downloadGamePairs = computed(() =>
  Object.keys(games.downloads).map((hash) => [hash, games.downloads[hash]])
);

const pauseAll = ref(false);
</script>
<style scoped lang="scss">
.downloads {
  width: 100%;
  display: flex;
  flex-direction: column;

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
    text-align: center;
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
