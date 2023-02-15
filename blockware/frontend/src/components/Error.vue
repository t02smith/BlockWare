<template>
  <div class="err"><strong>❌ Error</strong> {{ props.err }}</div>
</template>
<script setup>
const props = defineProps({
  err: {
    type: String,
    required: true,
  },
});
</script>
<style scoped lang="scss"></style>
