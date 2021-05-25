// +build js,wasm

package main

import (
	"syscall/js"

	"hackvm"
	"hackvm/games"
)

func main() {
	rom := hackvm.NewROM32K()
	if err := rom.Load(games.Pong); err != nil {
		panic(err)
	}

	screen := newDisplay()
	keyboard := newKeyboard()
	keyboard.Setup()

	ram, err := hackvm.NewRAM16K(screen, keyboard)
	if err != nil {
		panic(err)
	}

	cpu, err := hackvm.NewCPU(ram, rom)
	if err != nil {
		panic(err)
	}

	canvas := screen.Canvas()
	js.Global().Get("document").Get("body").Call("append", canvas)

	var (
		lastTimestamp float64
		renderFrame   js.Func
	)

	renderFrame = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		timestamp := args[0].Float()

		if timestamp-lastTimestamp >= (1000 / 30) {
			for i := 0; i < 50000; i++ {
				if err := cpu.Step(); err != nil {
					panic(err)
				}
			}

			if err := screen.Refresh(); err != nil {
				panic(err)
			}

			lastTimestamp = timestamp
		}

		js.Global().Get("window").Call("requestAnimationFrame", renderFrame)
		return nil
	})

	js.Global().Get("window").Call("requestAnimationFrame", renderFrame)

	<-make(chan struct{})
	keyboard.Dispose()
}
