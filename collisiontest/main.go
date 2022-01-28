package main

import (
	"image/color"
	"log"

	"collisiontest/component"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

// Game implements ebiten.Game interface.
type Game struct {
	ball *component.Square
	w    int
	h    int
}

// Update proceeds the game state.
// Update is called every tick (1/60 [s] by default).
func (g *Game) Update() error {
	// Write your game's logical update.

	if ebiten.IsKeyPressed(ebiten.KeySpace) {
		g.ball.IsRun = !g.ball.IsRun
	}

	if g.ball.IsRun {
		g.ball.CollisionDetection(float64(g.w), float64(g.h))
	}
	return nil
}

// Draw draws the game screen.
// Draw is called every frame (typically 1/60[s] for 60Hz display).
func (g *Game) Draw(screen *ebiten.Image) {

	screen.DrawImage(g.ball.Image, g.ball.Opts)

	if !g.ball.IsAlive() {
		ebitenutil.DebugPrint(screen, "遊戲結束")
	} else {
		ebitenutil.DebugPrint(screen, "目前座標:"+g.ball.GetCoordinate())
	}
	// Write your game's rendering.
}

// Layout takes the outside size (e.g., the window size) and returns the (logical) screen size.
// If you don't have to adjust the screen size with the outside size, just return a fixed size.
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return g.w, g.h
}

var (
	Red = color.RGBA{0xff, 0x00, 0x00, 0xff}
)

func main() {
	game := &Game{
		ball: component.NewSquare(Red, 5, 5, 157.5, 225, 3),
		w:    320,
		h:    240,
	}
	// Specify the window size as you like. Here, a doubled size is specified.
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("碰撞小測試")
	// Call ebiten.RunGame to start your game loop.
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
