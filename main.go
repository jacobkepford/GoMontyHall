package main

type Gamer interface {
	playGame() bool
}

type game struct {
	wins      int
	gameCount int
}

func (g *game) addWin() {
	g.wins++
}

func (g game) getWins() int {
	return g.wins
}

func (g game) determineWin(gamer Gamer) bool {
	return gamer.playGame()
}

func NewGame(gameCount int) game {
	return game{wins: 0, gameCount: gameCount}
}
