package test

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"testing"
)

type Solve func(io.Reader, io.Writer) error

func Test(t *testing.T, solve Solve, infile, expected string) {
	t.Helper()
	f, err := os.Open(infile)
	if err != nil {
		t.Error(err)
		return
	}
	defer f.Close()
	buf := &bytes.Buffer{}
	err = solve(f, buf)
	if err != nil {
		t.Error(err)
		return
	}
	actual := string(buf.Bytes())
	if actual != expected {
		t.Error(fmt.Errorf("expected: '%v' actual: '%v'", expected, actual))
	}
}
