package hackvm

import (
	"bufio"
	"errors"
	"fmt"
	"image"
	"os"
	"strconv"
)

const (
	DisplayWidth  = 512
	DisplayHeight = 256
)

// ErrBadInstruction error returned while executing invalid instructions.
var ErrBadInstruction = errors.New("invalid instruction")

// ErrBadAddress error returned while accessing an invalid memory address.
var ErrBadAddress = errors.New("invalid address")

type Computer struct {
	rom    []int
	ram    []int
	ra     int
	rd     int
	pc     int
	screen *image.RGBA
}

func (c *Computer) ReadRam(address int) (int, error) {
	if address > len(c.ram) {
		return 0, ErrBadAddress
	}
	return c.ram[address], nil
}

func (c *Computer) readScreen(address int) (int, error) {
	if address > 0x6000 {
		return 0, ErrBadAddress
	}
	return c.ram[0x4000+address], nil
}

func (c *Computer) WriteRam(address, data int) error {
	if address > len(c.ram) {
		return ErrBadAddress
	}
	c.ram[address] = data
	return nil
}

func (c *Computer) writeRom(address, data int) error {
	if address > len(c.rom) {
		return ErrBadAddress
	}
	c.rom[address] = data
	return nil
}

func (c *Computer) ReadRamCompM(compM bool) (int, error) {
	if compM {
		return c.ReadRam(c.ra)
	}
	return c.ra, nil
}

func (c *Computer) ReadRom(address int) (int, error) {
	if address > len(c.rom) {
		return 0, ErrBadAddress
	}
	return c.rom[address], nil
}

func (c *Computer) Step() error {
	instruction, err := c.ReadRom(c.pc)
	if err != nil {
		return fmt.Errorf("reading instruction: %w", err)
	}

	switch (instruction & 0x8000) >> 15 {
	case 0:
		c.ra = instruction & 0x7FFF
		c.pc += 1
	case 1:
		var (
			compM  = (instruction&0x1000)>>12 == 1
			comp   = instruction & 0xFC0
			dest   = (instruction & 0x038) >> 3
			destA  = (dest&0x4)>>2 == 1
			destD  = (dest&0x2)>>1 == 1
			destM  = dest&0x1 == 1
			jmp    = instruction & 0x7
			aluOut int
			word   int
			err    error
		)

		switch comp {
		case 0xA80:
			aluOut = 0
		case 0xFC0:
			aluOut = 1
		case 0xE80:
			aluOut = -1
		case 0x300:
			aluOut = c.rd
		case 0xC00:
			aluOut, err = c.ReadRamCompM(compM)
		case 0x340:
			aluOut = ^c.rd
		case 0xC40:
			word, err = c.ReadRamCompM(compM)
			aluOut = ^word
		case 0x3C0:
			aluOut = -c.rd
		case 0xCC0:
			aluOut = -c.ra
		case 0x7C0:
			aluOut = c.rd + 1
		case 0xDC0:
			word, err = c.ReadRamCompM(compM)
			aluOut = word + 1
		case 0x380:
			aluOut = c.rd - 1
		case 0xC80:
			word, err = c.ReadRamCompM(compM)
			aluOut = word - 1
		case 0x080:
			word, err = c.ReadRamCompM(compM)
			aluOut = c.rd + word
		case 0x4C0:
			word, err = c.ReadRamCompM(compM)
			aluOut = c.rd - word
		case 0x1C0:
			word, err = c.ReadRamCompM(compM)
			aluOut = word - c.rd
		case 0x000:
			word, err = c.ReadRamCompM(compM)
			aluOut = c.rd & word
		case 0x540:
			word, err = c.ReadRamCompM(compM)
			aluOut = c.rd | word
		default:
			return ErrBadInstruction
		}

		if err != nil {
			return fmt.Errorf("executing instruction 0x%X: %w", instruction, err)
		}

		if destM {
			if err := c.WriteRam(c.ra, aluOut); err != nil {
				return fmt.Errorf("executing instruction 0x%X: %w", instruction, err)
			}
		}

		if destA {
			c.ra = aluOut
		}

		if destD {
			c.rd = aluOut
		}

		jump := false
		switch jmp {
		case 0x01:
			if aluOut > 0 {
				jump = true
			}
		case 0x02:
			if aluOut == 0 {
				jump = true
			}
		case 0x03:
			if aluOut >= 0 {
				jump = true
			}
		case 0x04:
			if aluOut < 0 {
				jump = true
			}
		case 0x05:
			if aluOut != 0 {
				jump = true
			}
		case 0x06:
			if aluOut <= 0 {
				jump = true
			}
		case 0x07:
			jump = true
		default:
			jump = false
		}

		if jump {
			c.pc = c.ra
		} else {
			c.pc += 1
		}
	}
	return nil
}

