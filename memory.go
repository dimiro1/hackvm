package hackvm

import "errors"

// ErrBadAddress error returned while accessing an invalid memory address.
var ErrBadAddress = errors.New("invalid address")

// WriteOnlyMemory interface for write only memories.
type WriteOnlyMemory interface {
	// WriteWord writes `word` into `address`.
	// return ErrBadAddress if an invalid address is given.
	WriteWord(address int, word int) error
}

// ReadOnlyMemory interface for read only memories.
type ReadOnlyMemory interface {
	// ReadWord reads the contents of `address`.
	// return ErrBadAddress if an invalid address is given.
	ReadWord(address int) (int, error)
}

// ReadWriteMemory compound interface for read and write memories.
type ReadWriteMemory interface {
	ReadOnlyMemory
	WriteOnlyMemory
}

// ROM32K program ROM.
type ROM32K struct {
	rom []int
}

func (r *ROM32K) ReadWord(address int) (int, error) {
	if address < 0 || int(address) > len(r.rom) {
		return 0, ErrBadAddress
	}

	return r.rom[address], nil
}

// NewROM32K returns a new ROM32K
func NewROM32K() *ROM32K {
	return &ROM32K{
		rom: make([]int, 0x8000),
	}
}

// Load loads the program into ROM memory.
func (r *ROM32K) Load(program []int) error {
	r.rom = program
	return nil
}

// RAM16K abstract the main RAM, Screen and keyboard into a single addressable memory.
type RAM16K struct {
	ram      []int
	display  ReadWriteMemory
	keyboard ReadOnlyMemory
}

// Reset sets ram to zero value.
func (r *RAM16K) Reset() {
	for i := 0; i < len(r.ram); i++ {
		r.ram[i] = 0
	}
}

func (r *RAM16K) ReadWord(address int) (int, error) {
	if address == 0x6000 {
		return r.keyboard.ReadWord(address)
	}

	if address > 0x6000 || address < 0 {
		return 0, ErrBadAddress
	}

	if address <= 0x03FFF {
		return r.ram[address], nil
	} else {
		return r.display.ReadWord(address - 0x4000)
	}
}

func (r *RAM16K) WriteWord(address int, word int) error {
	if address > 0x6000 || address < 0 {
		return ErrBadAddress
	}

	if address <= 0x03FFF {
		r.ram[address] = word
	} else {
		return r.display.WriteWord(address-0x4000, word)
	}

	return nil
}

// NewRAM16K returns a new RAM16K.
// Return errors if the parameters are invalid.
func NewRAM16K(display ReadWriteMemory, keyboard ReadOnlyMemory) (*RAM16K, error) {
	if display == nil {
		return nil, errors.New("display is nil")
	}
	if keyboard == nil {
		return nil, errors.New("keyboard is nil")
	}

	return &RAM16K{
		ram:      make([]int, 0x4000),
		display:  display,
		keyboard: keyboard,
	}, nil
}
