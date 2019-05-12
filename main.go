package main

import (
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	flag.Parse()
	args := flag.Args()

	if len(args) != 1 {
		handleError(errors.New("no filename given"))
	}

	fileName := os.Args[1]
	code, err := ioutil.ReadFile(fileName)
	if err != nil {
		handleError(err)
	}

	compiler := NewCompiler(string(code))
	instructions := compiler.Compile()

	m := NewMachine(instructions, os.Stdin, os.Stdout)
	if err := m.Execute(); err != nil {
		handleError(err)
	}
}

func handleError(err error) {
	fmt.Fprintf(os.Stderr, "error: %s\n", err)
	os.Exit(-1)
}
