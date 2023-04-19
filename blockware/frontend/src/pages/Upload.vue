<template>
  <div class="wrapper">
    <form @submit.prevent="submit" class="upload">
      <h2>Upload your content to <strong>BlockWare</strong></h2>
      <Error />
      <div class="sections">
        <div class="section initial">
          <div class="title">
            <h3>1. What are you uploading?</h3>
            <div class="line"></div>
          </div>

          <div class="radio">
            <div class="in">
              <input
                type="radio"
                name=""
                id="own-game"
                value="own-game"
                v-model="type"
              />
              <label for="own-game"><strong>A brand new game</strong></label>
            </div>

            <p>
              Upload the first ever version of your game to the BlockWare
              network. Users will be able to find, buy and install your game
              just make sure you seed it! Once uploaded, you will be able to
              release as many updates as you desire and your users will help to
              distribute your game for you. To incentivise this, make sure you
              reward those users who contribute.
            </p>
          </div>

          <div class="radio">
            <div class="in">
              <input
                type="radio"
                name=""
                id="existing-game"
                value="existing-game"
                v-model="type"
              />
              <label for="existing-game"
                ><strong>An update to an exiting game</strong></label
              >
            </div>

            <p>
              Already have a game? Have a life or death patch you need to
              release? This is the option for you. Select which game you want to
              update, and most of the info will already be filled out for you
              just make sure to point to the right directory.
            </p>

            <div class="to-update">
              <select
                name=""
                id=""
                :disabled="ownedGames.length === 0"
                v-model="selectedOwnGame"
              >
                <option value="" disabled>Choose game:</option>

                <option :value="g" v-for="g in ownedGames">
                  <div>{{ g.title }} v{{ g.version }}</div>
                </option>
              </select>

              <p v-if="ownedGames && ownedGames.length === 0">
                You own no games
              </p>
            </div>
          </div>
        </div>

        <div class="section details">
          <div class="title">
            <h3>2. Your game info:</h3>
            <div class="line"></div>
          </div>

          <div class="fields">
            <div class="field">
              <div class="info">
                <h6>Title</h6>
                <p>what is your game called?</p>
              </div>
              <input
                type="text"
                name=""
                id=""
                :disabled="selectedOwnGame !== ''"
                placeholder="title"
                v-model="title"
              />
            </div>

            <div class="field">
              <div class="info">
                <h6>Developer</h6>
                <p>what is your/your companies name?</p>
              </div>
              <input
                type="text"
                name=""
                id=""
                :disabled="selectedOwnGame !== ''"
                placeholder="developer"
                v-model="dev"
              />
            </div>

            <div class="field">
              <div class="info">
                <h6>Version</h6>
                <p>what version are you releasing?</p>
              </div>
              <input
                v-model="version"
                type="text"
                name=""
                id=""
                placeholder="version number"
              />
            </div>

            <div class="field">
              <div class="info">
                <h6>Price</h6>
                <p>How expensive is your game (in Wei)?</p>
              </div>
              <input
                v-model="price"
                type="number"
                name=""
                id=""
                placeholder="price"
              />
            </div>
          </div>
        </div>

        <div class="section your-upload">
          <div class="title">
            <h3>3. Your upload:</h3>
            <div class="line"></div>
          </div>

          <div class="fields">
            <div class="field">
              <div class="info">
                <h6>Root Directory</h6>
                <p>select the root directory of your game</p>
              </div>
              <button
                value=""
                @click="async () => (file = await SelectFolder())"
              >
                {{ file === "" ? "Upload Folder" : file }}
              </button>
            </div>

            <div class="field">
              <div class="info">
                <h6>Assets Directory</h6>
                <p>required files: cover.png, description.md</p>
              </div>
              <button
                value=""
                @click="async () => (assets = await SelectFolder())"
              >
                {{ assets === "" ? "Upload Folder" : assets }}
              </button>
            </div>

            <div class="field">
              <div class="info">
                <h6>Shard Size</h6>
                <p>
                  what size shards (in bytes) should each file be broken into?
                </p>
              </div>
              <input
                v-model="shardSize"
                type="number"
                name=""
                id=""
                placeholder="shard size"
              />
            </div>
          </div>
        </div>

        <div class="section summary">
          <div class="title">
            <h3>4. Summary</h3>
            <div class="line"></div>
          </div>

          <div class="info">
            <p>
              Double check all the fields above and hit submit! After hitting
              submit this application will:
            </p>
            <ol>
              <li>Create a hash tree of your application</li>
              <li>Upload your hash tree and assets to IPFS</li>
              <li>Upload your game metadata to Ethereum</li>
              <li>Begin seeding your new game in the background</li>
            </ol>
          </div>

          <div class="license">
            <input type="checkbox" name="" id="license" v-model="license" />
            <label for="license">
              Click here to agree to BlockWare's game licensing policy
            </label>
          </div>

          <div class="submit">
            <button type="submit" :disabled="submitted">
              Upload your game!
            </button>

            <div class="file-counter">
              <div class="progress-bar">
                <div
                  class="progress"
                  :style="`width: ${progressWidth}px;`"
                ></div>
              </div>

              <p>{{ fileProgress }}/{{ fileCount }} files</p>
            </div>
          </div>
        </div>
      </div>
    </form>
  </div>
