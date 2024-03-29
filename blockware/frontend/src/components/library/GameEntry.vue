<template>
  <div class="details">
    <div class="header-wrapper">
      <div class="header">
        <img
          :src="`http://localhost:3003/${directory}/assets/${props.game.rootHash}/cover.png`"
          alt=""
        />
        <div class="header-text-wrapper">
          <!-- <img
            :src="`http://localhost:3003/test/data/.toolkit/assets/${props.game.rootHash}/background.png`"
          /> -->

          <div class="header-text">
            <div class="title">
              <h2>{{ props.game.title }}</h2>
              <h3>v{{ props.game.version }}</h3>
            </div>

            <div class="dev">
              <p>From</p>
              <h3>
                {{ props.game.dev }}
              </h3>
              <h4>/</h4>
              <p>{{ props.game.uploader }}</p>
            </div>

            <div class="release-date">
              <p>
                Released on <strong> {{ props.game.release }}</strong>
              </p>
            </div>
          </div>

          <div class="slot">
            <slot />
          </div>
        </div>
      </div>
    </div>

    <div v-html="content" class="markdown"></div>
  </div>
</template>
<script setup>
import { onMounted, ref, watch } from "vue";
import md from "markdown-it";
import { GetDirectory } from "../../../wailsjs/go/controller/Controller";

const props = defineProps({
  game: {
    type: Object,
    required: true,
  },
});

const directory = ref("");
const content = ref(null);

onMounted(() => updateContent());
watch(props.game, () => updateContent());

async function updateContent() {
  directory.value = await GetDirectory();
  const res = await fetch(
    `http://localhost:3003/${directory.value}/assets/${props.game.rootHash}/description.md`
  );

  content.value = md().render(await res.text());
}
</script>
<style scoped lang="scss">
.details {
  display: flex;
  flex-direction: column;
  gap: 1rem;
  padding: 1rem;
  overflow-y: auto;
  background-color: #131313;

  > .header-wrapper {
    position: relative;
    padding-bottom: 1rem;
    background-size: cover;

    > .header {
      display: flex;
      align-items: flex-end;
      gap: 1rem;
      position: sticky;

      > img {
        width: 200px;
        height: 200px;
        background-color: lighten(#131313, 5%);
        border-radius: 10px;
        padding: 2px;
        box-shadow: 0 4px 8px 0 rgba(0, 0, 0, 0.2),
          0 6px 20px 0 rgba(0, 0, 0, 0.19);
      }

      > .header-text-wrapper {
        display: grid;
        grid-template-columns: 3fr 1fr;
        width: 100%;

        > * {
          margin: 0.5rem 1rem;
        }

        > img {
          grid-column: span 2;
        }

        > .header-text {
          display: flex;
          flex-direction: column;

          > .title {
            display: flex;
            align-items: flex-end;
            gap: 0.25rem;

            > h2 {
              font-size: 4.25rem;
              margin-bottom: -0.9rem;
            }

            > h3 {
              color: rgb(0, 183, 255);
            }
          }

          > .release-date {
            margin-left: 5px;
            color: darken(white, 25%);
          }

          > .dev {
            display: flex;
            align-items: center;
            gap: 0.5rem;
            overflow-x: hidden;
            margin-top: 15px;
            margin-left: 5px;

            > h3 {
              color: rgb(0, 183, 255);
            }

            > p {
              color: darken(white, 25%);

              &:last-child {
                font-weight: bold;
              }
            }
          }
        }

        > .slot {
          place-self: flex-end;
          margin-left: auto;
          display: flex;
          flex-direction: column;
          justify-content: center;
          align-items: flex-end;
          gap: 0.5rem;

          > h3 {
            font-size: 1.6rem;
            color: darken(white, 15%);
          }

          > button {
            cursor: pointer;
            padding: 5px 20px;
            font-size: 1rem;
            border-radius: 4px;
            color: white;
            border: none;
            background-color: rgb(0, 116, 48);
            font-weight: bold;
            box-shadow: 0 4px 8px 0 rgba(0, 0, 0, 0.2),
              0 6px 20px 0 rgba(0, 0, 0, 0.19);
          }
        }
      }
    }
  }

  > p {
    color: darken(white, 20%);
  }
}
</style>
