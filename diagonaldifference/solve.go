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
	values, err := parse(r)
	if err != nil {
		return fmt.Errorf("parse: %v", err)
	}
	sum1 := 0
	for i := 0; i < len(values); i++ {
		sum1 += values[i][i]
	}
	sum2 := 0
	for i := 0; i < len(values); i++ {
		sum2 += values[i][len(values)-1-i]
	}
	diff := sum1 - sum2
	if diff < 0 {
		diff *= -1
	}
	fmt.Fprint(w, diff)
	return nil
}

func parse(r io.Reader) ([][]int, error) {
	p := parser.NewParser(r)
	n, err := p.Int()
	if err != nil {
		return nil, fmt.Errorf("count: %v", err)
	}
	values := make([][]int, n)
	for i := 0; i < n; i++ {
		vals, err := p.Ints(n)
		if err != nil {
			return nil, fmt.Errorf("row %v: %v", i, err)
		}
		values[i] = vals
	}
	return values, nil
}
