package main

import (
	"bytes"
	"fmt"
	"os"
	"testing"
)

func TestSolve0(t *testing.T) {
	testSolve(t, "example0", "5")
}

func testSolve(t *testing.T, infile, expected string) {
	t.Helper()
	f, err := os.Open(infile)
	if err != nil {
		t.Error(err)
	}
	defer f.Close()
	buf := &bytes.Buffer{}
	err = solve(f, buf)
	if err != nil {
		t.Error(err)
	}
	actual := string(buf.Bytes())
	if actual != expected {
		t.Error(fmt.Errorf("expected: '%v' actual: '%v'", expected, actual))
	}
}
