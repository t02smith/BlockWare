<template>
  <div class="home">
    <div class="title">
      <h1><strong>Block</strong>Ware</h1>
      <h2>by Tom Smith</h2>
    </div>

    <div class="quick-links">
      <Icon
        v-for="i in icons"
        :colorA="i.colorA"
        :colorB="i.colorB"
        :name="i.name"
        :icon="i.icon"
        :link="i.link"
      />
    </div>
  </div>
</template>
<script setup>
import Icon from "../components/home/Icon.vue";

const icons = [
  {
    name: "Library",
    icon: "🎮",
    colorA: "#790953",
    colorB: "#ff6464",
    link: "/library",
  },
  {
    name: "Downloads",
    icon: "🎁",
    colorA: "#8810d2",
    colorB: "#00ff2e",
    link: "/downloads",
  },
  {
    name: "Store",
    icon: "🏪",
    colorA: "#ffbf64",
    colorB: "#ff6474",
    link: "/store",
  },
  {
    name: "Upload",
    icon: "🌍",
    colorA: "#6472ff",
    colorB: "#ffa264",
    link: "/upload",
  },
];
</script>
<style scoped lang="scss">
.home {
  height: 100%;
  display: grid;
  place-items: center;

  > .title {
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: center;

    > h1 {
      font-size: 10rem;

      > strong {
        background: url("../assets/images/icon.png");
        background-attachment: fixed;
        -webkit-background-clip: text;
        background-clip: text;
        -webkit-text-fill-color: transparent;
      }
    }

    > h2 {
      margin-top: -1.75rem;
      color: darken(white, 20%);
    }
  }

  > .quick-links {
    margin: 1rem;
    display: flex;
    align-items: center;
    justify-content: center;
    gap: 2.5rem;
  }
}
</style>
