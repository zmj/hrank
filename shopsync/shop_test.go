package main

import (
	"bufio"
	"os"
	"strconv"
	"testing"
)

func TestSolve0(t *testing.T) {
	answer := tSolve("example")
	if answer != 30 {
		t.Fail()
	}
}

func TestSolve5(t *testing.T) {
	answer := tSolve("example5")
	if answer != 1571 {
		t.Fail()
	}
}

func tSolve(filename string) int {
	file, _ := os.Open(filename)
	rdr := bufio.NewReader(file)
	wr := &reader{}
	solve(rdr, bufio.NewWriter(wr))
	answer, _ := strconv.Atoi(wr.lines[0])
	return answer
}

type reader struct {
	lines []string
}

func (r *reader) Write(b []byte) (int, error) {
	r.lines = append(r.lines, string(b))
	return len(b), nil
}
