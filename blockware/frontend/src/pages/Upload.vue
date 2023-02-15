<template>
  <form @submit.prevent="submit" class="upload">
    <Error :err="err" v-if="err.length > 0" />

    <div class="success" v-if="success">✅ Game uploaded successfully</div>

    <div class="title">
      <h1>Upload your game here</h1>
      <p>
        Fill in the information below and add your game to a decentralised
        network of games today
      </p>
    </div>

    <div class="form-sections">
      <!-- Game details -->
      <div class="form-section">
        <header>
          <h2>Your Game</h2>
          <p>
            The metadata about your game to help identify it to other users.
          </p>
        </header>

        <div class="form-inputs">
          <!-- title -->
          <div class="form-group">
            <h3>Title</h3>
            <input
              type="text"
              name=""
              id=""
              placeholder="Enter your game's title"
              v-model="title"
            />
          </div>

          <!-- version -->
          <div class="form-group">
            <h3>Version</h3>
            <input
              type="text"
              name=""
              id=""
              placeholder="Enter your game's version"
              v-model="version"
            />
          </div>

          <!-- Developer -->
          <div class="form-group">
            <h3>Developer</h3>
            <input
              type="text"
              name=""
              id=""
              placeholder="Enter your name"
              v-model="dev"
            />
          </div>

          <!-- TODO release => auto generate? -->

          <!-- Price -->
          <div class="form-group">
            <h3>Price (Wei)</h3>
            <input
              type="number"
              name=""
              id=""
              placeholder="in Wei"
              v-model="price"
            />
          </div>
        </div>
      </div>

      <div class="form-section">
        <header>
          <h2>Your Upload</h2>
          <p>Details important to uploading your game to the network.</p>
        </header>

        <div class="form-inputs">
          <!-- Shard size -->
          <div class="form-group">
            <h3>Shard Size (in Bytes)</h3>
            <input
              type="number"
              name=""
              id=""
              placeholder="in bytes"
              v-model="shardSize"
            />
          </div>

          <!-- Root Directory -->
          <div class="form-group">
            <h3>Root Directory of Game</h3>
            <input
              type="text"
              name=""
              id=""
              placeholder="absolute path"
              v-model="file"
            />
          </div>
        </div>
      </div>
    </div>

    <footer>
      <button :disabled="submitted" type="submit">Upload Now!</button>

      <!-- progress -->
      <div class="file-counter">
        <!-- bar -->
        <div class="progress-bar">
          <div class="progress" :style="`width: ${progressWidth}px;`"></div>
        </div>

        <p>{{ fileProgress }}/{{ fileCount }} files</p>
      </div>
    </footer>
  </form>
</template>
<script setup>
import { onMounted, ref, computed } from "vue";
import { UploadGame } from "../../wailsjs/go/main/App.js";
import { EventsOn } from "../../wailsjs/runtime/runtime";
import Error from "../components/Error.vue";

const submitted = ref(false);
const success = ref(false);
const err = ref("");

const title = ref("");
const version = ref("");
const dev = ref("");
const price = ref(0);
const shardSize = ref(16384);
const file = ref("");

// upload progress
const fileCount = ref(0);
const fileProgress = ref(0);

const progressWidth = computed(() =>
  fileProgress.value === 0 ? 0 : (fileProgress.value / fileCount.value) * 300
);

onMounted(async () => {
  EventsOn("file-count", (count) => {
    fileCount.value = count;
    fileProgress.value = 0;
  });
  EventsOn("file-progress", (count) => {
    fileProgress.value = count;
  });
});

async function submit() {
  submitted.value = true;
  err.value = await UploadGame(
    title.value,
    version.value,
    dev.value,
    file.value,
    shardSize.value,
    price.value
  );

  if (err.value.length > 0) {
    title.value = "";
    version.value = "";
    dev.value = "";
    file.value = "";
    shardSize.value = 16384;
    price.value = 0;
    success.value = true;
  }

  submitted.value = false;
}
</script>
<style scoped lang="scss">
.upload {
  width: 100%;
  display: flex;
  flex-direction: column;
  align-items: center;

  > * {
    margin: 1rem;
  }

  > .title {
    display: flex;
    flex-direction: column;
    align-items: center;

    > p {
      font-style: italic;
      color: darken(white, 20%);
    }
  }

  button[type="submit"] {
    background-color: rgb(0, 90, 170);
    padding: 0.8rem 1.75rem;
    border-radius: 10px;
    box-shadow: 0 4px 8px 0 rgba(0, 0, 0, 0.2), 0 6px 20px 0 rgba(0, 0, 0, 0.19);
    font-weight: bold;
    font-size: 1.25rem;
    color: white;
    transition: 150ms;
    cursor: pointer;

    &:hover {
      scale: 1.01;
    }

    &:active {
      scale: 0.99;
    }
  }
}

.form-sections {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 2rem;
  justify-items: center;
  max-width: 1000px;
  margin-bottom: 2rem;

  .form-section {
    background-color: lighten(#131313, 5%);
    border-radius: 10px;
    box-shadow: 0 4px 8px 0 rgba(0, 0, 0, 0.2), 0 6px 20px 0 rgba(0, 0, 0, 0.19);
    padding-bottom: 0.6rem;
    max-width: 600px;
    height: 100%;

    > header {
      background-color: lighten(#131313, 15%);
      padding: 0.5rem 1rem;
      border-radius: 10px 10px 0 0;

      > p {
        font-style: italic;
        color: darken(white, 15%);
      }
    }

    .form-inputs {
      padding: 0.5rem 1.3rem;
      display: flex;
      flex-direction: column;
      gap: 0.6rem;

      > .form-group {
        > input {
          width: 95%;
          padding: 0.6rem 0.4rem;
          border-radius: 5px;
          outline: none;
          border: none;
        }
      }
    }
  }
}

.file-counter {
  display: flex;
  align-items: center;
  gap: 1rem;
  margin-left: auto;

  > .progress-bar {
    position: relative;
    background-color: gray;
    border-radius: 3px;
    width: 300px;
    height: 20px;

    > .progress {
      position: absolute;
      background-color: green;
      height: 20px;
    }
  }

  > p {
    font-weight: bold;
  }
}

footer {
  display: flex;
  align-items: center;
  width: 100%;
  max-width: 950px;
}
</style>
