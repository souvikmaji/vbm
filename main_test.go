package main

import (
	"bytes"
	"io/ioutil"
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

func BenchmarkHelloWorld(b *testing.B) {
	file, _ := ioutil.ReadFile("samples/hello_world.b")
	benchmarkCompileCode(b, string(file))
}

func BenchmarkMandelbrot(b *testing.B) {
	file, _ := ioutil.ReadFile("samples/mandelbrot.b")
	benchmarkCompileCode(b, string(file))
}

func benchmarkCompileCode(b *testing.B, code string) {
	in := bytes.NewBufferString("")
	out := new(bytes.Buffer)

	for n := 0; n < b.N; n++ {
		compile(code, in, out)
	}
}
