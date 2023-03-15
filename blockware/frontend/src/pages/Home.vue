<template>
  <div class="home">
    <div class="title">
      <h1><strong>Block</strong>Ware</h1>
      <h2>by Tom Smith</h2>
    </div>

    <div class="quick-links-wrapper">
      <div class="text">
        <p>{{ text }}</p>
        <p><strong> 1.0.0 </strong></p>
      </div>

      <div class="quick-links" @mouseleave="() => (text = '')">
        <Icon
          v-for="i in icons"
          :colorA="i.colorA"
          :colorB="i.colorB"
          :name="i.name"
          :icon="i.icon"
          :link="i.link"
          class="icon"
          @mouseover="() => (text = i.description)"
        />
      </div>
    </div>
  </div>
</template>
<script setup>
import { ref } from "vue";
import Icon from "../components/home/Icon.vue";

const text = ref("");

const icons = [
  {
    name: "Library",
    icon: "🎮",
    colorA: "#790953",
    colorB: "#ff6464",
    link: "/library",
    description: "View your current library of games",
  },
  {
    name: "Downloads",
    icon: "🎁",
    colorA: "#002cff",
    colorB: "#f800ff",
    link: "/downloads",
    description: "View the progress of your ongoing downloads",
  },
  {
    name: "Store",
    icon: "🏪",
    colorA: "#e200ff",
    colorB: "#ff5100",
    link: "/store",
    description: "Search for new games on BlockWare",
  },
  {
    name: "Upload",
    icon: "🌍",
    colorA: "#ff6a00",
    colorB: "#0012ff",
    link: "/upload",
    description: "Upload a new game to the BlockWare store",
  },
  {
    name: "Peers",
    icon: "🫂",
    colorA: "#00d70a",
    colorB: "#ff0052",
    link: "/peers",
    description: "Manage connections with your peers",
  },
];
</script>
<style scoped lang="scss">
.home {
  height: 100%;
  display: flex;
  flex-direction: column;
  align-items: center;
  padding: 2rem;

  > .quick-links-wrapper {
    margin-top: auto;
    display: flex;
    flex-direction: column;

    > .text {
      display: flex;
      align-items: center;
      padding: 2px 0.5rem;

      > p {
        font-size: 1.2rem;
        font-weight: bold;
        color: darken(white, 20%);
      }

      > *:last-child {
        margin-left: auto;
      }
    }

    > .quick-links {
      align-self: center;
      display: flex;
      position: relative;
      background-color: black;
      padding: 1rem;
      border: solid 5px rgb(75, 75, 75);
      box-shadow: 0 4px 8px 0 rgba(0, 0, 0, 0.2),
        0 6px 20px 0 rgba(0, 0, 0, 0.19);

      &:hover .icon:hover {
        opacity: 1;
      }

      &:hover .icon {
        opacity: 0.7;
      }
    }
  }

  > .title {
    justify-self: flex-start;
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: center;
    margin-top: auto;

    > h1 {
      font-size: 12.5rem;

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
}
</style>
