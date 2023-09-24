package main

import (
	"testing"
)

type TestAlwaysWinsGamer struct{}

func (t TestAlwaysWinsGamer) playGame() bool {
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
		winCount := game.RunGame(gameLogic)

		assertWinCount(t, winCount, game.gameCount)
	})

	t.Run("Test 0 Number Of Games", func(t *testing.T) {
		game := NewGame(0)
		gameLogic := TestAlwaysWinsGamer{}
		winCount := game.RunGame(gameLogic)

		assertWinCount(t, winCount, 0)
	})

	t.Run("Test Negative Number Of Games", func(t *testing.T) {
		game := NewGame(-3)
		gameLogic := TestAlwaysWinsGamer{}
		winCount := game.RunGame(gameLogic)

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

}

func assertWinCount(t *testing.T, winCount, expectedWinCount int) {
	t.Helper()

	if winCount != expectedWinCount {
		t.Errorf("Expected %d wins, but got %d", expectedWinCount, winCount)
	}
}
