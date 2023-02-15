<template>
  <nav>
    <div class="title">
      <h4>BlockWare</h4>
    </div>

    <div class="links">
      <router-link to="/library">🕹️ Library</router-link>
      <router-link to="/upload">🆕 Upload</router-link>
      <router-link to="/store">💵 Store</router-link>
    </div>
  </nav>
</template>
<script setup></script>
<style scoped lang="scss">
nav {
  width: 100vw;
  display: flex;
  align-items: center;
  justify-content: center;

  > * {
    padding: 1rem;
  }

  a {
    text-decoration: none;
  }

  > .title {
    font-size: 1.5rem;
  }

  > .links {
    margin-left: auto;
    display: flex;
    gap: 1rem;

    a {
      color: rgb(181, 181, 181);
    }

    a:hover {
      color: white;
    }
  }
}
</style>
