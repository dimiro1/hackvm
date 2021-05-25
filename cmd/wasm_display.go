// +build js,wasm

package main

import (
	"fmt"
	"hackvm"
	"syscall/js"
)

// display implements a HackVM screen rendering on HTML canvas element.
type display struct {
	ram           []int
	canvas        js.Value
	ctx2d         js.Value
	imageData     js.Value
	imageDataData js.Value
	buffer        []byte
}

// newDisplay returns a new HTMLCanvasDisplay.
func newDisplay() *display {
	canvas := js.Global().Get("document").Call("createElement", "canvas")
	canvas.Set("width", hackvm.DisplayWidth)
	canvas.Set("height", hackvm.DisplayHeight)

	ctx2d := canvas.Call("getContext", "2d")
	imageData := ctx2d.Call("createImageData", hackvm.DisplayWidth, hackvm.DisplayHeight)
	imageDataData := imageData.Get("data")

	return &display{
		ram:           make([]int, 0x2000),
		canvas:        canvas,
		ctx2d:         ctx2d,
		imageData:     imageData,
		imageDataData: imageDataData,
		buffer:        make([]byte, imageDataData.Length()),
	}
}

// Reset set memory values to 0.
func (d *display) Reset() error {
	for i := 0; i < len(d.ram); i++ {
		if err := d.WriteWord(i, 0); err != nil {
			return err
		}
	}

	return nil
}

func (d *display) ReadWord(address int) (int, error) {
	if address < 0 || address > len(d.ram) {
		return 0, hackvm.ErrBadAddress
	}

	word := d.ram[address]
	return word, nil
}

func (d *display) WriteWord(address int, word int) error {
	if address < 0 || address > len(d.ram) {
		return hackvm.ErrBadAddress
	}

	d.ram[address] = word
	return nil
}

// Refresh renders the screen data on the HTML canvas element.
func (d *display) Refresh() error {
	for row := 0; row < hackvm.DisplayHeight; row++ {
		for col := 0; col < hackvm.DisplayWidth; col++ {
			addr := 32*row + col/16
			word, err := d.ReadWord(addr)
			if err != nil {
				return fmt.Errorf("refreshing screen: %w", err)
			}

			var (
				nth   = col % 16
				bit   = (word&(1<<nth))>>nth == 1
				red   = row*(hackvm.DisplayWidth*4) + col*4
				green = red + 1
				blue  = red + 2
				alpha = red + 3
			)

			if bit {
				d.buffer[red] = 0x75
				d.buffer[green] = 0xf9
				d.buffer[blue] = 0x4c
			} else {
				d.buffer[red] = 0x00
				d.buffer[green] = 0x00
				d.buffer[blue] = 0x00
			}

			d.buffer[alpha] = 0xff
		}
	}

	js.CopyBytesToJS(d.imageDataData, d.buffer)
	d.ctx2d.Call("putImageData", d.imageData, 0, 0)
	return nil
}

// Canvas returns a HTML canvas element.
func (d *display) Canvas() js.Value {
	return d.canvas
}
