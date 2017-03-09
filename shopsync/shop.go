package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type fishmask int

type problem struct {
	allFish fishmask
	start   *node
	end     *node
}

type node struct {
	name  string
	sells fishmask
	edges []*edge
}

type edge struct {
	src  *node
	dest *node
	cost int
}

type path struct {
	cost int
	fish fishmask
	pos  *node
}

func main() {
	rdr := bufio.NewReader(os.Stdin)
	wr := bufio.NewWriter(os.Stdout)
	solve(rdr, wr)
}

func solve(rdr *bufio.Reader, wr *bufio.Writer) {
	defer wr.Flush()
	prob := (&parser{rdr}).parse()
	wr.WriteString(fmt.Sprintf("%v\n", prob.start.name))
	wr.WriteString(fmt.Sprintf("%v\n", prob.end.name))
}

type parser struct {
	rdr *bufio.Reader
}

func (p *parser) parse() *problem {
	prob := &problem{}
	n, m, k := p.firstLine()
	nodes := make([]*node, n)
	prob.allFish = (1 << uint(k)) - 1
	for i := 0; i < n; i++ {
		name := strconv.Itoa(i + 1)
		nodes[i] = &node{
			name:  name,
			sells: p.shopLine(),
		}
	}
	for i := 0; i < m; i++ {
		src, dest, cost := p.edgeLine()
		srcNode := nodes[src-1]
		edge := &edge{
			src:  srcNode,
			dest: nodes[dest-1],
			cost: cost,
		}
		srcNode.edges = append(srcNode.edges, edge)
	}
	prob.start = nodes[0]
	prob.end = nodes[len(nodes)-1]
	return prob
}

func (p *parser) firstLine() (n, m, k int) {
	line := p.line()
	return line[0], line[1], line[2]
}

func (p *parser) line() []int {
	line, _, _ := p.rdr.ReadLine()
	segments := strings.Split(string(line), " ")
	values := make([]int, len(segments))
	for i, s := range segments {
		v, _ := strconv.Atoi(s)
		values[i] = v
	}
	return values
}

func (p *parser) shopLine() fishmask {
	line := p.line()
	var fish fishmask
	for i := 0; i < line[0]; i++ {
		fish |= (1 << uint(line[i+1]))
	}
	return fish
}

func (p *parser) edgeLine() (src, dst, cost int) {
	line := p.line()
	return line[0], line[1], line[2]
}
