package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"hackvm"
	"log"
	"os"
)

var specialKeys = map[ebiten.Key]struct{}{
	ebiten.KeyLeft:      {},
	ebiten.KeyRight:     {},
	ebiten.KeyUp:        {},
	ebiten.KeyDown:      {},
	ebiten.KeyHome:      {},
	ebiten.KeyEnd:       {},
	ebiten.KeyEnter:     {},
	ebiten.KeyBackspace: {},
	ebiten.KeyPageUp:    {},
	ebiten.KeyPageDown:  {},
	ebiten.KeyInsert:    {},
	ebiten.KeyDelete:    {},
	ebiten.KeyEscape:    {},
	ebiten.KeyF1:        {},
	ebiten.KeyF2:        {},
	ebiten.KeyF3:        {},
	ebiten.KeyF4:        {},
	ebiten.KeyF5:        {},
	ebiten.KeyF6:        {},
	ebiten.KeyF7:        {},
	ebiten.KeyF8:        {},
	ebiten.KeyF9:        {},
	ebiten.KeyF10:       {},
	ebiten.KeyF11:       {},
	ebiten.KeyF12:       {},
}

type Runner struct {
	computer *hackvm.Computer
}

func (r *Runner) stepComputer() error {
	for i := 0; i < 50000; i++ {
		if err := r.computer.Step(); err != nil {
			return err
		}
	}
	return nil
}

func (r *Runner) updateKeys() error {
	r.computer.KeyUp()

	for _, key := range ebiten.AppendInputChars(nil) {
		r.computer.KeyDown(string(key))
	}

	for _, key := range inpututil.AppendPressedKeys(nil) {
		if _, ok := specialKeys[key]; ok {
			r.computer.KeyDown(key.String())
		}
	}
	return nil
}

func (r *Runner) Update() error {
	if err := r.stepComputer(); err != nil {
		return err
	}

	if err := r.updateKeys(); err != nil {
		return err
	}

	return nil
}

func (r *Runner) Draw(screen *ebiten.Image) {
	buffer, err := r.computer.ScreenBuffer()
	if err != nil {
		panic(err)
	}
	screen.DrawImage(ebiten.NewImageFromImage(buffer), &ebiten.DrawImageOptions{})
}

func (r *Runner) Layout(_, _ int) (int, int) {
	return hackvm.DisplayWidth, hackvm.DisplayHeight
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("usage: program path/to/rom.hack")
		return
	}

	ebiten.SetWindowSize(hackvm.DisplayWidth, hackvm.DisplayHeight)
	ebiten.SetWindowTitle("HackVm")
	ebiten.SetMaxTPS(30)

	var computer = hackvm.NewComputer()
	if err := computer.LoadRomFromFile(os.Args[1]); err != nil {
		panic(err)
	}

	var runner = &Runner{
		computer: computer,
	}

	if err := ebiten.RunGame(runner); err != nil {
		log.Fatal(err)
	}
}
