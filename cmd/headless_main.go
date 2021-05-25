// +build !js,!wasm

package main

import (
	"hackvm"
	"hackvm/games"
	"log"
	"os"
	"os/signal"
	"time"
)

func main() {
	rom := hackvm.NewROM32K()
	if err := rom.Load(games.Pong); err != nil {
		panic(err)
	}

	screen := newDisplay()
	keyboard := newKeyboard()

	ram, err := hackvm.NewRAM16K(screen, keyboard)
	if err != nil {
		panic(err)
	}

	cpu, err := hackvm.NewCPU(ram, rom)
	if err != nil {
		panic(err)
	}

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, os.Kill)

	go func() {
		for range sig {
			log.Println("Stopping VM")
			os.Exit(0)
		}
	}()

	log.Println("Starting VM in headless mode")
	for {
		for i := 0; i < 50000; i++ {
			if err := cpu.Step(); err != nil {
				panic(err)
			}
		}
		time.Sleep((1000 / 30) * time.Millisecond)
	}
}
