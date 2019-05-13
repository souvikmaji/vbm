package main

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fileName := getFileName()
	code, err := ioutil.ReadFile(fileName)
	handleError(err)

	err = compile(string(code), os.Stdin, os.Stdout)
	handleError(err)
}

func compile(code string, input io.Reader, output io.Writer) error {
	compiler := NewCompiler(code)
	instructions := compiler.Compile()

	m := NewMachine(instructions, input, output)
	return m.Execute()
}

func getFileName() string {
	flag.Parse()
	args := flag.Args()

	if len(args) != 1 {
		handleError(errors.New("no filename given"))
	}

	return args[0]
}

func handleError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
		os.Exit(-1)
	}
}
