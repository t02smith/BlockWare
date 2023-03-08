<template>
  <div class="peers">
    <div class="title-card">
      <div class="text">
        <h1>Connect to New Peers</h1>
        <p>
          You and your peers will form a distributed network who will share game
          data with each other.
        </p>
      </div>

      <form class="new" @submit.prevent="connect">
        <p>tcp://</p>
        <input
          type="text"
          name=""
          id=""
          placeholder="hostname"
          v-model="hostname"
        />
        <p>:</p>
        <input type="number" name="" id="" placeholder="port" v-model="port" />

        <button type="submit">connect</button>
      </form>
    </div>

    <!-- peer list -->
    <div class="content">
      <h2>
        <strong>{{ peers.peers ? peers.peers.length : 0 }}</strong> - Your Peers
      </h2>

      <p v-if="!peers.peers">Nothing to show here 🥲</p>

      <ul v-else>
        <li v-for="p in peers.peers">
          <strong>tcp://{{ p.Hostname }}:{{ p.Port }}</strong> -
          {{ p.Library ? p.Library.length : 0 }} games in common
        </li>
      </ul>
    </div>
  </div>
</template>
<script setup>
import { ref } from "vue";
import { usePeerStore } from "../stores/peers";

const peers = usePeerStore();

// connect form
const hostname = ref("");
const port = ref(null);

async function connect() {
  if (hostname.value.length === 0 || !port.value) return;

  peers.connect(hostname.value, port.value);
  hostname.value = "";
  port.value = null;
}
</script>
<style scoped lang="scss">
.peers {
  display: flex;
  flex-direction: column;
  gap: 1rem;
  margin: 2rem;

  > .content {
    ul {
      margin-left: 1.5rem;
    }
  }

  > .title-card {
    display: flex;
    flex-direction: column;
    gap: 1rem;

    > form {
      display: flex;
      background-color: gray;
      width: fit-content;
      padding: 0.5rem;
      border-radius: 5px;

      > p {
        background-color: darken(white, 10%);
        padding: 0.5rem 0.5rem;
        font-size: 1rem;
        color: black;

        &:first-child {
          border-radius: 5px 0 0 5px;
        }
      }

      > input {
        font-size: 1rem;
        padding: 0.5rem;
        outline: none;
        border: none;

        &[type="number"] {
          width: 50px;
        }
      }

      > button {
        cursor: pointer;
        font-size: 1rem;
        padding: 0.5rem;
        border-radius: 0 5px 5px 0;
        background-color: rgb(56, 121, 219);
        color: white;
        font-weight: bold;
        outline: none;
        border: none;
      }
    }
  }
}
</style>
