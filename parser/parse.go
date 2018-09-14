package parser

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

type Parser struct {
	scanner *bufio.Scanner
}

func (p *Parser) Int() (int, error) {
	s, err := p.readLine()
	if err != nil {
		return 0, err
	}
	return strconv.Atoi(s)
}

func (p *Parser) Ints(count int) ([]int, error) {
	s, err := p.readLine()
	if err != nil {
		return nil, err
	}
	ss := strings.Split(s, " ")
	if len(ss) != count {
		return nil, fmt.Errorf("expected %v values, got %v", count, len(ss))
	}
	values := make([]int, count)
	var v int
	for i, s := range ss {
		v, err = strconv.Atoi(s)
		values[i] = v
	}
	return values, err
}

func (p *Parser) Uint64s(n int) ([]uint64, error) {
	s, err := p.readLine()
	if err != nil {
		return nil, err
	}
	ss := strings.Split(s, " ")
	if len(ss) != n {
		return nil, fmt.Errorf("expected %v values, got %v", n, len(ss))
	}
	values := make([]uint64, n)
	var v uint64
	for i, s := range ss {
		v, err = strconv.ParseUint(s, 10, 64)
		values[i] = v
	}
	return values, err
}

func NewParser(r io.Reader) *Parser {
	return &Parser{bufio.NewScanner(r)}
}

func (p *Parser) readLine() (string, error) {
	ok := p.scanner.Scan()
	if !ok {
		if p.scanner.Err() != nil {
			return "", p.scanner.Err()
		}
		return "", io.EOF
	}
	return p.scanner.Text(), nil
}
