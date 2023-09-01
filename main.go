package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"

	tcell "github.com/gdamore/tcell/v2"
	"github.com/gdamore/tcell/v2/views"
	termbox "github.com/nsf/termbox-go"
	"golang.org/x/sys/unix"
)

func main() {
	// fmt.Print("\033[2J")
	// fmt.Print("\033[;H")

	runTcell()
	// readIn()
}

func printGreen() {
	// ANSI escape codes for green color
	const greenColor = "\033[32m"
	const resetColor = "\033[0m"

	message := "Hello, World!"
	// Print the message in green color
	fmt.Println(greenColor + message + resetColor)
}

func printSpinner() {
	loadingWheel := []string{"|", "/", "-", "\\"}
	for i := 0; i < 2; i++ {
		for _, symbol := range loadingWheel {
			fmt.Printf("\r%s Loading... %s", symbol, symbol)
			time.Sleep(100 * time.Millisecond)
		}
	}
	fmt.Print("\r\033[0K")
	fmt.Println("Done!")
}

func printProgressBar() {
	total := 100  // total number of items to process
	progress := 0 // current progress

	for progress <= total {
		// Calculate the percentage completion
		percentage := float64(progress) / float64(total) * 100

		// Render the progress bar
		fmt.Printf("\rProgress: [%-50s] %.1f%%", progressBar(percentage), percentage)

		time.Sleep(10 * time.Millisecond)

		// Update progress
		progress++
	}

	fmt.Print("\r\033[0K")
	fmt.Println("Done!")
}

func progressBar(percentage float64) string {
	barLength := 50
	completeLength := int(percentage / 2)
	var progressBar string

	for i := 0; i < completeLength; i++ {
		progressBar += "="
	}

	for i := 0; i < barLength-completeLength; i++ {
		progressBar += " "
	}

	return progressBar
}

func printColors() {
	fmt.Print("\033[?25l")
	for i := 0; i <= 10; i++ {
		fmt.Print("\033[2J\033[0;0H")
		fmt.Printf("\033[38;5;%dmColor %d\033[0m", i, i)
		time.Sleep(time.Second)
	}
	fmt.Print("\033[?25h")
}

func runTermbox() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()

	eventQueue := make(chan termbox.Event)
	go func() {
		for {
			eventQueue <- termbox.PollEvent()
		}
	}()

	for {
		select {
		case ev := <-eventQueue:
			if ev.Type == termbox.EventKey {
				switch ev.Key {
				case termbox.KeyArrowUp:
					fmt.Println("Up arrow key pressed")
				case termbox.KeyArrowDown:
					fmt.Println("Down arrow key pressed")
				case termbox.KeyArrowLeft:
					fmt.Println("Left arrow key pressed")
				case termbox.KeyArrowRight:
					fmt.Println("Right arrow key pressed")
				}
			}
		default:
		}
	}
}

type Shape struct {
	Frames []Frame
}

type Frame struct {
	Width  int
	Height int
	Data   []rune
}

var shapes = []Shape{
	Shape{
		Frames: []Frame{
			Frame{
				Width:  3,
				Height: 2,
				Data: []rune{
					0, '█', '█',
					'█', '█', 0,
				},
			},
			Frame{
				Width:  2,
				Height: 3,
				Data: []rune{
					'█', 0,
					'█', '█',
					0, '█',
				},
			},
			Frame{
				Width:  3,
				Height: 2,
				Data: []rune{
					'█', '█', 0,
					0, '█', '█',
				},
			},
			Frame{
				Width:  2,
				Height: 3,
				Data: []rune{
					0, '█',
					'█', '█',
					'█', 0,
				},
			},
		},
	},
	Shape{
		Frames: []Frame{
			Frame{
				Width:  3,
				Height: 2,
				Data: []rune{
					0, '█', 0,
					'█', '█', '█',
				},
			},
			Frame{
				Width:  2,
				Height: 3,
				Data: []rune{
					'█', 0,
					'█', '█',
					'█', 0,
				},
			},
			Frame{
				Width:  3,
				Height: 2,
				Data: []rune{
					'█', '█', '█',
					0, '█', 0,
				},
			},
			Frame{
				Width:  2,
				Height: 3,
				Data: []rune{
					0, '█',
					'█', '█',
					0, '█',
				},
			},
		},
	},
	Shape{
		Frames: []Frame{
			Frame{
				Width:  3,
				Height: 2,
				Data: []rune{
					'█', 0, 0,
					'█', '█', '█',
				},
			},
			Frame{
				Width:  2,
				Height: 3,
				Data: []rune{
					'█', '█',
					'█', 0,
					'█', 0,
				},
			},
			Frame{
				Width:  3,
				Height: 2,
				Data: []rune{
					'█', '█', '█',
					0, 0, '█',
				},
			},
			Frame{
				Width:  2,
				Height: 3,
				Data: []rune{
					0, '█',
					0, '█',
					'█', '█',
				},
			},
		},
	},
	Shape{
		Frames: []Frame{
			Frame{
				Width:  4,
				Height: 1,
				Data: []rune{
					'█', '█', '█', '█',
				},
			},
			Frame{
				Width:  1,
				Height: 4,
				Data: []rune{
					'█',
					'█',
					'█',
					'█',
				},
			},
		},
	},
	Shape{
		Frames: []Frame{
			Frame{
				Width:  2,
				Height: 2,
				Data: []rune{
					'█', '█',
					'█', '█',
				},
			},
		},
	},
}

