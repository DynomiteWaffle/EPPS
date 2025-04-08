package main

import (
	_ "embed"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var (

	//go:embed EPPS.kage
	EPPS []byte
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
	screen.Fill(color.White)

	// set up
	drawSopt := ebiten.DrawRectShaderOptions{}
	// add images
	drawSopt.Images[0] = g.pallet
	drawSopt.Images[1] = g.sprite
	drawSopt.GeoM.Translate(0, 16)

	screen.DrawRectShader(g.sprite.Bounds().Dx(), g.sprite.Bounds().Dy(), g.shader, &drawSopt)

	drawopt := &ebiten.DrawImageOptions{}
	drawopt.GeoM.Translate(16, 16)

	screen.DrawImage(g.sprite, drawopt)
	drawopt.GeoM.Translate(16, 0)
	screen.DrawImage(g.pallet, drawopt)

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 16 * 5, 16 * 5
}

func main() {
	game := &Game{}
	// init game objects
	var err error
	game.shader, err = ebiten.NewShader(EPPS)
	if err != nil {
		log.Fatal(err)
	}

	game.sprite, _, err = ebitenutil.NewImageFromFile("sprite.png")
	game.pallet, _, err = ebitenutil.NewImageFromFile("pallet.png")
	if err != nil {
		log.Fatal(err)
	}

	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Hello, World!")

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
