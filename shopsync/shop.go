package main

import (
	"bufio"
	"container/heap"
	"math"
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
	name         string
	sells        fishmask
	edges        []*edge
	minCostToEnd int
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
	problem.setCosts()
	completePaths := make(chan path)
	go problem.enumerate(completePaths)
	solution := problem.match(completePaths)
	wr.WriteString(strconv.Itoa(solution))
}

func (problem problem) match(completePaths <-chan path) int {
	var paths []path
	for {
		p := <-completePaths
		if p.fish == problem.allFish {
			return p.cost
		}
		add := true
		for _, q := range paths {
			fish := p.fish | q.fish
			if fish == q.fish {
				add = false
				break
			}
			if fish == problem.allFish {
				cost := p.cost
				if q.cost > cost {
					cost = q.cost
				}
				return cost
			}
		}
		if add {
			paths = append(paths, p)
		}
	}
}

func (problem problem) enumerate(completePaths chan<- path) {
	start := path{pos: problem.start}
	paths := &pathHeap{start}
	best := make(map[step]int)
	for paths.Len() > 0 {
		p := heap.Pop(paths).(path)
		p.fish |= p.pos.sells
		p.name += ":" + p.pos.name
		if p.pos == problem.end {
			completePaths <- p
		}
		for _, edge := range p.pos.edges {
			nextCost := p.cost + edge.cost
			step := step{p.fish, edge.dest}
			bestCost, ok := best[step]
			if ok && bestCost < nextCost {
				continue
			}
			next := path{
				name: p.name,
				fish: p.fish,
				cost: nextCost,
				pos:  edge.dest,
			}
			best[step] = nextCost
			heap.Push(paths, next)
		}
	}
}

type step struct {
	fish fishmask
	pos  *node
}

type pathHeap []path

func (h pathHeap) Len() int {
	return len(h)
}

func (p path) val() int {
	return p.cost + p.pos.minCostToEnd
}

func (h pathHeap) Less(i, j int) bool {
	return h[i].val() < h[j].val()
}

func (h pathHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *pathHeap) Push(p interface{}) {
	*h = append(*h, p.(path))
}

func (h *pathHeap) Pop() interface{} {
	old := *h
	last := old[len(old)-1]
	*h = old[0 : len(old)-1]
	return last
}

func (problem *problem) setCosts() {
	var set func(*node)
	set = func(n *node) {
		for _, edge := range n.edges {
			if n.minCostToEnd+edge.cost < edge.dest.minCostToEnd {
				edge.dest.minCostToEnd = n.minCostToEnd + edge.cost
				set(edge.dest)
			}
		}
	}
	problem.end.minCostToEnd = 0
	set(problem.end)
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
			name:         strconv.Itoa(i + 1),
			sells:        p.shopLine(),
			minCostToEnd: math.MaxInt32,
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
		fish |= (1 << uint(line[i+1]-1))
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
