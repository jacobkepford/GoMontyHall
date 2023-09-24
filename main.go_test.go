package main

import (
	"math/rand"
	"testing"
)

type TestAlwaysWinsGamer struct{}

func (t TestAlwaysWinsGamer) PlayGame() bool {
	return true
}

type TestRandomGamer struct{}

func (t TestRandomGamer) PlayGame() bool {
	return rand.Intn(4) == 1
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

	// t.Run("Test Concurrent Dynamic Win", func(t *testing.T) {
	// 	game := NewGame()
	// 	gamer := TestAlwaysTrueGamer{}
	// 	gameCount := 1000000
	// 	var wg sync.WaitGroup

	// 	wg.Add(gameCount)

	// 	for i := 0; i < gameCount; i++ {
	// 		go func() {
	// 			didWin := game.DetermineWin(gamer)
	// 			if didWin {
	// 				game.AddWin()
	// 			}
	// 			wg.Done()
	// 		}()
	// 	}

	// 	wg.Wait()

	// 	assertWinCount(t, game, gameCount)
	// })
}

// func BenchmarkDynamicWinsIfConcurrentIsFaster(b *testing.B) {
// 	b.Run("Wins One At A Time", func(b *testing.B) {
// 		game := NewGame()
// 		gamer := TestRandomGamer{}
// 		gameCount := 100000

// 		b.ResetTimer()

// 		for i := 0; i < gameCount; i++ {
// 			didWin := game.DetermineWin(gamer)
// 			if didWin {
// 				game.AddWin()
// 			}
// 		}
// 	})

// 	b.Run("Concurrent Dynamic Wins", func(b *testing.B) {
// 		game := NewGame()
// 		gamer := TestRandomGamer{}
// 		gameCount := 100000
// 		var wg sync.WaitGroup

// 		wg.Add(gameCount)

// 		b.ResetTimer()

// 		for i := 0; i < gameCount; i++ {
// 			go func() {
// 				didWin := game.DetermineWin(gamer)
// 				if didWin {
// 					game.AddWin()
// 				}
// 				wg.Done()
// 			}()
// 		}

// 		wg.Wait()

// 	})
// }

func assertWinCount(t testing.TB, game Game, wantedWinCount int) {
	t.Helper()

	if game.Wins() != wantedWinCount {
		t.Errorf("Expected %d wins, but got %d", wantedWinCount, game.Wins())
	}
}
