package hackvm

import (
	"errors"
	"fmt"
)

// ErrBadInstruction error returned while executing invalid instructions.
var ErrBadInstruction = errors.New("invalid instruction")

// CPU implements the HackVM instruction set.
type CPU struct {
	ram ReadWriteMemory
	rom ReadOnlyMemory
	a   int
	d   int
	pc  int
}

// NewCPU returns a new instance of a CPU.
func NewCPU(ram ReadWriteMemory, rom ReadOnlyMemory) (*CPU, error) {
	if ram == nil {
		return nil, errors.New("ram is nil")
	}
	if rom == nil {
		return nil, errors.New("rom is nil")
	}

	return &CPU{
		ram: ram,
		rom: rom,
	}, nil
}

// Load loads a program into memory.
func (c *CPU) Load(program ReadOnlyMemory) error {
	if program == nil {
		return errors.New("program is nil")
	}

	c.rom = program
	c.Reset()

	return nil
}

// A returns the value of the a register.
func (c *CPU) A() int {
	return c.a
}

// D returns the value of the d register.
func (c *CPU) D() int {
	return c.d
}

// PC returns the value of the pc register.
func (c *CPU) PC() int {
	return c.pc
}

// Reset clear the value of the pc register.
func (c *CPU) Reset() {
	c.pc = 0
}

// Step loads and execute a single instruction.
func (c *CPU) Step() error {
	instruction, err := c.rom.ReadWord(c.pc)
	if err != nil {
		return fmt.Errorf("reading instruction: %w", err)
	}

	switch (instruction & 0x8000) >> 15 {
	case 0:
		c.a = instruction & 0x7FFF
		c.pc += 1
	case 1:
		var (
			compM  = (instruction&0x1000)>>12 == 1
			comp   = instruction & 0xFC0
			dest   = instruction & 0x038 >> 3
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
			// 0
			aluOut = 0
		case 0xFC0:
			// 1
			aluOut = 1
		case 0xE80:
			// -1
			aluOut = -1
		case 0x300:
			// D
			aluOut = c.d
		case 0xC00:
			// A
			// M
			if compM {
				word, err = c.ram.ReadWord(c.a)
				aluOut = word
			} else {
				aluOut = c.a
			}
			break
		case 0x340:
			// !D
			aluOut = ^c.d
			break
		case 0xC40:
			// !A
			// !M
			if compM {
				word, err = c.ram.ReadWord(c.a)
				aluOut = ^word
			} else {
				aluOut = ^c.a
			}
		case 0x3C0:
			// -D
			aluOut = -c.d
		case 0xCC0:
			// -A
			aluOut = -c.a
		case 0x7C0:
			// D+1
			aluOut = c.d + 1
		case 0xDC0:
			// A+1
			// M+1
			if compM {
				word, err = c.ram.ReadWord(c.a)
				aluOut = word + 1
			} else {
				aluOut = c.a + 1
			}
		case 0x380:
			// D-1
			aluOut = c.d - 1
		case 0xC80:
			// A-1
			// M-1
			if compM {
				word, err = c.ram.ReadWord(c.a)
				aluOut = word - 1
			} else {
				aluOut = c.a - 1
			}
		case 0x080:
			// D+A
			// D+M
			if compM {
				word, err = c.ram.ReadWord(c.a)
				aluOut = c.d + word
			} else {
				aluOut = c.d + c.a
			}
		case 0x4C0:
			// D-A
			// D-M
			if compM {
				word, err = c.ram.ReadWord(c.a)
				aluOut = c.d - word
			} else {
				aluOut = c.d - c.a
			}
		case 0x1C0:
			// A-D
			// M-D
			if compM {
				word, err = c.ram.ReadWord(c.a)
				aluOut = word - c.d
			} else {
				aluOut = c.a - c.d
			}
		case 0x000:
			// D&A
			// D&M
			if compM {
				word, err = c.ram.ReadWord(c.a)
				aluOut = c.d & word
			} else {
				aluOut = c.d & c.a
			}
		case 0x540:
			// D|A
			// D|M
			if compM {
				word, err = c.ram.ReadWord(c.a)
				aluOut = c.d | word
			} else {
				aluOut = c.d | c.a
			}
		default:
			return ErrBadInstruction
		}

		if err != nil {
			return fmt.Errorf("executing instruction 0x%X: %w", instruction, err)
		}

		if destM {
			if err := c.ram.WriteWord(c.a, aluOut); err != nil {
				return fmt.Errorf("executing instruction 0x%X: %w", instruction, err)
			}
		}

		if destA {
			c.a = aluOut
		}

		if destD {
			c.d = aluOut
		}

		jump := false
		switch jmp {
		case 0x00:
			jump = false
		case 0x01:
			// JGT
			if aluOut > 0 {
				jump = true
			}
		case 0x02:
			// JEQ
			if aluOut == 0 {
				jump = true
			}
		case 0x03:
			// JGE
			if aluOut >= 0 {
				jump = true
			}
		case 0x04:
			// JLT
			if aluOut < 0 {
				jump = true
			}
		case 0x05:
			// JNE
			if aluOut != 0 {
				jump = true
			}
		case 0x06:
			// JLE
			if aluOut <= 0 {
				jump = true
			}
		case 0x07:
			// JMP
			jump = true
		}

		if jump {
			c.pc = c.a
		} else {
			c.pc += 1
		}
	}

	return nil
}
