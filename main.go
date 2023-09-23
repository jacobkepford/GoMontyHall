package main

import "sync"

type Game struct {
	mu   sync.Mutex
	wins int
}

func (g *Game) AddWin() {
	g.mu.Lock()
	defer g.mu.Unlock()
	g.wins++
}

func (g *Game) Wins() int {
	return g.wins
}

func NewGame() *Game {
	return &Game{}
}
