// +build !js,!wasm

package main

type keyboard struct{}

func newKeyboard() *keyboard {
	return &keyboard{}
}

func (h *keyboard) ReadWord(int) (int, error) {
	return 0, nil
}
