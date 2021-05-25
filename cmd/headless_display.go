// +build !js,!wasm

package main

type display struct{}

func newDisplay() *display {
	return &display{}
}

func (d *display) ReadWord(int) (int, error) {
	return 0, nil
}

func (d *display) WriteWord(int, int) error {
	return nil
}
