package main

import (
	"math/rand"
)

type GameLogic interface {
	playGame() bool
}

type game struct {
	wins       int
	gameCount  int
	prizeGroup [][]string
}

func (g *game) addWin() {
	g.wins++
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

	return g.wins
}

func (g *game) createPrizeSet() []string {
	prizeSet := []string{"O", "O", "O"}
	prizeNumber := rand.Intn(3)

	prizeSet[prizeNumber] = "X"

	return prizeSet
}

func NewGame(gameCount int) game {
	return game{wins: 0, gameCount: gameCount}
}
