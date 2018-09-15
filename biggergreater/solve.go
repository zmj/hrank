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
	words, err := parse(r)
	if err != nil {
		return fmt.Errorf("parse: %v", err)
	}
	var answ string
	for i, word := range words {
		p, ok := permute(word)
		if ok {
			answ = p
		} else {
			answ = "no answer"
		}
		if i != 0 {
			fmt.Fprintln(w)
		}
		fmt.Fprint(w, answ)
	}
	return nil
}

func permute(s string) (string, bool) {
	var i, j int
outer:
	for i = len(s) - 1; i >= 0; i-- {
		for j = len(s) - 1; j > i; j-- {
			if s[i] < s[j] {
				break outer
			}
		}
	}
	if i < 0 {
		return "", false
	}
	b := []byte(s)
	b[i], b[j] = b[j], b[i]
	c := b[i+1:]
	for k, l := 0, len(c)-1; l > k; k, l = k+1, l-1 {
		c[k], c[l] = c[l], c[k]
	}
	return string(b), true
}

func parse(r io.Reader) ([]string, error) {
	p := parser.NewParser(r)
	n, err := p.Int()
	if err != nil {
		return nil, fmt.Errorf("count: %v", err)
	}
	values := make([]string, n)
	for i := 0; i < n; i++ {
		v, err := p.String()
		if err != nil {
			return nil, fmt.Errorf("value %v: %v", i, err)
		}
		values[i] = v
	}
	return values, nil
}
