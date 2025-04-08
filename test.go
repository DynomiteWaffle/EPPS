package main

import (
	"log"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Game struct {
	shader *ebiten.Shader // shader
	sprite *ebiten.Image  // sprite
	pallet *ebiten.Image  // pallet
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, "Hello, World!")

	// set up
	drawSopt := ebiten.DrawRectShaderOptions{}
	// add images
	drawSopt.Images[0] = g.pallet
	drawSopt.Images[1] = g.sprite
	drawSopt.GeoM.Translate(20, 20)

	screen.DrawRectShader(g.sprite.Bounds().Dx(), g.sprite.Bounds().Dy(), g.shader, &drawSopt)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

func main() {
	game := &Game{}
	// init game objects
	bytes, _ := os.ReadFile("/EPPS.kage")

	game.shader, _ = ebiten.NewShader(bytes)
	game.sprite, _, _ = ebitenutil.NewImageFromFile("/sprite.png")
	game.pallet, _, _ = ebitenutil.NewImageFromFile("/pallet.png")

	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Hello, World!")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