</template>
<script setup>
import { onMounted, ref, computed, watch } from "vue";
import {
  UploadGame,
  SelectFolder,
} from "../../wailsjs/go/controller/Controller";
import { EventsOn } from "../../wailsjs/runtime/runtime";
import Error from "../components/Error.vue";
import { useGamesStore } from "../stores/games";
import { useErrStore } from "../stores/err";

/*

Upload a new game to the network by filling in its
information. Users may release a new game or an update
to an existing game

*/

// Pinia stores
const games = useGamesStore();

// form managers

// true whilst the form is being submitted but not complete
const submitted = ref(false);

// whether the upload was a success
const success = ref(false);

// input fields

// 1 => new game OR update
const type = ref(null);

// for tracking the user's owned games that are updateable
const ownedGames = ref([]);

const selectedOwnGame = ref("");

// 2 => Game details
const title = ref("");
const version = ref("");
const dev = ref("");
const price = ref(0);
const assets = ref("");
const file = ref("");

// 3 => upload details
const shardSize = ref(16384);
const workers = ref(5);

// 4 => accept license
const license = ref(false);

// upload progress for progress bar
const fileCount = ref(0);
const fileProgress = ref(0);
const progressWidth = computed(() =>
  fileProgress.value === 0 ? 0 : (fileProgress.value / fileCount.value) * 300
);

const err = useErrStore();

onMounted(async () => {
  // listen to progress bar events
  EventsOn("file-count", (count) => {
    fileCount.value = count;
    fileProgress.value = 0;
  });
  EventsOn("file-progress", (count) => {
    fileProgress.value = count;
  });

  await games.refreshOwnedGames();

  let gameLS = [];
  games.ownedGames
    .filter((g) => g.IsOwner)
    .sort((a, b) => a.release > b.release)
    .forEach(
      (g) => !gameLS.find((_g) => _g.title === g.title) && gameLS.push(g)
    );
  ownedGames.value = gameLS;
});

//
watch(selectedOwnGame, () => {
  if (!selectedOwnGame.value) return;

  selectOwnedGame(selectedOwnGame.value);
});

/*
Select an owned game to update
Will prevent certain fields from being changed
*/
function selectOwnedGame(g) {
  title.value = g.title;
  version.value = g.version;
  dev.value = g.dev;
  price.value = g.price;
}

