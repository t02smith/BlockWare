<template>
  <router-link to="/store" class="game">
    <div class="gradient"></div>

    <div class="content">
      <h3>+</h3>
      <h4>Add a new game</h4>
    </div>
  </router-link>
</template>
<script setup></script>
<style scoped lang="scss">
.game {
  position: relative;
  border-radius: 5px;
  margin: 0.5rem;
  width: 200px;
  height: 200px;
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

  > * {
    position: absolute;
    width: 200px;
    height: 200px;
  }

  > .gradient {
    background: linear-gradient(0deg, #171717b8 30%, #ffffff00 100%);
  }

  > .content {
    display: grid;
    grid-template-rows: 4fr 1fr;
    place-items: center;

    > h3 {
      font-size: 6rem;
      color: white;
    }

    > h4 {
      font-size: 1.2rem;
      color: darken(white, 30%);
    }
  }
}
</style>
