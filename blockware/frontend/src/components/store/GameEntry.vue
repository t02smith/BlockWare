<template>
  <div class="details">
    <div class="header-wrapper">
      <div class="gradient"></div>
      <div class="header">
        <img src="../../assets/images/logo-universal.png" alt="" />

        <div class="header-text-wrapper">
          <div class="header-text">
            <h2>{{ props.game.title }}</h2>
            <h3>{{ props.game.dev }}</h3>
          </div>

          <div class="buy" v-if="props.store">
            <h3>
              Buy today for <strong>{{ props.game.price }} ETH</strong>
            </h3>
            <button>🔒 Checkout now!</button>
          </div>
        </div>
      </div>
    </div>

    <p>
      Lorem ipsum dolor, sit amet consectetur adipisicing elit. Deleniti
      dignissimos aliquid quod, quae perspiciatis incidunt, aperiam a illum
      facilis voluptate provident consequuntur iste fugit voluptatem? Ducimus et
      similique eius, excepturi quo maiores! Incidunt saepe magnam laudantium
      earum sed dolores dolor natus beatae, tempore labore sunt exercitationem
      quidem nam quis non.
    </p>

    <p>
      Lorem ipsum dolor sit amet consectetur, adipisicing elit. Vitae nemo,
      dolorem quas excepturi totam itaque ducimus tempore fugit quo sed, amet
      adipisci consequuntur eos accusantium, iste impedit omnis cumque delectus!
      Iusto modi, ipsam omnis error soluta vel distinctio consequuntur? Eaque
      nihil nobis itaque libero? Praesentium, rem unde. At beatae rem laborum
      iste enim incidunt natus amet fugit id quos, eligendi adipisci illo ipsum
      nostrum fugiat necessitatibus architecto provident recusandae optio nemo
      molestiae ex veniam aperiam.
    </p>
  </div>
</template>
<script setup>
const props = defineProps({
  game: {
    type: Object,
    required: true,
  },
  store: {
    type: Boolean,
    default: false,
  },
});
</script>
<style scoped lang="scss">
.details {
  margin: 1rem;
  display: flex;
  flex-direction: column;
  gap: 1rem;
  width: 95%;

  > .header-wrapper {
    width: 100%;
    position: relative;
    background: url("../../assets/images/logo-universal.png") no-repeat;
    padding-bottom: 1rem;

    > .gradient {
      position: absolute;
      width: 150%;
      height: 100%;
      background: linear-gradient(0deg, #171717b8 30%, #ffffff00 100%);
    }

    > .header {
      display: flex;
      align-items: flex-end;
      gap: 1rem;
      width: 100%;
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
        display: flex;
        align-items: flex-end;
        width: 100%;

        > .header-text {
          display: flex;
          flex-direction: column;

          > h2 {
            font-size: 4rem;
            margin-bottom: -0.9rem;
          }

          > h3 {
            color: rgb(0, 183, 255);
          }
        }

        > .buy {
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
            outline: none;
            background-color: rgb(0, 116, 48);
            font-weight: bold;
            box-shadow: 0 4px 8px 0 rgba(0, 0, 0, 0.2),
              0 6px 20px 0 rgba(0, 0, 0, 0.19);
          }
        }
      }
    }
  }

  > .buy {
  }

  > p {
    color: darken(white, 20%);
  }
}
</style>
