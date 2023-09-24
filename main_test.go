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

	assertWinCount(t, game)
}

func TestPlayGame(t *testing.T) {
	game := NewGame(3)
	gamer := TestAlwaysWinsGamer{}

	for i := 0; i < game.gameCount; i++ {
		didWin := game.determineWin(gamer)
		if didWin {
			game.addWin()
		}
	}

	assertWinCount(t, game)
}

func assertWinCount(t testing.TB, game game) {
	t.Helper()

	if game.getWins() != game.gameCount {
		t.Errorf("Expected %d wins, but got %d", game.gameCount, game.getWins())
	}
}
