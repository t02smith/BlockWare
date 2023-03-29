<template>
  <div class="err" v-if="err.err">
    <strong>❌ ERROR</strong> {{ err.err }}
    <p class="clear" @click="err.clear">close</p>
  </div>
</template>
<script setup>
import { useErrStore } from "../stores/err";

const err = useErrStore();
</script>
<style scoped lang="scss">
.err {
  display: flex;
  align-items: center;
  gap: 0.25rem;
  place-self: center;
  margin: 1rem;
  font-size: 1.1rem;
  padding: 0.5rem 0.5rem;
  border-radius: 4px;
  background-color: lighten(red, 30%);
  border: solid 2px lighten(red, 20%);
  color: black;

  box-shadow: 0 4px 8px 0 rgba(0, 0, 0, 0.2), 0 6px 20px 0 rgba(0, 0, 0, 0.19);
  width: 100%;
  max-width: 80vw;

  > strong {
    color: rgb(255, 0, 47);
  }

  > .clear {
    margin-left: auto;
    cursor: pointer;
    color: rgba(123, 16, 16, 0.568);

    &:hover {
      text-decoration: underline;
    }
  }
}
</style>
