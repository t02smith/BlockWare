import { defineStore } from "pinia";
import { ref } from "vue";

export const useErrStore = defineStore("err", () => {
  const err = ref("");

  const clear = () => (err.value = "");
  const set = (newErr) => (err.value = newErr);

  return {
    err,
    clear,
    set,
  };
});
