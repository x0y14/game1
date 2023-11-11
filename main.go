package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"image/color"
	"log"
)

const (
	screenWidth  = 640
	screenHeight = 640
)

type Game struct {
	count   int
	player  *Object
	objects []*Object
}

func (g *Game) Update() error {
	g.count++

	if inpututil.IsKeyJustPressed(ebiten.KeyW) {
		g.player.SetPosition(Position{
			X: g.player.Center.X,
			Y: g.player.Center.Y - 3,
		})
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyS) {
		g.player.SetPosition(Position{
			X: g.player.Center.X,
			Y: g.player.Center.Y + 3,
		})
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.drawObject(screen)
	g.drawPlayer(screen)
	ebitenutil.DebugPrint(screen, fmt.Sprintf("TPS: %0.2f", ebiten.ActualTPS()))
}
func (g *Game) drawPlayer(screen *ebiten.Image) {
	for _, pol := range g.player.Component {
		//var realX, realY int
		switch pol.Figure {
		case Rectangle:
			vector.DrawFilledRect(
				screen,
				g.player.Center.X+pol.Center.X-pol.Width/2,
				g.player.Center.Y+pol.Center.Y-pol.Height/2,
				pol.Width,
				pol.Height,
				pol.Color,
				true,
			)
		case Circle:
			vector.DrawFilledCircle(
				screen,
				g.player.Center.X+pol.Center.X,
				g.player.Center.Y+pol.Center.Y,
				pol.Radius,
				pol.Color,
				true,
			)
		}
	}
}
func (g *Game) drawObject(screen *ebiten.Image) {
	for _, obj := range g.objects {
		for _, pol := range obj.Component {
			switch pol.Figure {
			case Rectangle:
				vector.DrawFilledRect(
					screen,
					obj.Center.X+pol.Center.X,
					obj.Center.Y+pol.Center.Y,
					pol.Width,
					pol.Height,
					pol.Color,
					true,
				)
			case Circle:
				vector.DrawFilledCircle(
					screen,
					obj.Center.X+pol.Center.X,
					obj.Center.Y+pol.Center.Y,
					pol.Radius,
					pol.Color,
					true,
				)
			}
		}
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Game1")
	game := &Game{}

	// head-part
	head := &Polygon{
		Figure: Circle,
		Center: Position{X: 0, Y: -10},
		Radius: 10,
		Color:  color.RGBA{R: 0x00, G: 0x80, B: 0x00, A: 0x80},
	}
	// body-part
	body := &Polygon{
		Figure: Rectangle,
		Center: Position{X: 0, Y: 20},
		Width:  20,
		Height: 40,
		Color:  color.RGBA{R: 0x00, G: 0x80, B: 0x00, A: 0x80},
	}
	human := &Object{
		Component: []*Polygon{head, body},
		Center:    Position{X: 0, Y: 0},
	}
	// register character to game
	//game.objects = append(game.objects, human)
	// move character to center
	human.SetPosition(Position{X: screenWidth / 2, Y: screenHeight / 2})
	game.player = human

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
