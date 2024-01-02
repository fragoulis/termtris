package main

import (
	tcell "github.com/gdamore/tcell/v2"
	"github.com/gdamore/tcell/v2/views"
)

type GridBlock struct {
	X        int
	Y        int
	Style    tcell.Style
	Occupied bool
}

type Grid struct {
	Width  int
	Height int
	Data   []GridBlock
}

func NewGrid(w, h int) *Grid {
	data := make([]GridBlock, w*h)
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			data[y*w+x].X = x
			data[y*w+x].Y = y
		}
	}

	return &Grid{
		Width:  w,
		Height: h,
		Data:   data,
	}
}
func (g *Grid) SetOccupied(x, y int, style tcell.Style) {
	block := &g.Data[y*g.Width+x]
	block.Occupied = true
	block.Style = style
}

func (g *Grid) Clear() {
	for i := range g.Data {
		g.Data[i].Occupied = false
	}
}

func (g *Grid) Draw(v *views.ViewPort) {
	for y := 0; y < g.Height; y++ {
		for x := 0; x < g.Width; x++ {
			i := y*g.Width + x
			block := &g.Data[i]

			// Debug the position
			// value := []rune(strconv.Itoa(i))[0]
			value := ' '

			if !block.Occupied {
				v.SetContent(x, y, value, nil, tcell.StyleDefault.
					Foreground(tcell.ColorWhite).
					Background(tcell.ColorBlack))
				continue
			}

			v.SetContent(x, y, value, nil, block.Style)
		}
	}
}

func (g *Grid) CanPieceMoveLeft(piece *Sprite) bool {
	return g.CanPieceMoveVertically(piece) &&
		piece.PosX > 0
}

func (g *Grid) CanPieceMoveRight(piece *Sprite) bool {
	return g.CanPieceMoveVertically(piece) &&
		piece.PosX < g.Width-piece.Width()
}

func (g *Grid) CanPieceMoveLaterally(piece *Sprite) bool {
	return g.CanPieceMoveLeft(piece) && g.CanPieceMoveRight(piece)
}

func (g *Grid) GetBottomRow(piece *Sprite) []GridBlock {
	from := (piece.PosY+piece.Height()-1)*g.Width + piece.PosX
	to := from + piece.Width()
	return g.Data[from:to]
}

func (g *Grid) GetBlock(x, y int) *GridBlock {
	index := y*g.Width + x
	if index >= len(g.Data) {
		return nil
	}

	return &g.Data[index]
}

func (g *Grid) CanPieceMoveVertically(piece *Sprite) bool {
	blocks := g.GetBottomRow(piece)
	for _, block := range blocks {
		blockBelow := g.GetBlock(block.X, block.Y+1)
		if blockBelow == nil {
			continue
		}
		if blockBelow.Occupied {
			return false
		}
	}

	return piece.PosY < g.Height-piece.Height()
}

func (g *Grid) CanPieceMove(piece *Sprite) bool {
	return g.CanPieceMoveLaterally(piece) || g.CanPieceMoveVertically(piece)
}
