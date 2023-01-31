package games

// TODO

// stores user's game information on owned and downloading games
type Library struct {
	games     map[[32]byte]*Game
	downloads map[[32]byte]*Download
}
