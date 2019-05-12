package main

import "io"

// Machine is the virtual brainfuck machine
type Machine struct {
	code string // the program that’s executed by the machine
	ip   int    // Instruction pointer: It points to the instruction in the code that’s to be executed next

	memory [30000]int // a brainfuck machine has 30000 memory cells
	dp     int        // data pointer:  It “points” to a memory cell, by holding the value of the cell’s index

	// i/o stream with a buffer slice
	input  io.Reader
	output io.Writer
	buf    []byte
}

// NewMachine initializes a new machine
func NewMachine(code string, in io.Reader, out io.Writer) *Machine {
	return &Machine{
		code:   code,
		input:  in,
		output: out,
		buf:    make([]byte, 1),
	}
}

// Execute starts the machine and and executes code
func (m *Machine) Execute() {
	for m.ip < len(m.code) {
		ins := m.code[m.ip]

		switch ins {
		case '+':
			m.memory[m.dp]++
		case '-':
			m.memory[m.dp]--
		case '>':
			m.dp++
		case '<':
			m.dp--
		case ',':
			m.readChar()
		case '.':
			m.putChar()
		case '[':
			if m.memory[m.dp] == 0 {
				depth := 1
				for depth != 0 {
					m.ip++
					switch m.code[m.ip] {
					case '[':
						depth++
					case ']':
						depth--
					}
				}
			}
		case ']':
			if m.memory[m.dp] != 0 {
				depth := 1
				for depth != 0 {
					m.ip--
					switch m.code[m.ip] {
					case ']':
						depth++
					case '[':
						depth--
					}
				}
			}
		}

		m.ip++
	}
}

// reads a character from input
func (m *Machine) readChar() {
	n, err := m.input.Read(m.buf)
	if err != nil {
		panic(err)
	}
	if n != 1 {
		panic("wrong num bytes read")
	}

	m.memory[m.dp] = int(m.buf[0])
}

// writes a character to output
func (m *Machine) putChar() {
	m.buf[0] = byte(m.memory[m.dp])

	n, err := m.output.Write(m.buf)
	if err != nil {
		panic(err)
	}
	if n != 1 {
		panic("wrong num bytes written")
	}
}
