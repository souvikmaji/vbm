package main

import (
	"bytes"
	"testing"
)

const HelloWorld = `++++++++[>++++[>++>+++>+++>+<<<<-]>+> +>->>+[<]<-]>>.>---.+++++++ ..+ ++.>>.<-.<.+++.------.--------.>>+.>++.`

func TestCompileCode(t *testing.T) {
	in := bytes.NewBufferString("")
	out := new(bytes.Buffer)

	err := compile(HelloWorld, in, out)
	if err != nil {
		t.Errorf("error should not occur. got=%q", err)
	}

	output := out.String()
	if output != "Hello World!\n" {
		t.Errorf("output wrong. got=%q", output)
	}
}
