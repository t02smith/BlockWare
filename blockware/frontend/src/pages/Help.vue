<template>
  <div class="wrapper">
    <div class="help">
      <div class="title">
        <h1>Frequently Asked Questions</h1>
        <p>
          Want to learn more about <strong>BlockWare</strong> or have some
          questions about how to use the app? Consult some of the questions
          below to try to find an answer.
        </p>
      </div>

      <div class="questions">
        <hr />
        <div class="question">
          <h4>What is <strong>BlockWare</strong>?</h4>
          <div class="answer">
            <p>
              <strong>BlockWare</strong> is a decentralised video game
              marketplace that allows developers and gamers to interact
              directly. It is built on top of technologies like Ethereum and
              IPFS and includes a custom P2P file sharing network for
              distributing games. Here you can buy or sell games to other people
              in the network without fear of censorship and without large,
              recurring fees from a middle-man to facilitate the selling and
              distribution. Using Ethereum means your ownership of everything
              you buy here is much more secure and not reliant on any single
              platform staying active.
            </p>
            <p>
              <strong>BlockWare</strong> was created as part of my (Tom Smith
              tcsg20) final year project for a BSc on Computer Science at the
              University of Southampton.
            </p>
          </div>
        </div>
        <hr />
        <div class="question">
          <h4>What does each page do?</h4>
          <div class="answer">
            <p>There are 5 main pages available for you to use:</p>
            <ul>
              <li>
                <strong>Library</strong> view a list of your owned games. Here
                you can also choose to install/uninstall any game. Also if you
                own a game but it hasn't appeared in your library you can import
                it using its root hash.
              </li>
              <li>
                <strong>Downloads</strong> Here you can view the progress of
                each of your downloads from the setup till its complete.
              </li>
              <li>
                <strong>Upload</strong> Here you can enter the details about
                your new game (or an update to an existing one) and upload it to
                the BlockWare network to be seen and purchased by anyone on the
                network.
              </li>
              <li>
                <strong>Store</strong> Here you can view and search for games in
                the BlockWare store. If you see a game you like you can purchase
                it and add it to your own library.
              </li>
              <li>
                <strong>Peers</strong> Here you manage your connections with
                your peers whom you will download games off of. Simply enter in
                the addresses of all the peers you want to connect to and
                BlockWare will do the rest. When you're downloading a game these
                peers will be sent requests for data from that game.
              </li>
            </ul>
          </div>
        </div>
        <hr />
        <div class="question">
          <h4>What do I need to include in my assets folder?</h4>
          <div class="answer">
            <p>
              You need to include at least the following files in the root
              directory:
            </p>
            <ul>
              <li>
                <strong>cover.png</strong> The main piece of cover art for your
                game.
              </li>
              <li>
                <strong>description.md</strong> A description file that gives an
                overview of your game.
              </li>
            </ul>
            <p>
              You can include any number of other files inside the assets folder
              and even reference them in your description file to make your
              library entry stand out.
            </p>
          </div>
        </div>
        <hr />
        <div class="question">
          <h4>What does Shard Size mean from the game upload page?</h4>
          <div class="answer">
            <p>
              When you upload a game, we break each file in every directory and
              subdirectory into fixed size shards. You download games by
              requesting shards of data off of your peers and adding them to the
              correct location. Each shard of data is represented by a SHA256
              hash, which acts as a unique fingerprint of the data so that when
              you download it you can easily verify that what you've been sent
              is what it was supposed to be.
            </p>
            <p>
              A small shard size will mean you will create many blocks so we
              advise that you use a minimum shard size of
              <strong>4MiB</strong> (or <strong>4,194,304 bytes</strong>) to
              improve download speeds. Don't worry about data smaller than the
              shard size as it is compressed before its sent.
            </p>
          </div>
        </div>
        <hr />
        <div class="question">
          <h4>How can I create an Ethereum private key?</h4>
          <div class="answer">
            <p>
              You'll need an Ethereum private key to access the BlockWare
              network and to purchase and upload games. We recommend you use the
              browser extension MetaMask by following this guide
              <a
                href="https://support.metamask.io/hc/en-us/articles/360015489531"
                >https://support.metamask.io/hc/en-us/articles/360015489531</a
              >
              to download it. This should set you up with a default account,
              which you can use. Then follow this tutorial
              <a
                href="https://support.metamask.io/hc/en-us/articles/360015289632-How-to-export-an-account-s-private-key"
                >https://support.metamask.io/hc/en-us/articles/360015289632-How-to-export-an-account-s-private-key</a
              >
              to export your private key.
            </p>
            <p class="warning">
              ⚠️ Remember not to share you private key with anybody! This
              application will only use it for signing transactions and
              determining your address and it will never be transmitted
              elsewhere or saved between instances of the application.
            </p>
          </div>
        </div>
        <hr />
        <div class="question">
          <h4>My download is stuck even though I'm connected to peers?</h4>
          <div class="answer">
            <p>
              Simply click the unstick button on the downloads page. This will
              load any requests that have been postponed to be re-requested from
              your peers. However, if none of your peers have the block you
              need, you will not be able to download it until they do or you
              connect to a peer who does.
            </p>
            <p>
              If you are certain that a peer has the data you need but is not
              sending it, try raising an issue by contacting
              <strong>tcs1g20@soton.ac.uk</strong>.
            </p>
          </div>
        </div>
        <hr />
        <div class="question">
          <h4>I own a game but it is not appearing in my library?</h4>
          <div class="answer">
            <p>
              The only games that will appear in your library are the ones you
              have cached locally, which may not be your full library. If you
              hit the "🌍 import" button and enter the root hash of any missing
              game it will be fetched from the BlockWare network and stored
              locally.
            </p>
          </div>
        </div>
        <hr />
        <div class="question">
          <h4>Can I trust any peer I connect to?</h4>
          <div class="answer">
            <p>
              We have several mechanisms in place to boost the trust you can
              have with your peers:
            </p>
            <ul>
              <li>
                When connecting to a peer you will exchange signed messages that
                prove your Ethereum identity. The message is generated
                differently at each connection so a valid peer is way more
                likely to be who they say they are.
              </li>
              <li>
                Each authorised command will include the data required to check
                the integrity of the data provided. This data will correspond to
                an element of a game's hash tree and thanks to Ethereum's and
                IPFS's inherent immutability properties we know what data we
                should receive and what data we don't need.
              </li>
              <li>
                If you think a peer may be attempting to act maliciously you can
                remove them from the Peers page.
              </li>
            </ul>
          </div>
        </div>
        <hr />
        <div class="question">
          <h4>Why are some games smaller in my library?</h4>
          <div class="answer">
            <p>
              The smaller games represent old versions of games and the larger
              version represents the most recent versions. This is to help you
              see the latest changes without losing access to old updates.
            </p>
            <p>
              You can toggle this on and off by ticking the
              <strong>show old versions</strong> box in the top right hand
              corner of your library page.
            </p>
          </div>
        </div>
        <hr />
        <div class="question">
          <h4>What is peer discovery?</h4>
          <div class="answer">
            <p>
              When you connect to a new peer you will exchange your list of
              connected peers to allow each other to discover new peers. Peers
              will connect to you using your public hostname (or IP address) but
              the peers that you connect to won't know what this is unless we
              tell them.
            </p>
            <p>
              By entering your public hostname in the peers menu, we allow other
              users to share our server information to their peers.
            </p>
            <p>
              If you don't know what your public IP address is you can use a
              website like
              <a href="https://www.whatismyip.com/"
                >https://www.whatismyip.com/</a
              >. It's likely that you may need to enable port forwarding for
              your device at your home router.
            </p>
          </div>
        </div>
        <hr />
      </div>
    </div>
  </div>
</template>
<script setup></script>
<style scoped lang="scss">
hr {
  opacity: 0.6;
}

.wrapper {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  height: 100%;
  overflow: hidden;
  background-color: lighten(#131313, 2%);
}

.help {
  background-color: #131313;
  display: flex;
  flex-direction: column;
  padding: 1rem;
  padding-bottom: 3rem;
  width: 75%;
  max-width: 1200px;
  height: 100%;
  overflow: auto;
  gap: 1rem;

  > .title {
    h1 {
      font-size: 3rem;
    }
  }

  > .questions {
    > .question {
      margin: 0.5rem 0rem;

      > h4 {
        font-size: 1.5rem;
      }

      > .answer {
        display: flex;
        flex-direction: column;
        gap: 0.5rem;

        p,
        li {
          color: rgb(202, 202, 202);

          > a {
            text-decoration: none;
            color: rgb(0, 131, 253);

            &:hover {
              text-decoration: underline;
            }
          }

          &.warning {
            font-weight: bold;
            color: orange;
          }
        }

        > ul {
          margin-left: 2rem;
        }
      }
    }
  }
}
</style>
