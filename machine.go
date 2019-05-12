package main

import "io"

// Machine is the virtual brainfuck machine
type Machine struct {
	code []*Instruction // the program that’s executed by the machine
	ip   int            // Instruction pointer: It points to the instruction in the code that’s to be executed next

	memory [30000]int // a brainfuck machine has 30000 memory cells
	dp     int        // data pointer:  It “points” to a memory cell, by holding the value of the cell’s index

	// i/o stream with a buffer slice
	input   io.Reader
	output  io.Writer
	readBuf []byte
}

// NewMachine initializes a new machine
func NewMachine(instructions []*Instruction, in io.Reader, out io.Writer) *Machine {
	return &Machine{
		code:    instructions,
		input:   in,
		output:  out,
		readBuf: make([]byte, 1),
	}
}

// Execute starts the machine and and executes the intermediate representation
func (m *Machine) Execute() {
	for m.ip < len(m.code) {
		ins := m.code[m.ip]

		switch ins.Type {
		case Plus:
			m.memory[m.dp] += ins.Argument
		case Minus:
			m.memory[m.dp] -= ins.Argument
		case Right:
			m.dp += ins.Argument
		case Left:
			m.dp -= ins.Argument
		case ReadChar:
			for i := 0; i < ins.Argument; i++ {
				m.readChar()
			}
		case PutChar:
			for i := 0; i < ins.Argument; i++ {
				m.putChar()
			}
		case JumpIfZero:
			if m.memory[m.dp] == 0 {
				m.ip = ins.Argument
				continue
			}
		case JumpIfNotZero:
			if m.memory[m.dp] != 0 {
				m.ip = ins.Argument
				continue
			}
		}

		m.ip++
	}
}

// reads a character from input
func (m *Machine) readChar() {
	n, err := m.input.Read(m.readBuf)
	if err != nil {
		panic(err)
	}
	if n != 1 {
		panic("wrong num bytes read")
	}

	m.memory[m.dp] = int(m.readBuf[0])
}

// writes a character to output
func (m *Machine) putChar() {
	m.readBuf[0] = byte(m.memory[m.dp])

	n, err := m.output.Write(m.readBuf)
	if err != nil {
		panic(err)
	}
	if n != 1 {
		panic("wrong num bytes written")
	}
}
