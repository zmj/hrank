package main

import (
	"bufio"
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
	dest *node
	cost int
}

type path struct {
	name string
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
	problem := (&parser{rdr}).parse()
	completePaths := make(chan *path, 100)
	go enumerate(completePaths)
	solution := match(completePaths)
	wr.WriteString(strconv.Itoa(solution))
}

func match(completePaths <-chan *path) int {

}

func enumerate(completePaths chan<- *path) {

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
		nodes[i] = &node{
			name:  strconv.Itoa(i + 1),
			sells: p.shopLine(),
		}
	}
	for i := 0; i < m; i++ {
		src, dest, cost := p.edgeLine()
		srcNode := nodes[src-1]
		destNode := nodes[dest-1]
		srcNode.edges = append(srcNode.edges, &edge{
			dest: destNode,
			cost: cost,
		})
		destNode.edges = append(destNode.edges, &edge{
			dest: srcNode,
			cost: cost,
		})
	}
	prob.start = nodes[0]
	prob.end = nodes[len(nodes)-1]
	return prob
}

func (p *parser) firstLine() (n, m, k int) {
	line := p.line()
	n, m, k = line[0], line[1], line[2]
	if n < 2 || m < 1 || k < 0 {
		panic("param out of range")
	}
	return
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

func (p *parser) edgeLine() (src, dest, cost int) {
	line := p.line()
	src, dest, cost = line[0], line[1], line[2]
	if src == dest {
		panic("circular edge")
	}
	return
}
