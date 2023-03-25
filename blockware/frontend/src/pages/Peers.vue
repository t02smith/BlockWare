<template>
  <div class="peers">
    <div class="side-panel">
      <h3>
        Connect to New Peers:
        <hr style="opacity: 0.5" />
      </h3>

      <form class="new" @submit.prevent="connect">
        <h4>Connect to a <strong>Single Peer</strong>:</h4>

        <input
          type="text"
          name=""
          id=""
          placeholder="hostname (localhost)"
          v-model="hostname"
        />

        <input
          type="number"
          name=""
          id=""
          placeholder="port (6750)"
          v-model="port"
        />

        <button type="submit">connect</button>
      </form>

      <hr style="opacity: 0.5" />
      <form @submit.prevent="connectAll" class="many-peers">
        <h4>Connect to <strong>Many Peers</strong>:</h4>

        <textarea
          name=""
          id=""
          cols="30"
          rows="10"
          v-model="manyPeers"
          placeholder="localhost:6750
localhost:6751
..."
        ></textarea>

        <button type="submit">Connect</button>
      </form>
    </div>

    <div class="content">
      <div class="title">
        <h2>Your Peers</h2>
        <p>Below are the list of peers you're currently connected to.</p>
        <p>
          Peers are used to share game data and contribute to the distribution
          of games uploaded to the <strong>BlockWare</strong> network.
        </p>
        <br />
        <p>🔒 = We have verified the peer's Eth address</p>
        <p>
          ⛔ = We have't received verification.
          <strong> Click it to resend.</strong>
        </p>
      </div>

      <div class="peer-list">
        <h3>
          You are connected to
          <strong>{{ peers.peers ? peers.peers.length : 0 }} peers</strong>!
        </h3>

        <p v-if="!peers.peers">Nothing to show here 🥲</p>

        <div v-else v-for="p in peers.peers.slice(0, 15)" class="peer">
          <button
            class="disconnect"
            @click="() => disconnect(p.Hostname, p.Port)"
          >
            ❌
          </button>

          <p v-if="p.Validated">🔒</p>
          <button
            class="validation"
            @click="() => peers.resendValidation(p.Hostname, p.Port)"
            v-else
          >
            ⛔
          </button>

          <p>
            <strong>tcp://{{ p.Hostname }}:{{ p.Port }}</strong> -
            {{ p.Library ? p.Library.length : 0 }} games in common
          </p>
        </div>
        <p
          v-if="peers.peers && peers.peers.length > 15"
          style="font-style: italic"
        >
          {{ peers.peers.length - 15 }} more peers connected...
        </p>
      </div>
    </div>
  </div>
</template>
<script setup>
import { ref } from "vue";
import { usePeerStore } from "../stores/peers";

/*

Allow users to connect to and manage their connections with
peers.

*/

// Peers pinia store
const peers = usePeerStore();

// connect to single peer
const hostname = ref("");
const port = ref(null);

/*
connect to the single peer 
*/
async function connect() {
  if (hostname.value.length === 0 || !port.value) return;

  peers.connect(hostname.value, port.value);
  hostname.value = "";
  port.value = null;
}

// conenct to many peers
const manyPeers = ref("");

/*
connect to all peers listed in the manyPeers field
*/
async function connectAll() {
  if (manyPeers.value.length === 0) return;

  peers.connectToAll(manyPeers.value);
  manyPeers.value = "";
}

/*
disconnect to an existing peer
*/
function disconnect(hostname, port) {
  peers.disconnect(hostname, port);
}
</script>
<style scoped lang="scss">
.peers {
  display: flex;
  gap: 1rem;
  height: 100%;

  > .side-panel {
    background-color: lighten(#131313, 5%);
    height: 100%;
    max-width: fit-content;
    padding: 1rem 1rem;
    box-shadow: 0 4px 8px 0 rgba(0, 0, 0, 0.2), 0 6px 20px 0 rgba(0, 0, 0, 0.19);
    border-radius: 0 10px 0;
    border: solid 2px rgba(0, 132, 255, 0.24);
    border-left: none;
    border-bottom: none;
    display: flex;
    flex-direction: column;
    gap: 1rem;
    min-width: 200px;

    > form {
      display: flex;
      flex-direction: column;
      gap: 0.5rem;

      > button {
        padding: 5px 15px;
        font-weight: bold;
        border-radius: 3px;
        background-color: rgb(0, 132, 255);
        color: white;
        border: none;
        width: fit-content;
        align-self: flex-end;
        cursor: pointer;
        transition: 150ms;

        &:hover {
          opacity: 0.75;
        }

        &:focus {
          background-color: rgb(20, 187, 20);
        }
      }
    }

    > .new {
      > input {
        padding: 5px 7px;
        border-radius: 2px;
        border: none;

        &:focus {
          font-weight: bold;
        }
      }

      > p {
        background-color: gray;
      }
    }

    > .many-peers {
      > textarea {
        border: none;
        padding: 5px 2px;
        width: 100%;
        resize: none;
        border-radius: 2px;

        &:focus {
          font-weight: bold;
        }
      }
    }
  }

  > .content {
    display: flex;
    flex-direction: column;
    gap: 1rem;

    > .title {
      max-width: min(100%, 900px);

      > h2 {
        font-size: 2rem;
      }

      > p {
        color: darken(white, 20%);
      }
    }

    > .peer-list {
      > .peer {
        display: flex;
        align-items: center;
        gap: 0.5rem;

        > .disconnect,
        .validation {
          background-color: transparent;
          font-size: 1.2rem;
          border: none;
          transition: 150ms;
          cursor: pointer;

          &:hover {
            opacity: 0.75;
          }
        }

        > p {
          font-size: 1.2rem;
        }
      }
    }
  }
}
</style>
