// +build js,wasm

package main

import (
	"syscall/js"
)

// keyboard returns the current pressed key upon ReadWord call.
type keyboard struct {
	keyCode         int
	keyDownCallback js.Func
	keyUpCallback   js.Func
}

// newKeyboard creates a new Keyboard.
// It is required to call Setup to setup event handlers.
func newKeyboard() *keyboard {
	kbd := &keyboard{
		keyCode: 0,
	}

	kbd.keyDownCallback = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		e := args[0]
		switch e.Get("key").String() {
		case " ":
			kbd.keyCode = 32
		case "!":
			kbd.keyCode = 33
		case "\"":
			kbd.keyCode = 34
		case "#":
			kbd.keyCode = 35
		case "$":
			kbd.keyCode = 36
		case "%":
			kbd.keyCode = 37
		case "&":
			kbd.keyCode = 38
		case "'":
			kbd.keyCode = 39
		case "(":
			kbd.keyCode = 40
		case ")":
			kbd.keyCode = 41
		case "*":
			kbd.keyCode = 42
		case "+":
			kbd.keyCode = 43
		case ",":
			kbd.keyCode = 44
		case "-":
			kbd.keyCode = 45
		case ".":
			kbd.keyCode = 46
		case "0":
			kbd.keyCode = 48
		case "1":
			kbd.keyCode = 49
		case "2":
			kbd.keyCode = 50
		case "3":
			kbd.keyCode = 51
		case "4":
			kbd.keyCode = 52
		case "5":
			kbd.keyCode = 53
		case "6":
			kbd.keyCode = 54
		case "7":
			kbd.keyCode = 55
		case "8":
			kbd.keyCode = 56
		case "9":
			kbd.keyCode = 57
		case ":":
			kbd.keyCode = 58
		case ";":
			kbd.keyCode = 59
		case "<":
			kbd.keyCode = 60
		case "=":
			kbd.keyCode = 61
		case ">":
			kbd.keyCode = 62
		case "?":
			kbd.keyCode = 63
		case "@":
			kbd.keyCode = 64
		case "A":
			kbd.keyCode = 65
		case "B":
			kbd.keyCode = 66
		case "C":
			kbd.keyCode = 67
		case "D":
			kbd.keyCode = 68
		case "E":
			kbd.keyCode = 69
		case "F":
			kbd.keyCode = 70
		case "G":
			kbd.keyCode = 71
		case "H":
			kbd.keyCode = 72
		case "I":
			kbd.keyCode = 73
		case "J":
			kbd.keyCode = 74
		case "K":
			kbd.keyCode = 75
		case "L":
			kbd.keyCode = 76
		case "M":
			kbd.keyCode = 77
		case "N":
			kbd.keyCode = 78
		case "O":
			kbd.keyCode = 79
		case "P":
			kbd.keyCode = 80
		case "Q":
			kbd.keyCode = 81
		case "R":
			kbd.keyCode = 82
		case "S":
			kbd.keyCode = 83
		case "T":
			kbd.keyCode = 84
		case "U":
			kbd.keyCode = 85
		case "V":
			kbd.keyCode = 86
		case "W":
			kbd.keyCode = 87
		case "X":
			kbd.keyCode = 88
		case "Y":
			kbd.keyCode = 89
		case "Z":
			kbd.keyCode = 90
		case "[":
			kbd.keyCode = 91
		case "/":
			kbd.keyCode = 92
		case "]":
			kbd.keyCode = 93
		case "^":
			kbd.keyCode = 94
		case "_":
			kbd.keyCode = 95
		case "`":
			kbd.keyCode = 96
		case "a":
			kbd.keyCode = 97
		case "b":
			kbd.keyCode = 98
		case "c":
			kbd.keyCode = 99
		case "d":
			kbd.keyCode = 100
		case "e":
			kbd.keyCode = 101
		case "f":
			kbd.keyCode = 102
		case "g":
			kbd.keyCode = 103
		case "h":
			kbd.keyCode = 104
		case "i":
			kbd.keyCode = 105
		case "j":
			kbd.keyCode = 106
		case "k":
			kbd.keyCode = 107
		case "l":
			kbd.keyCode = 108
		case "m":
			kbd.keyCode = 109
		case "n":
			kbd.keyCode = 110
		case "o":
			kbd.keyCode = 111
		case "p":
			kbd.keyCode = 112
		case "q":
			kbd.keyCode = 113
		case "r":
			kbd.keyCode = 114
		case "s":
			kbd.keyCode = 115
		case "t":
			kbd.keyCode = 116
		case "u":
			kbd.keyCode = 117
		case "v":
			kbd.keyCode = 118
		case "w":
			kbd.keyCode = 119
		case "x":
			kbd.keyCode = 120
		case "y":
			kbd.keyCode = 121
		case "z":
			kbd.keyCode = 122
		case "{":
			kbd.keyCode = 123
		case "|":
			kbd.keyCode = 124
		case "}":
			kbd.keyCode = 125
		case "~":
			kbd.keyCode = 126
		case "Enter":
			kbd.keyCode = 138
		case "Backspace":
			kbd.keyCode = 129
		case "ArrowLeft":
			kbd.keyCode = 130
		case "ArrowUp":
			kbd.keyCode = 131
		case "ArrowRight":
			kbd.keyCode = 132
		case "ArrowDown":
			kbd.keyCode = 133
		case "Home":
			kbd.keyCode = 134
		case "End":
			kbd.keyCode = 135
		case "PageUp":
			kbd.keyCode = 136
		case "PageDown":
			kbd.keyCode = 137
		case "Insert":
			kbd.keyCode = 138
		case "Delete":
			kbd.keyCode = 139
		case "Escape":
			kbd.keyCode = 140
		case "F1":
			kbd.keyCode = 141
		case "F2":
			kbd.keyCode = 142
		case "F3":
			kbd.keyCode = 143
		case "F4":
			kbd.keyCode = 144
		case "F5":
			kbd.keyCode = 145
		case "F6":
			kbd.keyCode = 146
		case "F7":
			kbd.keyCode = 147
		case "F8":
			kbd.keyCode = 148
		case "F9":
			kbd.keyCode = 149
		case "F10":
			kbd.keyCode = 150
		case "F11":
			kbd.keyCode = 121
		case "F12":
			kbd.keyCode = 152
		}
		return nil
	})

	kbd.keyUpCallback = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		kbd.keyCode = 0
		return nil
	})

	return kbd
}

// Setup configure keydown and keyup event listeners.
func (k *keyboard) Setup() {
	js.Global().Get("window").Call("addEventListener", "keydown", k.keyDownCallback)
	js.Global().Get("window").Call("addEventListener", "keyup", k.keyUpCallback)
}

// Dispose remove event listeners.
func (k *keyboard) Dispose() {
	js.Global().Get("window").Call("removeEventListener", "keydown", k.keyDownCallback)
	js.Global().Get("window").Call("removeEventListener", "keyup", k.keyUpCallback)
}

// ReadWord return the current pressed key.
func (k *keyboard) ReadWord(int) (int, error) {
	return k.keyCode, nil
}
