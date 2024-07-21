package ctx_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/jcbhmr/go-typescript/internal/ctx"
)

func TestRunModule(t *testing.T) {
	file, err := os.CreateTemp("", "hello-world")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(file.Name())
	file.WriteString("Hello world!")
	js := "import { readFileSync } from 'node:fs';\n" +
		"console.log(readFileSync(" + fmt.Sprintf("%q", file.Name()) + ", 'utf8'));"
	p, err := ctx.Eval(js)
	defer p.Free()
	if err != nil {
		t.Fatal(err)
	}
	m, err := ctx.Await(p)
	defer m.Free()
	if err != nil {
		t.Fatal(err)
	}
}
