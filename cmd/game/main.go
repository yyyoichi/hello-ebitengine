package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Game struct {
	word string
}

func (g *Game) Update() error {
	resp, err := http.Get("/api/echo")
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status: %s", resp.Status)
	}
	var buf bytes.Buffer
	buf.ReadFrom(resp.Body)
	g.word = buf.String()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, g.word)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Hello, World!!")
	ebiten.SetTPS(10)
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
