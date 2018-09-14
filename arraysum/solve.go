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
		return err
	}
	sum := 0
	for _, x := range values {
		sum += x
	}
	w.Write([]byte(strconv.Itoa(sum)))
	return nil
}

func parse(r io.Reader) ([]int, error) {
	p := parser.NewParser(r)
	size, err := p.Int()
	if err != nil {
		return nil, fmt.Errorf("read size: %v", err)
	}
	return p.Ints(size)
}
