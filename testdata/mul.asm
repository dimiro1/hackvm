// This file is part of www.nand2tetris.org
// and the book "The Elements of Computing Systems"
// by Nisan and Schocken, MIT Press.
// File name: projects/04/Mult.asm

// Multiplies R0 and R1 and stores the result in R2.
// (R0, R1, R2 refer to RAM[0], RAM[1], and RAM[2], respectively.)

@R2
M=0 // @R2 = 0

@R1
D=M
@i
M=D // @i = @R1

(LOOP)
	@i
	D=M
	@HALT
	D;JLE // if (@i <= 0); goto @HALT

	@R0
	D=M
	@R2
	D=D+M
	M=D  // @R2 = @R2 + R0

	@i
	D=M
	D=D-1
	@i
	M=D   // @i = @i - 1

	@LOOP
	0;JMP

(HALT)
	@HALT
	0;JMP