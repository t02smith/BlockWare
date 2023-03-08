<template>
  <div class="login">
    <div class="title">
      <img src="../assets/images/icon.png" alt="" />
      <h1>BlockWare</h1>
    </div>

    <div class="form">
      <form @submit.prevent="join" v-if="isJoining">
        <div class="form-group">
          <p>The address of your chosen BlockWare instance</p>
          <input
            type="text"
            name=""
            id=""
            placeholder="Contract Address"
            v-model="addr"
          />
        </div>

        <div class="form-group">
          <p>Your ETH private key for uploading and buying</p>
          <input
            type="password"
            name=""
            id=""
            placeholder="Private Key"
            v-model="key"
          />
        </div>

        <button type="submit">Join</button>
      </form>

      <form @submit.prevent="create" v-else>
        <div class="form-group">
          <p>Your ETH private key for deploying</p>
          <input
            type="password"
            name=""
            id=""
            placeholder="Private Key"
            v-model="key"
          />
        </div>

        <button type="submit">Deploy</button>
      </form>

      <p @click="() => (isJoining = !isJoining)">
        Click here to
        {{
          isJoining ? "deploy your own instance" : "join an existing instance"
        }}
      </p>
    </div>
  </div>
</template>
<script setup>
import { ref } from "vue";
import { useRouter } from "vue-router";
import { useEthStore } from "../stores/eth";

const key = ref(
  "af9668cd6ebc3ba4c0e5036c284e128ed66e18ba9e4ed87b2c0c6d9642f2b879"
);
const addr = ref("0x750cf6392175f94ff5014803a0bb6b79de543337");
const isJoining = ref(true);

const eth = useEthStore();
const router = useRouter();

async function create() {
  await eth.deployNewLibInstance(key.value);
  key.value = "";
  router.push("/home");
}

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

  > .form {
    display: flex;
    flex-direction: column;
    align-items: center;

    > form {
      display: flex;
      flex-direction: column;

      > .form-group {
        margin-bottom: 1rem;

        > p {
          color: darken(white, 20%);
          font-style: italic;
        }

        > input {
          width: 100%;
          padding: 5px 6px;
          border-radius: 4px;
          outline: none;
          min-width: 400px;
        }
      }

      > button {
        background-color: rgb(18, 141, 18);
        font-weight: bold;
        width: fit-content;
        place-self: center;
        padding: 4px 30px;
        border-radius: 4px;
        color: white;
        margin-top: 1rem;
        cursor: pointer;
      }
    }

    > p {
      font-size: 0.9rem;
      color: darken(white, 20%);
      font-style: italic;
      cursor: pointer;

      &:hover {
        text-decoration: underline;
      }
    }
  }
}
</style>
