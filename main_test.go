package main

import (
	"testing"
)

type TestAlwaysWinsGamer struct{}

func (t TestAlwaysWinsGamer) playGame(emptyPrizeSet []string) bool {
	return true
}

func TestIncrementWin(t *testing.T) {
	game := NewGame(3)
	for i := 0; i < game.gameCount; i++ {
		game.addWin()
	}

	assertWinCount(t, game.wins, game.gameCount)
}

func TestPlayGame(t *testing.T) {
	t.Run("Test Positive Number Of Games", func(t *testing.T) {
		game := NewGame(3)
		gameLogic := TestAlwaysWinsGamer{}
		winCount := game.configurableRunGame(gameLogic)

		assertWinCount(t, winCount, game.gameCount)
	})

	t.Run("Test 0 Number Of Games", func(t *testing.T) {
		game := NewGame(0)
		gameLogic := TestAlwaysWinsGamer{}
		winCount := game.configurableRunGame(gameLogic)

		assertWinCount(t, winCount, 0)
	})

	t.Run("Test Negative Number Of Games", func(t *testing.T) {
		game := NewGame(-3)
		gameLogic := TestAlwaysWinsGamer{}
		winCount := game.configurableRunGame(gameLogic)

		assertWinCount(t, winCount, 0)
	})
}

func TestPrizeBoxes(t *testing.T) {
	t.Run("Test single prize set", func(t *testing.T) {
		prizeSet := createPrizeSet()
		var xCount, oCount int

		for _, prize := range prizeSet {
			if prize == "X" {
				xCount++
			} else if prize == "O" {
				oCount++
			}
		}

		if xCount != 1 && oCount != 2 {
			t.Errorf("There should always be one prize (X) but got %d and two goats (O) but got %d", xCount, oCount)
		}
	})

	t.Run("Test game set creation", func(t *testing.T) {
		game := NewGame(5)
		if len(game.prizeSets) != game.gameCount {
			t.Errorf("Expected %d game sets, but got %d", game.gameCount, len(game.prizeSets))
		}
	})
}

func TestShowCorrectGoat(t *testing.T) {
	for i := 0; i < 100; i++ {
		prizeSet := createPrizeSet()
		userChosenPrize := chooseRandomPrize()

		prizeToShow, err := selectPrizeToShow(prizeSet, userChosenPrize)

		if err != nil {
			t.Fatalf("Error: %s", err)
		}

		if prizeSet[prizeToShow] == "X" {
			t.Error("Prize shown was the actual prize")
		} else if userChosenPrize == prizeToShow {
			t.Error("Prize shown was the users selected prize")
		}
	}
}

func TestChooseRandomPrize(t *testing.T) {

	for i := 0; i < 100; i++ {
		userSelectedPrize := chooseRandomPrize()

		if userSelectedPrize < 0 || userSelectedPrize > prizeAmount {
			t.Errorf("Randomly selected user prize is %d but should be in range of 0 and %d", userSelectedPrize, prizeAmount)
		}
	}

}

func TestSwitchPrize(t *testing.T) {
	userSelectedPrize := chooseRandomPrize()
	prizeSet := createPrizeSet()
	prizeToShow, err := selectPrizeToShow(prizeSet, userSelectedPrize)

	if err != nil {
		t.Fatal("Error: ", err)
	}

	finalPrize := selectSwitchPrize(userSelectedPrize, prizeToShow)

	if finalPrize == userSelectedPrize {
		t.Errorf("Prize to switch to should not be the originally chosen prize")
	}

	if finalPrize == prizeToShow {
		t.Errorf("Prize to switch to should not be the goat shown")
	}
}

func assertWinCount(t *testing.T, winCount, expectedWinCount int) {
	t.Helper()

	if winCount != expectedWinCount {
		t.Errorf("Expected %d wins, but got %d", expectedWinCount, winCount)
	}
}
