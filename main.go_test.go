package main

import (
	"sync"
	"testing"
)

func TestIncrementWin(t *testing.T) {
	t.Run("Test Incrementing Wins One At A Time", func(t *testing.T) {
		game := NewGame()
		gameCount := 1000000

		for i := 0; i < gameCount; i++ {
			game.AddWin()
		}

		wantedWinCount := gameCount

		assertWinCount(t, game, wantedWinCount)

	})

	t.Run("Test Incrementing Wins Concurrently", func(t *testing.T) {
		game := NewGame()
		gameCount := 1000000
		var wg sync.WaitGroup

		wg.Add(gameCount)

		for i := 0; i < gameCount; i++ {
			go func() {
				game.AddWin()
				wg.Done()
			}()
		}

		wg.Wait()

		assertWinCount(t, game, gameCount)
	})

}

func BenchmarkWins(b *testing.B) {
	b.Run("Test Wins One At A Time", func(b *testing.B) {
		game := NewGame()
		gameCount := 1000000

		b.ResetTimer()

		for i := 0; i < gameCount; i++ {
			game.AddWin()
		}
	})

	b.Run("Test Wins Concurrently", func(b *testing.B) {
		game := NewGame()
		gameCount := 1000000
		var wg sync.WaitGroup

		wg.Add(gameCount)

		b.ResetTimer()

		for i := 0; i < gameCount; i++ {
			go func() {
				game.AddWin()
				wg.Done()
			}()
		}

		wg.Wait()
	})
}

func assertWinCount(t testing.TB, game *Game, wantedWinCount int) {
	t.Helper()

	if game.Wins() != wantedWinCount {
		t.Errorf("Expected %d wins, but got %d", wantedWinCount, game.Wins())
	}
}
