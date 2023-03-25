import { defineStore } from "pinia";
import { ref } from "vue";
import {
  DeployLibraryInstance,
  JoinLibraryInstance,
} from "../../wailsjs/go/controller/Controller";

/**
 * Interact and form a connection with the Eth smart contract
 */
export const useEthStore = defineStore("eth", () => {
  // the address of the current library smart contract
  const contractAddress = ref("");

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
  }

  return { contractAddress, deployNewLibInstance, joinLibInstance };
});
