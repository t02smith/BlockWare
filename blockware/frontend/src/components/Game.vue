<template>
  <div class="game">
    <div class="gradient"></div>

    <div class="content">
      <h3>{{ props.title }}</h3>
      <h4>{{ props.dev }}</h4>
      <h5>{{ props.release }} - {{ props.version }}</h5>
    </div>
  </div>
</template>
<script setup>
const props = defineProps({
  title: {
    type: String,
    required: true,
  },
  dev: {
    type: String,
    required: true,
  },
  release: {
    type: String,
    required: true,
  },
  version: {
    type: String,
    required: true,
  },
});
</script>
<style scoped lang="scss">
.game {
  position: relative;
  border-radius: 5px;
  margin: 0.5rem;
  width: 200px;
  height: 200px;
  background: url("../assets/images/logo-universal.png") no-repeat center center;
  background-size: cover;
  background-color: #101010;
  cursor: pointer;
  box-shadow: 0 4px 8px 0 rgba(0, 0, 0, 0.2), 0 6px 20px 0 rgba(0, 0, 0, 0.19);
  transition: 150ms;

  &:hover {
    scale: 1.02;
  }

  > .gradient {
    position: absolute;
    width: 200px;
    height: 200px;
    background: linear-gradient(0deg, #171717b8 30%, #ffffff00 100%);
  }

  > .content {
    position: absolute;
    bottom: 1rem;
    left: 1rem;
  }
}
</style>
