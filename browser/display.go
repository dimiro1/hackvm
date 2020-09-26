// +build js,wasm

package browser

import (
	"fmt"
	"syscall/js"

	"hackvm"
)

const (
	DisplayWidth  = 512
	DisplayHeight = 256
)

// HTMLCanvasDisplay implements a HackVM screen rendering on HTML canvas element.
type HTMLCanvasDisplay struct {
	ram           []int
	canvas        js.Value
	ctx2d         js.Value
	imageData     js.Value
	imageDataData js.Value
	buffer        []byte
}

// NewHTMLCanvasDisplay returns a new HTMLCanvasDisplay.
func NewHTMLCanvasDisplay() *HTMLCanvasDisplay {
	canvas := js.Global().Get("document").Call("createElement", "canvas")
	canvas.Set("width", DisplayWidth)
	canvas.Set("height", DisplayHeight)

	ctx2d := canvas.Call("getContext", "2d")
	imageData := ctx2d.Call("createImageData", DisplayWidth, DisplayHeight)
	imageDataData := imageData.Get("data")

	return &HTMLCanvasDisplay{
		ram:           make([]int, 0x2000),
		canvas:        canvas,
		ctx2d:         ctx2d,
		imageData:     imageData,
		imageDataData: imageDataData,
		buffer:        make([]byte, imageDataData.Length()),
	}
}

// Reset set memory values to 0.
func (s *HTMLCanvasDisplay) Reset() error {
	for i := 0; i < len(s.ram); i++ {
		if err := s.WriteWord(i, 0); err != nil {
			return err
		}
	}

	return nil
}

func (s *HTMLCanvasDisplay) ReadWord(address int) (int, error) {
	if address < 0 || address > len(s.ram) {
		return 0, hackvm.BadAddressErr
	}

	word := s.ram[address]
	return word, nil
}

func (s *HTMLCanvasDisplay) WriteWord(address int, word int) error {
	if address < 0 || address > len(s.ram) {
		return hackvm.BadAddressErr
	}

	s.ram[address] = word
	return nil
}

// Refresh renders the screen data on the HTML canvas element.
func (s *HTMLCanvasDisplay) Refresh() error {
	for row := 0; row < DisplayHeight; row++ {
		for col := 0; col < DisplayWidth; col++ {
			addr := 32*row + col/16
			word, err := s.ReadWord(addr)
			if err != nil {
				return fmt.Errorf("refreshing screen: %w", err)
			}

			var (
				nth   = col % 16
				bit   = (word&(1<<nth))>>nth == 1
				red   = row*(DisplayWidth*4) + col*4
				green = red + 1
				blue  = red + 2
				alpha = red + 3
			)

			if bit {
				s.buffer[red] = 0x75
				s.buffer[green] = 0xf9
				s.buffer[blue] = 0x4c
			} else {
				s.buffer[red] = 0x00
				s.buffer[green] = 0x00
				s.buffer[blue] = 0x00
			}

			s.buffer[alpha] = 0xff
		}
	}

	js.CopyBytesToJS(s.imageDataData, s.buffer)
	s.ctx2d.Call("putImageData", s.imageData, 0, 0)
	return nil
}

// Canvas returns a HTML canvas element.
func (s *HTMLCanvasDisplay) Canvas() js.Value {
	return s.canvas
}