func (c *Computer) Reset() {
	c.pc = 0
}

func (c *Computer) KeyUp() {
	if err := c.WriteRam(0x6000, 0); err != nil {
		panic(fmt.Sprintf("error is not expected: %s", err.Error()))
	}
}

func (c *Computer) KeyDown(keyChar string) {
	var keyCode = 0
	switch keyChar {
	case " ":
		keyCode = 32
	case "!":
		keyCode = 33
	case "\"":
		keyCode = 34
	case "#":
		keyCode = 35
	case "$":
		keyCode = 36
	case "%":
		keyCode = 37
	case "&":
		keyCode = 38
	case "'":
		keyCode = 39
	case "(":
		keyCode = 40
	case ")":
		keyCode = 41
	case "*":
		keyCode = 42
	case "+":
		keyCode = 43
	case ",":
		keyCode = 44
	case "-":
		keyCode = 45
	case ".":
		keyCode = 46
	case "0":
		keyCode = 48
	case "1":
		keyCode = 49
	case "2":
		keyCode = 50
	case "3":
		keyCode = 51
	case "4":
		keyCode = 52
	case "5":
		keyCode = 53
	case "6":
		keyCode = 54
	case "7":
		keyCode = 55
	case "8":
		keyCode = 56
	case "9":
		keyCode = 57
	case ":":
		keyCode = 58
	case ";":
		keyCode = 59
	case "<":
		keyCode = 60
	case "=":
		keyCode = 61
	case ">":
		keyCode = 62
	case "?":
		keyCode = 63
	case "@":
		keyCode = 64
	case "A":
		keyCode = 65
	case "B":
		keyCode = 66
	case "C":
		keyCode = 67
	case "D":
		keyCode = 68
	case "E":
		keyCode = 69
	case "F":
		keyCode = 70
	case "G":
		keyCode = 71
	case "H":
		keyCode = 72
	case "I":
		keyCode = 73
	case "J":
		keyCode = 74
	case "K":
		keyCode = 75
	case "L":
		keyCode = 76
	case "M":
		keyCode = 77
	case "N":
		keyCode = 78
	case "O":
		keyCode = 79
	case "P":
		keyCode = 80
	case "Q":
		keyCode = 81
	case "R":
		keyCode = 82
	case "S":
		keyCode = 83
	case "T":
		keyCode = 84
	case "U":
		keyCode = 85
	case "V":
		keyCode = 86
	case "W":
		keyCode = 87
	case "X":
		keyCode = 88
	case "Y":
		keyCode = 89
	case "Z":
		keyCode = 90
	case "[":
		keyCode = 91
	case "/":
		keyCode = 92
	case "]":
		keyCode = 93
	case "^":
		keyCode = 94
	case "_":
		keyCode = 95
	case "`":
		keyCode = 96
	case "a":
		keyCode = 97
	case "b":
		keyCode = 98
	case "c":
		keyCode = 99
	case "d":
		keyCode = 100
	case "e":
		keyCode = 101
	case "f":
		keyCode = 102
	case "g":
		keyCode = 103
	case "h":
		keyCode = 104
	case "i":
		keyCode = 105
	case "j":
		keyCode = 106
	case "k":
		keyCode = 107
	case "l":
		keyCode = 108
	case "m":
		keyCode = 109
	case "n":
		keyCode = 110
	case "o":
		keyCode = 111
	case "p":
		keyCode = 112
	case "q":
		keyCode = 113
	case "r":
		keyCode = 114
	case "s":
		keyCode = 115
	case "t":
		keyCode = 116
	case "u":
		keyCode = 117
	case "v":
		keyCode = 118
	case "w":
		keyCode = 119
	case "x":
		keyCode = 120
	case "y":
		keyCode = 121
	case "z":
		keyCode = 122
	case "{":
		keyCode = 123
	case "|":
		keyCode = 124
	case "}":
		keyCode = 125
	case "~":
		keyCode = 126
	case "Enter":
		keyCode = 138
	case "Backspace":
		keyCode = 129
	case "ArrowLeft":
		keyCode = 130
	case "ArrowUp":
		keyCode = 131
	case "ArrowRight":
		keyCode = 132
	case "ArrowDown":
		keyCode = 133
	case "Home":
		keyCode = 134
	case "End":
		keyCode = 135
	case "PageUp":
		keyCode = 136
	case "PageDown":
		keyCode = 137
	case "Insert":
		keyCode = 138
	case "Delete":
		keyCode = 139
	case "Escape":
		keyCode = 140
	case "F1":
		keyCode = 141
	case "F2":
		keyCode = 142
	case "F3":
		keyCode = 143
	case "F4":
		keyCode = 144
	case "F5":
		keyCode = 145
	case "F6":
		keyCode = 146
	case "F7":
		keyCode = 147
	case "F8":
		keyCode = 148
	case "F9":
		keyCode = 149
	case "F10":
		keyCode = 150
	case "F11":
		keyCode = 121
	case "F12":
		keyCode = 152
	}

	if err := c.WriteRam(0x6000, keyCode); err != nil {
		panic(fmt.Sprintf("error is not expected: %s", err.Error()))
	}
}

