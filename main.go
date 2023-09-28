package montyHall

import (
	"errors"
	"fmt"
	"math/rand"
)

const prizeAmount int = 3

type gameLogic interface {
	playGame([]string) bool
}

type montyHallLogic struct{}

func (m montyHallLogic) playGame(prizeSet []string) bool {
	userChoice := chooseRandomPrize()
	prizeToShow, err := selectPrizeToShow(prizeSet, userChoice)

	if err != nil {
		return false
	}

	finalPrize := selectSwitchPrize(userChoice, prizeToShow)

	return prizeSet[finalPrize] == "X"
}

type game struct {
	wins      int
	gameCount int
	prizeSets [][]string
}

func (g *game) addWin() {
	g.wins++
}

func (g *game) determineWin(gameLogic gameLogic, prizeSet []string) bool {
	return gameLogic.playGame(prizeSet)
}

func (g *game) configurableRunGame(gameLogic gameLogic) int {
	for _, prizeSet := range g.prizeSets {
		didWin := g.determineWin(gameLogic, prizeSet)
		if didWin {
			g.addWin()
		}
	}

	return g.wins
}

func (g *game) RunGame() int {
	gameLogic := montyHallLogic{}
	return g.configurableRunGame(gameLogic)
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
	prizeNumber := chooseRandomPrize()

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

func chooseRandomPrize() int {
	return rand.Intn(prizeAmount)
}

func selectSwitchPrize(userSelectedPrize, prizeToShow int) int {
	var switchPrize int

	for i := 0; i < prizeAmount; i++ {
		if i != userSelectedPrize && i != prizeToShow {
			switchPrize = i
		}
	}

	return switchPrize
}

func sample() {
	gameAmount := 10000
	game := NewGame(gameAmount)
	wins := game.RunGame()

	winPercent := int(float64(wins) / float64(gameAmount) * 100)

	fmt.Printf("Monty Hall strategy resulted in a %d win rate \n", winPercent)
}
