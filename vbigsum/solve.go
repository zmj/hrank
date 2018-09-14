package main

import (
	"fmt"
	"hrank/parser"
	"io"
	"os"
	"strconv"
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
	var sum uint64
	for _, x := range values {
		sum += x
	}
	w.Write([]byte(strconv.FormatUint(sum, 10)))
	return nil
}

func parse(r io.Reader) ([]uint64, error) {
	p := parser.NewParser(r)
	n, err := p.Int()
	if err != nil {
		return nil, fmt.Errorf("read count: %v", err)
	}
	return p.Uint64s(n)
}
