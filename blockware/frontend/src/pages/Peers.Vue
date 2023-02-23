<template>
  <div class="peers">
    <!-- peer list -->
    <div class="content">
      <h2>Your Peers</h2>

      <p v-if="!peers">Loading...</p>

      <p v-else v-for="p in peers">{{ p }}</p>
    </div>

    <form class="new" @submit.prevent="connect">
      <input
        type="text"
        name=""
        id=""
        placeholder="hostname"
        v-model="hostname"
      />
      <input type="number" name="" id="" placeholder="port" v-model="port" />

      <button type="submit">Connect</button>
    </form>
  </div>
</template>
<script setup>
import { onMounted, ref } from "vue";
import {
  GetPeerInformation,
  ConnectToPeer,
} from "../../wailsjs/go/controller/Controller";

const peers = ref([]);

onMounted(async () => {
  peers.value = await GetPeerInformation();
});

// connect form
const hostname = ref("");
const port = ref(null);

async function connect() {
  if (hostname.value.length === 0 || !port.value) return;

  await ConnectToPeer(hostname.value, port.value);
}
</script>
<style scoped lang="scss">
.peers {
  display: grid;
  grid-template-columns: 1fr 1fr;

  > * {
    margin: 2rem;
  }
}
</style>
