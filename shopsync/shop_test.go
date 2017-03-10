package main

import "testing"
import "os"
import "bufio"
import "strconv"

func TestSolve(t *testing.T) {
	file, _ := os.Open("example")
	rdr := bufio.NewReader(file)
	wr := &reader{}
	solve(rdr, bufio.NewWriter(wr))
	answer, _ := strconv.Atoi(wr.lines[0])
	if answer != 30 {
		t.Fail()
	}
}

type reader struct {
	lines []string
}

func (r *reader) Write(b []byte) (int, error) {
	r.lines = append(r.lines, string(b))
	return len(b), nil
}
