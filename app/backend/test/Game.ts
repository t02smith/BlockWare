import { ethers } from "hardhat";
import { expect } from "chai";
import { loadFixture } from "@nomicfoundation/hardhat-network-helpers";

describe("Game", () => {
  async function deploySampleGame() {
    const [owner, otherAccount] = await ethers.getSigners();

    const gameName = "Test Game";
    const gameVersion = "1.1.4";
    const gamePrice = 100;
    const gameDev = {
      name: "Test Game",
      addr: owner.address,
      domain: "tcs1g20.com",
    };

    const Game = await ethers.getContractFactory("Game");
    const game = await Game.deploy(gameName, gameVersion, gamePrice, gameDev);

    return { game, gameName, gameVersion, gameDev, gamePrice };
  }

  describe("deployment", () => {
    it("Should set the right name", async () => {
      const { game, gameName } = await loadFixture(deploySampleGame);
      expect(await game.title()).to.equal(gameName);
    });

    it("Should set the right version", async () => {
      const { game, gameVersion } = await loadFixture(deploySampleGame);
      expect(await game.version()).to.equal(gameVersion);
    });

    it("Should set the right price", async () => {
      const { game, gamePrice } = await loadFixture(deploySampleGame);
      expect(await game.price()).to.equal(gamePrice);
    });

    it("Should set the right developer", async () => {
      const { game, gameDev } = await loadFixture(deploySampleGame);
      expect(await game.developer()).to.have.members([
        gameDev.name,
        gameDev.addr,
        gameDev.domain,
      ]);
    });
  });
});
