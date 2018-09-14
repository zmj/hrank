package main

import (
	"bufio"
	"fmt"
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
	input, err := parse(r)
	if err != nil {
		return err
	}
	answer := input.a + input.b
	_, err = w.Write([]byte(strconv.Itoa(answer)))
	return err
}

type input struct {
	a int
	b int
}

func parse(r io.Reader) (*input, error) {
	scanner := bufio.NewScanner(r)
	rd := func(val *int) error {
		ok := scanner.Scan()
		if !ok {
			if scanner.Err() != nil {
				return scanner.Err()
			}
			return io.EOF
		}
		s := scanner.Text()
		i, err := strconv.Atoi(s)
		if err != nil {
			return fmt.Errorf("parse %v: %v", s, err)
		}
		*val = i
		return nil
	}
	input := &input{}
	err := rd(&input.a)
	if err != nil {
		return nil, fmt.Errorf("read a: %v", err)
	}
	err = rd(&input.b)
	if err != nil {
		return nil, fmt.Errorf("read b: %v", err)
	}
	return input, nil
}
