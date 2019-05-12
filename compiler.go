package main

// Compiler compiler brainfuck code to our intermediate representation
type Compiler struct {
	code       string
	codeLength int
	position   int

	instructions []*Instruction
}

// NewCompiler creates a new compiler
func NewCompiler(code string) *Compiler {
	return &Compiler{
		code:         code,             // original code
		codeLength:   len(code),        //length of code
		instructions: []*Instruction{}, // our representation
	}
}

// Compile compiles brainfucks to our instruction set
func (c *Compiler) Compile() []*Instruction {
	loopStack := []int{} // stack to keep track of '['s

	for c.position < c.codeLength {
		current := c.code[c.position]

		switch current {
		case '[':
			insPos := c.EmitWithArg(JumpIfZero, 0)
			loopStack = append(loopStack, insPos)
		case ']':
			// Pop position of last JumpIfZero ("[") instruction off stack
			openInstruction := loopStack[len(loopStack)-1]
			loopStack = loopStack[:len(loopStack)-1]

			// Emit the new JumpIfNotZero ("]") instruction,
			// with correct position as argument
			closeInstructionPos := c.EmitWithArg(JumpIfNotZero, openInstruction)

			// Patch the old JumpIfZero ("[") instruction with new position
			c.instructions[openInstruction].Argument = closeInstructionPos
		case '+':
			c.CompileFoldableInstruction('+', Plus)
		case '-':
			c.CompileFoldableInstruction('-', Minus)
		case '<':
			c.CompileFoldableInstruction('<', Left)
		case '>':
			c.CompileFoldableInstruction('>', Right)
		case '.':
			c.CompileFoldableInstruction('.', PutChar)
		case ',':
			c.CompileFoldableInstruction(',', ReadChar)
		}

		c.position++
	}

	return c.instructions
}

// CompileFoldableInstruction scans the input for repeated codes that can be folded to one instruction
func (c *Compiler) CompileFoldableInstruction(char byte, insType InsType) {
	count := 1

	for c.position < c.codeLength-1 && c.code[c.position+1] == char {
		count++
		c.position++
	}

	c.EmitWithArg(insType, count)
}

// EmitWithArg creates a new instruction object
func (c *Compiler) EmitWithArg(insType InsType, arg int) int {
	ins := &Instruction{Type: insType, Argument: arg}
	c.instructions = append(c.instructions, ins)
	return len(c.instructions) - 1
}
