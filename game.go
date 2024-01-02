package main

import (
	"errors"
	"fmt"
	"sync"
	"time"

	tcell "github.com/gdamore/tcell/v2"
	"github.com/gdamore/tcell/v2/views"
)

type Game struct {
	quitq      chan struct{}
	screen     tcell.Screen
	eventq     chan tcell.Event
	errmsg     string
	quitone    sync.Once
	leftView   *views.ViewPort
	centerView *views.ViewPort
	rightView  *views.ViewPort
	grid       *Grid
	sprites    []*Sprite
	active     *Sprite
	pause      bool
	spawnPiece chan bool
	dropPiece  chan bool
	speed      time.Duration

	sync.Mutex
}

func (g *Game) Init() error {
	var err error
	g.screen, err = tcell.NewScreen()
	if err != nil {
		return err
	}

	err = g.screen.Init()
	if err != nil {
		return err
	}

	defaultStyle := tcell.StyleDefault.
		Background(tcell.ColorBlack).
		Foreground(tcell.ColorWhite)

	g.screen.SetStyle(defaultStyle)

	g.screen.EnableMouse()

	screenWidth, screenHeight := g.screen.Size()
	sidebarWidth := (screenWidth - MainWidth) / 2

	g.grid = NewGrid(MainWidth, screenHeight)

	g.leftView = views.NewViewPort(g.screen, 0, 0, sidebarWidth, screenHeight)
	g.centerView = views.NewViewPort(g.screen, sidebarWidth, 0, MainWidth, screenHeight)
	g.rightView = views.NewViewPort(g.screen, sidebarWidth+MainWidth, 0, sidebarWidth, screenHeight)

	g.quitq = make(chan struct{})
	g.eventq = make(chan tcell.Event)
	g.spawnPiece = make(chan bool, 1)
	g.dropPiece = make(chan bool, 1)

	g.speed = 100 * time.Millisecond

	g.Reset()

	return nil
}

func (g *Game) Run() error {
	go g.EventPoller()
	go g.Updater()

loop:
	for {
		g.Draw()
		select {
		case <-g.quitq:
			// log.Printf("quit 79\n")
			break loop
		case <-time.After(time.Millisecond * 10):
		case ev := <-g.eventq:
			g.HandleEvent(ev)
		}
	}

	// Inject a wakeup interrupt
	iev := tcell.NewEventInterrupt(nil)
	g.screen.PostEvent(iev)

	g.screen.Fini()
	// wait for updaters to finish
	if g.errmsg != "" {
		return errors.New(g.errmsg)
	}
	return nil
}

func (g *Game) Error(msg string) {
	g.errmsg = msg
	g.Quit()
}

func (g *Game) Quit() error {
	g.quitone.Do(func() {
		close(g.quitq)
	})
	return nil
}

func (g *Game) Reset() {
	g.grid.Clear()

	g.sprites = []*Sprite{
		NewSprite(),
	}
	g.active = g.sprites[0]
}

func (g *Game) Draw() {
	g.Lock()
	defer g.Unlock()

	g.grid.Clear()
	for _, sprite := range g.sprites {
		sprite.Draw(g.grid)
	}

	g.screen.Clear()
	g.leftView.Fill('.', tcell.StyleDefault.
		Background(tcell.ColorPurple).
		Foreground(tcell.ColorWhite))
	g.rightView.Fill('.', tcell.StyleDefault.
		Background(tcell.ColorPurple).
		Foreground(tcell.ColorWhite))
	g.grid.Draw(g.centerView)
	g.screen.Show()
}

func (g *Game) EventPoller() {
	for {
		select {
		case <-g.quitq:
			// log.Printf("quit 140\n")
			return
		default:
		}
		ev := g.screen.PollEvent()
		if ev == nil {
			return
		}
		select {
		case <-g.quitq:
			// log.Printf("quit 150\n")
			return
		case g.eventq <- ev:
		}
	}
}

func (g *Game) HandleEvent(ev tcell.Event) bool {
	switch ev := ev.(type) {
	case *tcell.EventResize:
		// g.lview.Resize(0, 1, -1, -1)
		// g.sview.Resize(0, 0, -1, 1)
		// g.level.HandleEvent(ev)
	case *tcell.EventKey:
		if ev.Key() == tcell.KeyEscape || ev.Key() == tcell.KeyCtrlC {
			fmt.Printf("Pressing quit\n")
			g.Quit()
			return true
		}
		if ev.Key() == tcell.KeyEnter {
			g.pause = !g.pause
			return true
		}
		if ev.Key() == tcell.KeyUp {
			if !g.pause && g.grid.CanPieceMove(g.active) {
				g.active.NextFrame()
			}
			return true
		}
		if ev.Key() == tcell.KeyDown {
			// g.dropPiece <- true
			g.speed = time.Millisecond * 10
			return true
		}
		if ev.Key() == tcell.KeyLeft {
			if !g.pause && g.grid.CanPieceMoveLeft(g.active) {
				g.active.PosX -= 1
			}
			return true
		}
		if ev.Key() == tcell.KeyRight {
			if !g.pause && g.grid.CanPieceMoveRight(g.active) {
				g.active.PosX += 1
			}
			return true
		}
	}

	return true
}

func (g *Game) Updater() {
	for {
		select {
		case <-g.quitq:
			// log.Printf("quit 204\n")
			return
		case <-time.After(g.speed):
			g.Lock()
			// log.Printf("updating 1\n")
			// g.level.Update(time.Now())

			if !g.pause {
				if g.grid.CanPieceMoveVertically(g.active) {
					g.active.PosY += 1
				} else {
					g.SpawnPiece()
					g.speed = time.Millisecond * 100
				}
			}

			g.Unlock()
		}
	}
}

func (g *Game) SpawnPiece() {
	// log.Printf("spawn piece\n")
	sprite := NewSprite()
	g.sprites = append(g.sprites, sprite)
	g.active = sprite
}
