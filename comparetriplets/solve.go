package main

import (
	"fmt"
	"hrank/parser"
	"io"
	"os"
)

func main() {
	err := solve(os.Stdin, os.Stdout)
	if err != nil {
		fmt.Println(err)
	}
}

func solve(r io.Reader, w io.Writer) error {
	input, err := parse(r)
	if err != nil {
		return err
	}
	var a, b int
	for i := 0; i < 3; i++ {
		if input.alice[i] > input.bob[i] {
			a++
		}
		if input.bob[i] > input.alice[i] {
			b++
		}
	}
	fmt.Fprintf(w, "%v %v", a, b)
	return nil
}

type input struct {
	alice []int
	bob   []int
}

func parse(r io.Reader) (*input, error) {
	p := parser.NewParser(r)
	v, err := p.Ints(3)
	if err != nil {
		return nil, fmt.Errorf("line 1: %v", err)
	}
	input := &input{alice: v}
	v, err = p.Ints(3)
	if err != nil {
		return nil, fmt.Errorf("line 2: %v", err)
	}
	input.bob = v
	return input, nil
}
