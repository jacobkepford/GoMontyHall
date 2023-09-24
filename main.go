package main

type Gamer interface {
	PlayGame() bool
}

type Game struct {
	wins int
}

func (g *Game) AddWin() {
	g.wins++
}

func (g *Game) Wins() int {
	return g.wins
}

func (g *Game) DetermineWin(gamer Gamer) bool {
	return gamer.PlayGame()
}

func NewGame() *Game {
	return &Game{}
}
