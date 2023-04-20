<template>
  <div class="login">
    <div class="title">
      <h1><strong>Block</strong>Ware</h1>
      <h2>by Tom Smith</h2>
    </div>

    <h3>
      🌍 Connecting to Library <strong>{{ addr }}</strong>
    </h3>

    <div class="form">
      <form @submit.prevent="join" v-if="isJoining">
        <input
          type="password"
          name=""
          id=""
          placeholder="🔑Your ETH Private Key"
          v-model="key"
        />

        <button type="submit">Go</button>
      </form>

      <form @submit.prevent="create" v-else>
        <input
          type="password"
          name=""
          id=""
          placeholder="🔑Your ETH Private Key"
          v-model="key"
        />
        <button type="submit">Deploy</button>
      </form>

      <p>
        Need some help?
        <router-link to="/help"><strong>Click Here!</strong></router-link>
      </p>
    </div>
  </div>
</template>
<script setup>
import { ref, onMounted } from "vue";
import { useRouter } from "vue-router";
import { useEthStore } from "../stores/eth";

/*

Connect to the BlockWare network

*/

// user's private key
const key = ref("");

// address of the deployed smart contract
const addr = ref("");

// join existing network OR deploy new one
const isJoining = ref(true);

// hooks
const eth = useEthStore();
const router = useRouter();

onMounted(async () => {
  addr.value = await eth.getContractAddress();
});

/*
Deploy a new instance
*/
async function create() {
  await eth.deployNewLibInstance(key.value);
  key.value = "";
  router.push("/home");
}

/*
Join an existing instance
*/
async function join() {
  await eth.joinLibInstance(addr.value, key.value);
  addr.value = "";
  router.push("/home");
}
</script>
<style scoped lang="scss">
.login {
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  gap: 2rem;
  height: 100vh;

  > h3 {
    margin-top: 5rem;
    background-color: #303030;
    padding: 3px 15px;
    border-radius: 10px;
    font-size: 1.15rem;
  }

  > .title {
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: center;

    > h1 {
      font-size: 12rem;

      > strong {
        background: url("../assets/images/icon.png");
        background-attachment: fixed;
        -webkit-background-clip: text;
        background-clip: text;
        -webkit-text-fill-color: transparent;
      }
    }

    > h2 {
      margin-top: -1.75rem;
      color: darken(white, 20%);
    }
  }

  > .form {
    display: flex;
    flex-direction: column;
    align-items: center;

    > form {
      display: flex;
      align-items: center;
      justify-content: center;
      background-color: lighten(#131313, 10%);
      padding: 0.5rem;
      border-radius: 5px;

      > input {
        padding: 10px;
        min-width: 350px;
        height: 20px;
        border-radius: 5px 0 0 5px;
        border: none;
      }

      > button {
        background-color: rgb(21, 64, 255);
        font-weight: bold;
        width: fit-content;
        padding: 10px;
        border: none;
        border-radius: 0 5px 5px 0;
        color: white;
        height: 40px;
        font-size: 1.2rem;
        cursor: pointer;
        transition: 150ms;

        &:hover {
          background-color: lighten(rgb(21, 64, 255), 4%);
        }
      }
    }

    > p {
      margin-top: 1rem;
      font-size: 0.9rem;
      color: darken(white, 20%);
      font-style: italic;
      cursor: pointer;

      a {
        text-decoration: none;
        color: white;

        &:hover {
          text-decoration: underline;
        }
      }
    }
  }
}
</style>
