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

func assertWinCount(t *testing.T, winCount, expectedWinCount int) {
	t.Helper()

	if winCount != expectedWinCount {
		t.Errorf("Expected %d wins, but got %d", expectedWinCount, winCount)
	}

}
