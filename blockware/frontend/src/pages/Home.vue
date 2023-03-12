<template>
  <div class="home">
    <div class="title">
      <h1><strong>Block</strong>Ware</h1>
      <h2>by Tom Smith</h2>
    </div>

    <div class="quick-links-wrapper">
      <div class="left-joycon joycon">
        <div class="bumper">
          <div class="cover"></div>
        </div>
        <div class="minus"></div>

        <div class="joystick"></div>

        <div class="buttons">
          <div></div>
          <div></div>
          <div></div>
          <div></div>
        </div>

        <div class="screenshot"></div>
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
      <div class="right-joycon joycon">
        <div class="bumper">
          <div class="cover"></div>
        </div>
        <div class="plus">
          <div></div>
          <div></div>
        </div>

        <div class="buttons">
          <div></div>
          <div></div>
          <div></div>
          <div></div>
        </div>

        <div class="joystick"></div>

        <div class="home"></div>
      </div>
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
    colorA: "#002cff",
    colorB: "#f800ff",
    link: "/downloads",
  },
  {
    name: "Upload",
    icon: "☁️",
    colorA: "#00c2ff",
    colorB: "#ff0081",
    link: "/upload",
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
  {
    name: "Peers",
    icon: "🫂",
    colorA: "#f7fb3f",
    colorB: "#35de00",
    link: "/peers",
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

  > .quick-links-wrapper {
    display: flex;
    align-items: center;
    justify-content: center;
    background-color: lighten(#131313, 10%);
    border-radius: 40px;
    box-shadow: 0 4px 8px 0 rgba(0, 0, 0, 0.2), 0 6px 20px 0 rgba(0, 0, 0, 0.19);

    > .joycon {
      width: 80px;
      height: 300px;
      padding: 1rem;
      position: relative;

      > * {
        position: absolute;
      }

      > .joystick {
        border-radius: 50%;
        background-color: black;
        height: 70px;
        width: 70px;
      }

      > .buttons {
        position: relative;

        > * {
          position: absolute;
          border-radius: 50%;
          background-color: black;
          height: 28px;
          width: 28px;
        }
      }

      > .bumper {
        top: -5px;
        width: 100px;
        height: 100px;
        z-index: -1;
        background-color: black;

        > .cover {
          position: absolute;
          background-color: #131313;
          top: 40px;
          width: 100px;
          height: 100px;
        }
      }
    }

    > .left-joycon {
      background-color: green;
      border-radius: 40px 0 0 40px;

      > .bumper {
        border-top-left-radius: 40px;
        left: -5px;
      }

      > .minus {
        background-color: black;
        height: 7px;
        width: 20px;
        right: 10px;
      }

      > .joystick {
        left: 23px;
        top: 60px;
      }

      > .screenshot {
        background-color: black;
        width: 20px;
        height: 20px;
        right: 15px;
        top: 240px;
      }

      > .buttons {
        top: 140px;

        > div:first-child {
          left: 30px;
        }

        > div:nth-child(2) {
          top: 24px;
          left: 5px;
        }

        > div:nth-child(3) {
          top: 24px;
          right: 0px;
        }

        > div:nth-child(4) {
          top: 48px;
          left: 30px;
        }
      }
    }

    > .right-joycon {
      background-color: rgb(255, 22, 80);
      border-radius: 0 40px 40px 0;

      > .bumper {
        border-top-right-radius: 40px;
        right: -5px;
      }

      > .plus {
        position: relative;
        right: 8px;
        top: -4px;

        > * {
          background-color: black;
          position: absolute;
        }

        > div:first-child {
          left: 6px;
          height: 20px;
          width: 7px;
        }

        > div:nth-child(2) {
          top: 7px;
          height: 7px;
          width: 20px;
        }
      }

      > .buttons {
        top: 40px;
        right: 7px;

        > div:first-child {
          left: 30px;
        }

        > div:nth-child(2) {
          top: 24px;
          left: 5px;
        }

        > div:nth-child(3) {
          top: 24px;
          right: 0px;
        }

        > div:nth-child(4) {
          top: 48px;
          left: 30px;
        }
      }

      > .joystick {
        right: 23px;
        top: 160px;
      }

      > .home {
        border-radius: 50%;
        background-color: black;
        width: 28px;
        height: 28px;
        top: 240px;
      }
    }

    > .quick-links {
      background-color: black;
      margin: 0.5rem;
      padding: 0.5rem 3rem;
      display: grid;
      grid-template-columns: 1fr 1fr 1fr;
      place-items: center;
      gap: 1rem;
      height: 300px;
    }
  }
}
</style>
