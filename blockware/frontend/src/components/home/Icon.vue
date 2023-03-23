<template>
  <router-link :to="props.link" class="wrapper" :class="`${props.class}`">
    <div class="icon">
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
  class: {
    type: String,
    default: "",
  },
});

const background = `linear-gradient(144deg, ${props.colorA} 0%, ${props.colorB} 150%)`;
const backgroundHover = `linear-gradient(144deg, ${props.colorB} 0%, ${props.colorA} 150%)`;
</script>
<style>
.wrapper > .icon {
  background: v-bind("background");
}

.wrapper:hover > .icon {
  background: v-bind("backgroundHover");
  transition: background 1000ms linear;
}
</style>
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
  transition: 150ms;

  &:hover,
  &:focus {
    scale: 1.05;
    z-index: 100;
  }

  > .icon {
    width: 200px;
    height: 200px;
    display: grid;
    place-items: center;
    font-size: 8rem;
    border: 2px solid gray;

    &:hover {
      box-shadow: 0 4px 8px 0 rgba(0, 0, 0, 0.2),
        0 6px 20px 0 rgba(0, 0, 0, 0.19);
      opacity: 1;
    }
  }

  > h4 {
    font-size: 1.75rem;
    font-weight: bold;
  }
}
</style>
