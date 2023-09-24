package main

import (
	"testing"
)

type TestAlwaysWinsGamer struct{}

func (t TestAlwaysWinsGamer) PlayGame() bool {
	return true
}

func TestIncrementWin(t *testing.T) {
	game := Game{}
	gameCount := 3

	for i := 0; i < gameCount; i++ {
		game.AddWin()
	}

	wantedWinCount := gameCount

	assertWinCount(t, game, wantedWinCount)
}

func TestDetermineWin(t *testing.T) {
	t.Run("Test Synchronous Dynamic Win", func(t *testing.T) {
		game := Game{}
		gamer := TestAlwaysWinsGamer{}
		gameCount := 3

		for i := 0; i < gameCount; i++ {
			didWin := game.DetermineWin(gamer)
			if didWin {
				game.AddWin()
			}
		}

		assertWinCount(t, game, gameCount)
	})

}

func assertWinCount(t testing.TB, game Game, wantedWinCount int) {
	t.Helper()

	if game.Wins() != wantedWinCount {
		t.Errorf("Expected %d wins, but got %d", wantedWinCount, game.Wins())
	}
}
