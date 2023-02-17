<template>
  <div class="login">
    <div class="title">
      <img src="../assets/images/icon.png" alt="" />
      <h1>BlockWare</h1>
    </div>

    <form @submit.prevent="submit">
      <input
        type="password"
        name=""
        id=""
        placeholder="Your ETH private key"
        v-model="key"
      />
      <button type="submit">Login</button>
    </form>
  </div>
</template>
<script setup>
import { ref } from "vue";
import { useRouter } from "vue-router";
import { DeployLibraryInstance } from "../../wailsjs/go/main/App";

const key = ref("");
const router = useRouter();

async function submit() {
  const e = await DeployLibraryInstance(key.value);
  if (e.length === 0) {
    key.value = "";
    router.push("/home");
  }
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

  > form {
    font-size: 1.25rem;
    border-radius: 10px;

    > * {
      outline: none;
      border: none;
    }

    > input {
      padding: 1rem;
      border-radius: 10px 0 0 10px;
      width: 300px;
    }

    > button {
      padding: 1rem 1rem;
      border-radius: 0 10px 10px 0;
      background-color: rgb(0, 130, 206);
      color: white;
      font-weight: bold;
      cursor: pointer;

      &:active {
        background-color: lighten(rgb(0, 130, 206), 10%);
      }
    }
  }
}
</style>
