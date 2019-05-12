package main

// InsType represents New Instruction type
type InsType byte

// intermediate instruction set
const (
	Plus          InsType = '+'
	Minus         InsType = '-'
	Right         InsType = '>'
	Left          InsType = '<'
	PutChar       InsType = '.'
	ReadChar      InsType = ','
	JumpIfZero    InsType = '['
	JumpIfNotZero InsType = ']'
)

// Instruction represent a single instruction of the intermediate instruction set
type Instruction struct {
	Type     InsType
	Argument int
}
