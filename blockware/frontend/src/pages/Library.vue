<template>
  <div>
    <Library />
  </div>
</template>
<script setup>
import Library from "../components/library/Library.vue";
</script>
<style scoped lang="scss">
.title {
  background-color: #131313;
  display: flex;
  flex-direction: column;
  gap: 0;

  > h1 {
    font-size: 5rem;
  }

  > p {
    font-size: 2rem;
  }
}
</style>
