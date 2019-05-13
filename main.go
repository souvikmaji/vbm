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

	fileName := args[0]
	code, err := ioutil.ReadFile(fileName)
	handleError(err)

	compiler := NewCompiler(string(code))
	instructions := compiler.Compile()

	m := NewMachine(instructions, os.Stdin, os.Stdout)
	err = m.Execute()
	handleError(err)
}

func handleError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
		os.Exit(-1)
	}
}