func (c *Computer) LoadRomFromFile(path string) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer func() {
		_ = f.Close()
	}()

	var (
		scanner = bufio.NewScanner(f)
		address = 0
	)

	for scanner.Scan() {
		data, err := strconv.ParseInt(scanner.Text(), 2, strconv.IntSize)
		if err != nil {
			return fmt.Errorf("%s is not a .hack file", path)
		}
		if err := c.writeRom(address, int(data)); err != nil {
			return fmt.Errorf("unable to write into memory: %w", err)
		}
		address += 1
	}

	return nil
}

func (c *Computer) ScreenBuffer() (image.Image, error) {
	for row := 0; row < DisplayHeight; row++ {
		for col := 0; col < DisplayWidth; col++ {
			addr := 32*row + col/16
			word, err := c.readScreen(addr)
			if err != nil {
				return nil, fmt.Errorf("refreshing screen: %w", err)
			}

			var (
				nth = col % 16
				bit = (word&(1<<nth))>>nth == 1
				r   = row*(DisplayWidth*4) + col*4
				g   = r + 1
				b   = r + 2
				a   = r + 3
			)

			if bit {
				c.screen.Pix[r] = 0x75
				c.screen.Pix[g] = 0xf9
				c.screen.Pix[b] = 0x4c
			} else {
				c.screen.Pix[r] = 0x00
				c.screen.Pix[g] = 0x00
				c.screen.Pix[b] = 0x00
			}

			c.screen.Pix[a] = 0xFF
		}
	}

	return c.screen, nil
}

func NewComputer() *Computer {
	return &Computer{
		ram:    make([]int, 0x6001),
		rom:    make([]int, 0x8000),
		screen: image.NewRGBA(image.Rect(0, 0, DisplayWidth, DisplayHeight)),
	}
}
