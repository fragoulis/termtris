package main

import (
	"math/rand"

	tcell "github.com/gdamore/tcell/v2"
)

type Sprite struct {
	PosX     int
	PosY     int
	Rotation int
	Shape    *Shape
	Style    tcell.Style
}

var (
	foreground = tcell.ColorBlack
	styles     = []tcell.Style{
		tcell.StyleDefault.
			Foreground(foreground).
			Background(tcell.ColorLightBlue),
		tcell.StyleDefault.
			Foreground(foreground).
			Background(tcell.ColorLightGreen),
		tcell.StyleDefault.
			Foreground(foreground).
			Background(tcell.ColorLightSalmon),
		tcell.StyleDefault.
			Foreground(foreground).
			Background(tcell.ColorLightCoral),
	}
)

func NewSprite() *Sprite {
	shape := &shapes[rand.Int()%len(shapes)]
	style := styles[rand.Int()%len(styles)]

	return &Sprite{
		PosX:     MainWidth / 2,
		PosY:     0,
		Rotation: rand.Int() % len(shape.Frames),
		Shape:    shape,
		Style:    style,
	}
}

func (s *Sprite) Frame() *Frame {
	return &s.Shape.Frames[s.Rotation]
}

func (s *Sprite) NextFrame() {
	frames := len(s.Shape.Frames)
	s.Rotation += 1
	if s.Rotation >= frames {
		s.Rotation = 0
	}
}

func (s *Sprite) Width() int {
	return s.Frame().Width
}

func (s *Sprite) Height() int {
	return s.Frame().Height
}

func (s *Sprite) Draw(grid *Grid) {
	for y := 0; y < s.Height(); y++ {
		for x := 0; x < s.Width(); x++ {
			i := y*s.Width() + x
			if s.Frame().Data[i] == 0 {
				continue
			}
			grid.SetOccupied(s.PosX+x, s.PosY+y, s.Style)
		}
	}
}
