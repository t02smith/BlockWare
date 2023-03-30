<template>
  <nav>
    <router-link to="/home" class="title">
      <img src="../assets/images/icon.png" alt="" />
      <h4>BlockWare</h4>

      <h3 class="address" v-if="eth.contractAddress">
        🌍 Using library {{ eth.contractAddress }}
      </h3>
    </router-link>

    <div class="links">
      <router-link v-if="eth.connected" to="/library">🕹️ Library</router-link>
      <router-link v-if="eth.connected" to="/downloads"
        >🎁 Downloads</router-link
      >
      <router-link v-if="eth.connected" to="/upload">🆕 Upload</router-link>
      <router-link v-if="eth.connected" to="/store">💵 Store</router-link>
      <router-link v-if="eth.connected" to="/peers">🫂 Peers</router-link>

      <router-link v-if="!eth.connected" to="/">🔒 Connect</router-link>
      <router-link to="/help">❓ Help</router-link>
    </div>
  </nav>
</template>
<script setup>
import { useEthStore } from "../stores/eth";

/*
Standard navbar with links to all major pages
*/

// used to get the contract address
const eth = useEthStore();
</script>
<style scoped lang="scss">
nav {
  max-width: 100vw;
  display: flex;
  align-items: center;
  justify-content: center;

  > * {
    padding: 1rem;
  }

  a {
    text-decoration: none;
  }

  > .title {
    font-size: 1.5rem;
    display: flex;
    gap: 0.4rem;
    align-items: center;
    color: white;

    > img {
      width: 30px;
      height: 30px;
    }

    > .address {
      display: flex;
      gap: 0.35rem;
      align-items: center;
      background-color: #303030;
      padding: 3px 8px;
      border-radius: 10px;
      font-size: 0.9rem;
      overflow: hidden;
      height: 25px;
      cursor: pointer;
      width: auto;
      margin-top: 4px;
      color: darken(white, 15%);
    }
  }

  > .links {
    margin-left: auto;
    display: flex;
    gap: 0.5rem;

    a {
      color: rgb(181, 181, 181);
    }

    a:hover {
      color: white;
    }
  }
}
</style>
