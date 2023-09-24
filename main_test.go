package main

import (
	"testing"
)

type TestAlwaysWinsGamer struct{}

func (t TestAlwaysWinsGamer) PlayGame() bool {
	return true
}

func TestIncrementWin(t *testing.T) {
	game := NewGame(3)
	for i := 0; i < game.gameCount; i++ {
		game.AddWin()
	}

	assertWinCount(t, game)
}

func TestPlayGame(t *testing.T) {
	game := NewGame(3)
	gamer := TestAlwaysWinsGamer{}
	gameCount := 3

	for i := 0; i < gameCount; i++ {
		didWin := game.DetermineWin(gamer)
		if didWin {
			game.AddWin()
		}
	}

	assertWinCount(t, game)
}

func assertWinCount(t testing.TB, game Game) {
	t.Helper()

	if game.Wins() != game.gameCount {
		t.Errorf("Expected %d wins, but got %d", game.gameCount, game.Wins())
	}
}
