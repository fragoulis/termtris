package main

import (
	"fmt"
	"log"
	"os"
)

const (
	Cell      = '█'
	MainWidth = 30
)

// type Animation struct {
// 	Ticker   *time.Ticker
// 	Duration time.Duration
// 	Piece    *Sprite
// }

// func NewAnimation(duration time.Duration, piece *Sprite) *Animation {
// 	return &Animation{
// 		Duration: duration,
// 		Piece:    piece,
// 	}
// }

// func (a *Animation) Start() {
// 	a.Ticker = time.NewTicker(a.Duration)

// 	go func() {
// 		for {
// 			select {
// 			case <-a.Ticker.C:
// 				if g.CanPieceMoveVertically(a.Piece) {
// 					a.Piece.PosY += 1
// 				} else {
// 					a.Ticker.Stop()
// 				}
// 			}
// 		}
// 	}()
// }

// func (a *Animation) Stop() {
// 	a.Ticker.Stop()
// }

func main() {
	if f, e := os.Create("debug.log"); e == nil {
		log.SetOutput(f)
	}

	game := &Game{}
	if err := game.Init(); err != nil {
		fmt.Printf("Failed to initialize game: %v\n", err)
		os.Exit(1)
	}
	if err := game.Run(); err != nil {
		fmt.Printf("Failed to run game: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Thanks for playing!\n")
}
