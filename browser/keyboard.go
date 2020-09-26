package browser

import (
	"syscall/js"
)

// Keyboard returns the current pressed key upon ReadWord call.
type Keyboard struct {
	keyCode         int
	keyDownCallback js.Func
	keyUpCallback   js.Func
}

// NewKeyboard creates a new Keyboard.
// It is required to call Setup to setup event handlers.
func NewKeyboard() *Keyboard {
	keyboard := &Keyboard{
		keyCode: 0,
	}

	keyboard.keyDownCallback = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		e := args[0]
		switch e.Get("key").String() {
		case " ":
			keyboard.keyCode = 32
		case "!":
			keyboard.keyCode = 33
		case "\"":
			keyboard.keyCode = 34
		case "#":
			keyboard.keyCode = 35
		case "$":
			keyboard.keyCode = 36
		case "%":
			keyboard.keyCode = 37
		case "&":
			keyboard.keyCode = 38
		case "'":
			keyboard.keyCode = 39
		case "(":
			keyboard.keyCode = 40
		case ")":
			keyboard.keyCode = 41
		case "*":
			keyboard.keyCode = 42
		case "+":
			keyboard.keyCode = 43
		case ",":
			keyboard.keyCode = 44
		case "-":
			keyboard.keyCode = 45
		case ".":
			keyboard.keyCode = 46
		case "0":
			keyboard.keyCode = 48
		case "1":
			keyboard.keyCode = 49
		case "2":
			keyboard.keyCode = 50
		case "3":
			keyboard.keyCode = 51
		case "4":
			keyboard.keyCode = 52
		case "5":
			keyboard.keyCode = 53
		case "6":
			keyboard.keyCode = 54
		case "7":
			keyboard.keyCode = 55
		case "8":
			keyboard.keyCode = 56
		case "9":
			keyboard.keyCode = 57
		case ":":
			keyboard.keyCode = 58
		case ";":
			keyboard.keyCode = 59
		case "<":
			keyboard.keyCode = 60
		case "=":
			keyboard.keyCode = 61
		case ">":
			keyboard.keyCode = 62
		case "?":
			keyboard.keyCode = 63
		case "@":
			keyboard.keyCode = 64
		case "A":
			keyboard.keyCode = 65
		case "B":
			keyboard.keyCode = 66
		case "C":
			keyboard.keyCode = 67
		case "D":
			keyboard.keyCode = 68
		case "E":
			keyboard.keyCode = 69
		case "F":
			keyboard.keyCode = 70
		case "G":
			keyboard.keyCode = 71
		case "H":
			keyboard.keyCode = 72
		case "I":
			keyboard.keyCode = 73
		case "J":
			keyboard.keyCode = 74
		case "K":
			keyboard.keyCode = 75
		case "L":
			keyboard.keyCode = 76
		case "M":
			keyboard.keyCode = 77
		case "N":
			keyboard.keyCode = 78
		case "O":
			keyboard.keyCode = 79
		case "P":
			keyboard.keyCode = 80
		case "Q":
			keyboard.keyCode = 81
		case "R":
			keyboard.keyCode = 82
		case "S":
			keyboard.keyCode = 83
		case "T":
			keyboard.keyCode = 84
		case "U":
			keyboard.keyCode = 85
		case "V":
			keyboard.keyCode = 86
		case "W":
			keyboard.keyCode = 87
		case "X":
			keyboard.keyCode = 88
		case "Y":
			keyboard.keyCode = 89
		case "Z":
			keyboard.keyCode = 90
		case "[":
			keyboard.keyCode = 91
		case "/":
			keyboard.keyCode = 92
		case "]":
			keyboard.keyCode = 93
		case "^":
			keyboard.keyCode = 94
		case "_":
			keyboard.keyCode = 95
		case "`":
			keyboard.keyCode = 96
		case "a":
			keyboard.keyCode = 97
		case "b":
			keyboard.keyCode = 98
		case "c":
			keyboard.keyCode = 99
		case "d":
			keyboard.keyCode = 100
		case "e":
			keyboard.keyCode = 101
		case "f":
			keyboard.keyCode = 102
		case "g":
			keyboard.keyCode = 103
		case "h":
			keyboard.keyCode = 104
		case "i":
			keyboard.keyCode = 105
		case "j":
			keyboard.keyCode = 106
		case "k":
			keyboard.keyCode = 107
		case "l":
			keyboard.keyCode = 108
		case "m":
			keyboard.keyCode = 109
		case "n":
			keyboard.keyCode = 110
		case "o":
			keyboard.keyCode = 111
		case "p":
			keyboard.keyCode = 112
		case "q":
			keyboard.keyCode = 113
		case "r":
			keyboard.keyCode = 114
		case "s":
			keyboard.keyCode = 115
		case "t":
			keyboard.keyCode = 116
		case "u":
			keyboard.keyCode = 117
		case "v":
			keyboard.keyCode = 118
		case "w":
			keyboard.keyCode = 119
		case "x":
			keyboard.keyCode = 120
		case "y":
			keyboard.keyCode = 121
		case "z":
			keyboard.keyCode = 122
		case "{":
			keyboard.keyCode = 123
		case "|":
			keyboard.keyCode = 124
		case "}":
			keyboard.keyCode = 125
		case "~":
			keyboard.keyCode = 126
		case "Enter":
			keyboard.keyCode = 138
		case "Backspace":
			keyboard.keyCode = 129
		case "ArrowLeft":
			keyboard.keyCode = 130
		case "ArrowUp":
			keyboard.keyCode = 131
		case "ArrowRight":
			keyboard.keyCode = 132
		case "ArrowDown":
			keyboard.keyCode = 133
		case "Home":
			keyboard.keyCode = 134
		case "End":
			keyboard.keyCode = 135
		case "PageUp":
			keyboard.keyCode = 136
		case "PageDown":
			keyboard.keyCode = 137
		case "Insert":
			keyboard.keyCode = 138
		case "Delete":
			keyboard.keyCode = 139
		case "Escape":
			keyboard.keyCode = 140
		case "F1":
			keyboard.keyCode = 141
		case "F2":
			keyboard.keyCode = 142
		case "F3":
			keyboard.keyCode = 143
		case "F4":
			keyboard.keyCode = 144
		case "F5":
			keyboard.keyCode = 145
		case "F6":
			keyboard.keyCode = 146
		case "F7":
			keyboard.keyCode = 147
		case "F8":
			keyboard.keyCode = 148
		case "F9":
			keyboard.keyCode = 149
		case "F10":
			keyboard.keyCode = 150
		case "F11":
			keyboard.keyCode = 121
		case "F12":
			keyboard.keyCode = 152
		}
		return nil
	})

	keyboard.keyUpCallback = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		keyboard.keyCode = 0
		return nil
	})

	return keyboard
}

// Setup configure keydown and keyup event listeners.
func (k *Keyboard) Setup() {
	js.Global().Get("window").Call("addEventListener", "keydown", k.keyDownCallback)
	js.Global().Get("window").Call("addEventListener", "keyup", k.keyUpCallback)
}

// Dispose remove event listeners.
func (k *Keyboard) Dispose() {
	js.Global().Get("window").Call("removeEventListener", "keydown", k.keyDownCallback)
	js.Global().Get("window").Call("removeEventListener", "keyup", k.keyUpCallback)
}

// ReadWord return the current pressed key.
func (k *Keyboard) ReadWord(int) (int, error) {
	return k.keyCode, nil
}
