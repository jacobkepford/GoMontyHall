package main

type GameLogic interface {
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

func (g game) determineWin(gameLogic GameLogic) bool {
	return gameLogic.playGame()
}

func (g *game) RunGame(gameLogic GameLogic) int {
	for i := 0; i < g.gameCount; i++ {
		didWin := g.determineWin(gameLogic)
		if didWin {
			g.addWin()
		}
	}

	return g.getWins()
}

func NewGame(gameCount int) game {
	return game{wins: 0, gameCount: gameCount}
}
