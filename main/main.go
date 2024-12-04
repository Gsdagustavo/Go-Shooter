package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Game struct{}

func (g *Game) Update() error {
	// Update game logic here
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// Draw game objects here
	ebitenutil.DebugPrint(screen, "Hello, Ebiten!")
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return width, height // Game window size
}

const width int = 1280
const height int = 720

func main() {
	game := &Game{}
	ebiten.SetWindowSize(width, height)
	ebiten.SetWindowTitle("My Game in Go")
	if err := ebiten.RunGame(game); err != nil {
		panic(err)
	}
}
