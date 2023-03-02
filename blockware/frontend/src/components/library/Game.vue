<template>
  <router-link
    :to="`/${props.linkTo}?game=${props.hash}`"
    class="game"
    :style="`background: url('http://localhost:3003/test/data/.toolkit/assets/${props.hash}/cover.png') no-repeat center;`"
  >
    <div class="gradient"></div>
    <h5 class="version">v{{ props.version }}</h5>

    <div class="content">
      <h3>{{ props.title }}</h3>
      <h4>{{ props.dev }}</h4>
    </div>
  </router-link>
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
  version: {
    type: String,
    required: true,
  },
  hash: {
    type: String,
    required: true,
  },
  linkTo: {
    type: String,
    default: "library",
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
  background-size: cover;
  background-color: #101010;
  cursor: pointer;
  box-shadow: 0 4px 8px 0 rgba(0, 0, 0, 0.2), 0 6px 20px 0 rgba(0, 0, 0, 0.19);
  transition: 150ms;

  &:hover {
    scale: 1.01;
  }

  &:active {
    scale: 0.99;
  }

  > .version {
    position: absolute;
    top: 1rem;
    right: 1rem;
    font-size: 1rem;
    color: white;
    background-color: rgb(0, 88, 219);
    border-radius: 7px;
    padding: 1px 7px;
    opacity: 0.8;
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
    line-height: 1.2rem;
    color: white;

    h3 {
      font-size: 1.5rem;
    }

    > h4 {
      font-style: italic;
      color: darken(white, 30%);
    }
  }
}
</style>
