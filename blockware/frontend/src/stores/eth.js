import { defineStore } from "pinia";
import { ref } from "vue";
import {
  DeployLibraryInstance,
  JoinLibraryInstance,
} from "../../wailsjs/go/controller/Controller";

export const useEthStore = defineStore("eth", () => {
  const contractAddress = ref("");

  async function deployNewLibInstance(privKey) {
    contractAddress.value = await DeployLibraryInstance(privKey);
  }

  async function joinLibInstance(address, privKey) {
    await JoinLibraryInstance(address, privKey);
    contractAddress.value = address;
  }

  return { contractAddress, deployNewLibInstance, joinLibInstance };
});
