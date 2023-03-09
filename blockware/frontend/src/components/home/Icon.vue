<template>
  <router-link :to="props.link" class="wrapper">
    <div
      class="icon"
      :style="`background: linear-gradient(144deg, ${props.colorA} 0%, ${props.colorB} 150%);`"
    >
      {{ props.icon }}
    </div>
    <h4>{{ props.name }}</h4>
  </router-link>
</template>
<script setup>
const props = defineProps({
  colorA: {
    type: String,
    required: true,
  },
  colorB: {
    type: String,
    default: "#fff",
  },
  name: {
    type: String,
    required: true,
  },
  icon: {
    type: String,
    required: true,
  },
  link: {
    type: String,
    default: "/home",
  },
});
</script>
<style scoped lang="scss">
.wrapper {
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  transition: 150ms;
  cursor: pointer;
  text-decoration: none;
  color: white;

  &:hover {
    scale: 1.02;
  }

  > .icon {
    width: 100px;
    height: 100px;
    display: grid;
    place-items: center;
    font-size: 4rem;
    border-radius: 10px;
    box-shadow: 0 4px 8px 0 rgba(0, 0, 0, 0.2), 0 6px 20px 0 rgba(0, 0, 0, 0.19);
  }

  > h4 {
    font-size: 2rem;
    font-weight: bold;
  }
}
</style>
