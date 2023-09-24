package main

import (
	"errors"
	"math/rand"
)

type GameLogic interface {
	playGame() bool
}

type game struct {
	wins      int
	gameCount int
	prizeSets [][]string
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

func NewGame(gameCount int) game {
	if gameCount < 0 {
		gameCount = 0
	}
	return game{wins: 0, gameCount: gameCount, prizeSets: createAllPrizeSets(gameCount)}
}

func createAllPrizeSets(gameCount int) [][]string {
	prizeSets := make([][]string, gameCount)

	for i := 0; i < gameCount; i++ {
		prizeSets[i] = createPrizeSet()
	}

	return prizeSets
}

func createPrizeSet() []string {
	prizeSet := []string{"O", "O", "O"}
	prizeNumber := rand.Intn(3)

	prizeSet[prizeNumber] = "X"

	return prizeSet
}

func selectPrizeToShow(prizeSet []string, userChosenPrize int) (prizeToShow int, err error) {
	if prizeSet[userChosenPrize] != "X" {
		for index := range prizeSet {
			if index != userChosenPrize && prizeSet[index] != "X" {
				return index, nil
			}
		}
	}

	goatToShow := rand.Intn(2)
	goatCount := -1

	for index, prize := range prizeSet {
		if index == userChosenPrize {
			continue
		} else if prize == "X" {
			continue
		} else if prize == "O" {
			goatCount++
		}

		if goatCount == goatToShow {
			return index, nil
		}
	}

	return 0, errors.New("was unable to find goat to show")

}
