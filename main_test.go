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

	if game.getWins() != game.gameCount {
		t.Errorf("Expected %d wins, but got %d", game.gameCount, game.getWins())
	}
}

func TestPlayGame(t *testing.T) {
	game := NewGame(3)
	gameLogic := TestAlwaysWinsGamer{}
	winCount := game.RunGame(gameLogic)

	if winCount != game.gameCount {
		t.Errorf("Expected %d wins, but got %d", game.gameCount, winCount)
	}
}
