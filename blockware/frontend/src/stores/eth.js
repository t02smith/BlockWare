import { defineStore } from "pinia";
import { ref } from "vue";
import {
  DeployLibraryInstance,
  JoinLibraryInstance,
  GetContractAddress,
} from "../../wailsjs/go/controller/Controller";

/**
 * Interact and form a connection with the Eth smart contract
 */
export const useEthStore = defineStore("eth", () => {
  // the address of the current library smart contract
  const contractAddress = ref("");

  const connected = ref(false);

  /**
   * Deploy a new library smart contract
   * ! Not used
   * @param {String} privKey
   */
  async function deployNewLibInstance(privKey) {
    contractAddress.value = await DeployLibraryInstance(privKey);
  }

  /**
   * Join an existing library smart contract
   * @param {String} address the address of the contract
   * @param {String} privKey
   */
  async function joinLibInstance(address, privKey) {
    await JoinLibraryInstance(address, privKey);
    contractAddress.value = address;
    connected.value = true;
  }

  async function getContractAddress() {
    return await GetContractAddress();
  }

  return {
    contractAddress,
    deployNewLibInstance,
    joinLibInstance,
    getContractAddress,
    connected,
  };
});