type Sprite struct {
	PosX     int
	PosY     int
	Rotation int
	Shape    *Shape
}

func NewSprite(x int) *Sprite {
	shape := &shapes[rand.Int()%4]
	return &Sprite{
		PosX:     x,
		PosY:     0,
		Rotation: rand.Int() % len(shape.Frames),
		Shape:    shape,
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

func (s *Sprite) Draw(v *views.ViewPort) {
	for y := 0; y < s.Frame().Height; y++ {
		for x := 0; x < s.Frame().Width; x++ {
			i := y*s.Frame().Width + x
			if s.Frame().Data[i] == 0 {
				continue
			}
			style := tcell.StyleDefault
			v.SetContent(x+s.PosX, y+s.PosY, s.Frame().Data[i], nil, style)
		}
	}
}

func detectCollision(a, b *Sprite) bool {
	if a.PosX < b.PosX+b.Width() &&
		a.PosX+a.Width() > b.PosX &&
		a.PosY < b.PosY+b.Height() &&
		a.PosY+a.Height() > b.PosY {
		return true
	}

	return false
}

func collisionDetection(active *Sprite, sprites []*Sprite) bool {
	for _, sprite := range sprites {
		if sprite == active {
			continue
		}

		return detectCollision(active, sprite)
	}

	return false
}

func runTcell() {
	screen, err := tcell.NewScreen()
	if err != nil {
		panic(err)
	}
	defer screen.Fini()

	err = screen.Init()
	if err != nil {
		panic(err)
	}

	// screen.SetCursor(0, 0)
	// screen.ShowCursor(2, 2)

	screen.SetStyle(tcell.StyleDefault.
		Foreground(tcell.ColorWhite).
		Background(tcell.ColorBlack))

	screen.EnableMouse()

	w, h := screen.Size()
	thirdW := w / 3

	sprites := []*Sprite{
		NewSprite(thirdW / 2),
	}
	active := sprites[0]

	left := views.NewViewPort(screen, 0, 0, thirdW, h)
	center := views.NewViewPort(screen, thirdW, 0, thirdW, h)
	right := views.NewViewPort(screen, thirdW*2, 0, thirdW, h)
	screen.Show()

	eventQueue := make(chan tcell.Event)
	go func() {
		for {
			event := screen.PollEvent()
			eventQueue <- event
		}
	}()

	sigterm := make(chan os.Signal, 1)
	signal.Notify(sigterm, syscall.SIGINT, syscall.SIGTERM)

	speed := 100 * time.Millisecond
	ticker := time.NewTicker(speed)
	pause := false

	go func() {
		for {
			select {
			case <-ticker.C:
				if !pause && active.PosY < h-active.Frame().Height && !collisionDetection(active, sprites) {
					active.PosY += 1
				}
			}
		}
	}()

	spawn := make(chan bool, 1)
	go func() {
		for {
			select {
			case <-spawn:
				active.PosY = h - active.Frame().Height
				sprite := NewSprite(thirdW / 2)
				sprites = append(sprites, sprite)
				active = sprite
			}
		}
	}()

	for {
		select {
		case ev := <-eventQueue:
			switch ev := ev.(type) {
			case *tcell.EventKey:
				switch ev.Key() {
				case tcell.KeyUp:
					// fmt.Print("\033[1A")
					// active.PosY -= 1
					if !pause && active.PosY < h-active.Height() {
						active.NextFrame()
					}
				case tcell.KeyDown:
					// fmt.Print("\033[1B")
					// active.PosY += 1
					spawn <- true
				case tcell.KeyLeft:
					// fmt.Print("\033[1D")
					if !pause && active.PosX > 0 && !collisionDetection(active, sprites) {
						active.PosX -= 1
					}
				case tcell.KeyRight:
					// fmt.Print("\033[1C")
					if !pause && active.PosX < thirdW-active.Width() && !collisionDetection(active, sprites) {
						active.PosX += 1
					}
				case tcell.KeyEnter:
					pause = !pause
				case tcell.KeyCtrlC:
					ticker.Stop()
					sigterm <- syscall.SIGTERM
				case tcell.KeyEscape:
					ticker.Stop()
					sigterm <- syscall.SIGTERM
				}
			}
		case <-sigterm:
			return
		default:
			screen.Clear()
			left.Fill('x', tcell.StyleDefault)
			right.Fill('.', tcell.StyleDefault)

			for _, sprite := range sprites {
				sprite.Draw(center)
			}
			screen.Show()
		}
	}
}

func readIn() {
	reader := bufio.NewReader(os.Stdin)

	for {
		// Read a single rune from stdin
		char, _, err := reader.ReadRune()
		if err != nil {
			fmt.Println("Error reading input:", err)
			os.Exit(1)
		}

		// Check for arrow keys
		switch char {
		case '\x1b': // Escape key
			// Read two more runes to check for arrow keys
			char2, _, err := reader.ReadRune()
			if err != nil {
				fmt.Println("Error reading input:", err)
				os.Exit(1)
			}
			char3, _, err := reader.ReadRune()
			if err != nil {
				fmt.Println("Error reading input:", err)
				os.Exit(1)
			}

			if char2 == '[' {
				// Arrow key detected
				switch char3 {
				case 'A':
					fmt.Println("Up arrow key pressed")
				case 'B':
					fmt.Println("Down arrow key pressed")
				case 'C':
					fmt.Println("Right arrow key pressed")
				case 'D':
					fmt.Println("Left arrow key pressed")
				}
			}
		default:
			fmt.Printf("Key pressed:%#v\n", string(char))
		}
	}
}

func runRawTerm() {
	oldState, err := makeRawTerminal()
	if err != nil {
		fmt.Println("Error setting terminal to raw mode:", err)
		os.Exit(1)
	}
	defer restoreTerminal(oldState)

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)

	var b [3]byte
	for {
		select {
		case <-ch:
			// Exit the program when Ctrl+C is pressed
			fmt.Println("Program terminated by Ctrl+C")
			os.Exit(1)
		default:
			n, err := os.Stdin.Read(b[:])
			if err != nil {
				fmt.Println("Error reading input:", err)
				os.Exit(1)
			}

			// Check for arrow keys
			if n == 3 && b[0] == 0x1b && b[1] == '[' {
				switch b[2] {
				case 'A':
					// fmt.Println("Up arrow key pressed")
					fmt.Print("\033[1A")
				case 'B':
					// fmt.Println("Down arrow key pressed")
					fmt.Print("\033[1B")
				case 'C':
					// fmt.Println("Right arrow key pressed")
					fmt.Print("\033[1C")
				case 'D':
					// fmt.Println("Left arrow key pressed")
					fmt.Print("\033[1D")
				}
			}
		}
	}
}

// makeRawTerminal puts the terminal into raw mode and returns the old terminal state
func makeRawTerminal() (*unix.Termios, error) {
	oldState, err := unix.IoctlGetTermios(syscall.Stdin, unix.TCGETS)
	if err != nil {
		return nil, err
	}

	newState := *oldState
	newState.Lflag &^= unix.ICANON | unix.ECHO | unix.ISIG
	err = unix.IoctlSetTermios(syscall.Stdin, unix.TCSETS, &newState)
	if err != nil {
		return nil, err
	}

	return oldState, nil
}

// restoreTerminal restores the terminal to its original state
func restoreTerminal(oldState *unix.Termios) {
	unix.IoctlSetTermios(syscall.Stdin, unix.TCSETS, oldState)
}
