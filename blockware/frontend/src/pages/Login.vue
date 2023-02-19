<template>
  <div class="login">
    <div class="title">
      <img src="../assets/images/icon.png" alt="" />
      <h1>BlockWare</h1>
    </div>

    <div class="forms">
      <form @submit.prevent="join">
        <input
          type="text"
          name=""
          id=""
          placeholder="Contract address"
          v-model="addr"
        />
        <button type="submit">Join</button>
      </form>

      <form @submit.prevent="create">
        <p>Or</p>
        <input
          type="password"
          name=""
          id=""
          placeholder="Your ETH private key"
          v-model="key"
        />
        <button type="submit">Create</button>
      </form>
    </div>
  </div>
</template>
<script setup>
import { ref } from "vue";
import { useRouter } from "vue-router";
import { useEthStore } from "../stores/eth";

const key = ref("");
const addr = ref("");
const router = useRouter();
const eth = useEthStore();

async function create() {
  await eth.deployNewLibInstance(key.value);
  key.value = "";
  router.push("/home");
}

async function join() {
  await eth.joinLibInstance(addr.value);
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

  > .title {
    display: flex;
    align-items: center;
    gap: 1rem;

    > h1 {
      font-size: 5rem;
    }

    > img {
      width: 100px;
      height: 100px;
      object-fit: cover;
    }
  }

  > .forms {
    display: flex;
    flex-direction: column;
    gap: 0.5rem;

    > form {
      font-size: 1.25rem;
      border-radius: 10px;
      display: flex;
      align-items: center;

      > * {
        outline: none;
        border: none;
      }

      > p {
        color: darken(white, 20%);
        margin-right: 1rem;
      }

      > input {
        padding: 0.5rem 1rem;
        border-radius: 5px 0 0 5px;
        width: 300px;
      }

      > button {
        padding: 0.5rem 1rem;
        border-radius: 0 5px 5px 0;
        background-color: rgb(0, 130, 206);
        color: white;
        font-weight: bold;
        cursor: pointer;

        &:active {
          background-color: lighten(rgb(0, 130, 206), 10%);
        }
      }

      .contract {
        display: flex;
        justify-content: center;
        gap: 1.5rem;
      }
    }
  }
}
</style>