/*
Upload the user's new game
*/
async function submit() {
  // validate input
  if (workers.value <= 0) workers.value = 1;
  if (shardSize.value <= 0) shardSize.value = 16384;
  if (!license.value) return;

  // submit
  submitted.value = true;
  await UploadGame(
    title.value,
    version.value,
    dev.value,
    file.value,
    shardSize.value,
    price.value,
    workers.value,
    assets.value
  );

  // reset fields
  if (err.err.length === 0) {
    title.value = "";
    version.value = "";
    dev.value = "";
    file.value = "";
    shardSize.value = 16384;
    price.value = 0;
    workers.value = 5;
    assets.value = "";
    success.value = true;
    games.refreshOwnedGames();
  } else {
    success.value = false;
  }

  submitted.value = false;
}
</script>
<style scoped lang="scss">
.wrapper {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  height: 100%;
  overflow: hidden;
  background-color: lighten(#131313, 2%);
}

.upload {
  background-color: #131313;
  display: flex;
  flex-direction: column;
  padding: 1rem;
  padding-bottom: 3rem;
  width: 75%;
  max-width: 1200px;
  height: 100%;
  overflow: auto;

  > h2 {
    margin: 2rem;
    font-size: 4rem;
    background-color: lighten(#131313, 5%);
    border-radius: 10px;
    box-shadow: 0 4px 8px 0 rgba(0, 0, 0, 0.2), 0 6px 20px 0 rgba(0, 0, 0, 0.19);
    padding: 0.5rem 1.25rem;
  }

  > .sections {
    > .section {
      align-self: flex-start;
      width: 100%;

      > .title {
        display: flex;
        align-items: center;
        gap: 1rem;
        font-size: 2rem;
        width: 100%;

        > h3 {
          width: fit-content;
        }

        > .line {
          background-color: white;
          height: 3px;
          margin-top: 10px;
          flex: 1;
        }
      }

      &.initial {
        display: flex;
        flex-direction: column;
        gap: 0.75rem;
        > .radio {
          padding: 0 1rem;
          > .in {
            display: flex;
            align-items: center;
            gap: 4px;

            > label {
              font-size: 1.1rem;
              font-weight: bold;
            }
          }

          > p {
            color: darken(white, 15%);
          }
        }

        .to-update {
          margin-top: 10px;
          display: flex;
          gap: 1rem;

          > p {
            font-style: italic;
            color: red;
          }

          > select {
            padding: 4px 10px;
            border-radius: 2px;
            font-size: 1.1rem;
            font-weight: bold;
            background-color: rgb(0, 131, 253);
            border: none;
            color: white;
            cursor: pointer;

            > * {
              font-weight: bold;
              cursor: pointer;
              padding: 8px 10px;
            }
          }
        }
      }

      &.details,
      &.your-upload {
        > .fields {
          display: flex;
          flex-direction: column;
          gap: 10px;

          > .field {
            display: flex;
            width: 100%;

            > .info {
              display: flex;
              align-items: center;
              gap: 5px;

              > p {
                color: darken(white, 25%);
                font-style: italic;
              }

              > h6 {
                font-size: 1.2rem;
                color: rgb(0, 131, 253);
                font-weight: bold;
              }
            }

            > input,
            button {
              margin-left: auto;
              width: 250px;
              padding: 3px 5px;
              background-color: lighten(#131313, 10%);
              outline: none;
              border: none;
              color: white;
            }

            > button {
              width: 260px;

              &:hover {
                color: rgb(0, 131, 253);
              }
            }
          }
        }
      }

      &.summary {
        display: flex;
        flex-direction: column;
        gap: 1rem;

        > *:first-child {
          margin-bottom: -1rem;
        }

        > .info {
          > ol {
            margin-left: 1rem;
          }
        }

        > .license {
          color: darken(white, 15%);
          font-style: italic;
        }

        > .submit {
          display: flex;
          align-items: center;

          > button {
            cursor: pointer;
            background-color: rgb(0, 104, 0);
            color: white;
            font-weight: bold;
            padding: 5px 10px;
            border-radius: 5px;
            font-size: 1.15rem;
            transition: 150ms;

            &:hover {
              background-color: lighten(rgb(0, 104, 0), 6%);
            }
          }

          .file-counter {
            display: flex;
            align-items: center;
            gap: 1rem;
            margin-left: auto;
            margin-right: 1rem;

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
        }
      }
    }
  }
}
</style>
