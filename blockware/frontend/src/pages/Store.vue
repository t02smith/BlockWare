<template>
  <div class="store">
    <div class="title">
      <h2>Welcome to the <strong>BlockWare Store</strong></h2>
      <p>
        Here you can view all the games uploaded to BlockWare and purchase them
        for yourself. Enjoy the selection and be sure to send any feedback to
        <a href="mailto:email@notreal.com">this email</a>.
      </p>
    </div>

    <div class="featured">
      <CustomLibrary :games="[]" name="Featured" />
    </div>

    <div class="search">
      <input type="text" name="" id="" placeholder="Search for a game" />
    </div>
  </div>
</template>
<script setup>
import CustomLibrary from "../components/library/CustomLibrary.vue";
</script>
<style scoped lang="scss">
.store {
  strong {
    color: rgb(0, 132, 255);
  }

  > * {
    margin: 2rem;
  }

  > .title {
    display: flex;
    flex-direction: column;

    > * {
      max-width: 700px;
    }

    > h2 {
      font-size: 2rem;
    }

    > p {
      color: darken(white, 15%);
    }
  }

  > .featured {
    > h3 {
      color: orangered;
      font-size: 1.4rem;
    }
  }
}
</style>
