package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	flag.Parse()
	args := flag.Args()

	if len(args) != 1 {
		fmt.Fprintf(os.Stderr, "no filename given\n")
		os.Exit(-1)
	}

	fileName := os.Args[1]
	code, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
		os.Exit(-1)
	}

	compiler := NewCompiler(string(code))
	instructions := compiler.Compile()

	m := NewMachine(instructions, os.Stdin, os.Stdout)
	m.Execute()
}
