package main

import (
	"container/heap"
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
	a, err := parse(r)
	if err != nil {
		return err
	}
	h := &mHeap{&matrix{a: a}}
	var answer *matrix
	for h.Len() > 0 {
		src := heap.Pop(h).(*matrix)
		fmt.Println(src)
		if src.magic() {
			answer = src
			break
		}
		for _, e := range src.edges() {
			dest := e.next()
			heap.Push(h, dest)
		}
	}
	fmt.Fprint(w, answer.cost)
	return nil
}

type matrix struct {
	a    [][]int
	cost int
	prev *edge
	h    int
}

func (m *matrix) heuristic() int {
	if m.h > 0 {
		return m.h
	}
	m.h = 9 - m.distinct()
	h := maxDiff(m.row(0), m.row(1), m.row(2))
	if h > m.h {
		m.h = h
	}
	h = maxDiff(m.col(0), m.col(1), m.col(2))
	if h > m.h {
		m.h = h
	}
	h = maxDiff(m.d1(), m.d2())
	if h > m.h {
		m.h = h
	}
	return m.h
}

func maxDiff(xs ...int) int {
	max := 0
	for i, x := range xs {
		for _, y := range xs[i:] {
			diff := x - y
			if diff < 0 {
				diff = y - x
			}
			if diff > max {
				max = diff
			}
		}
	}
	return max
}

func (m *matrix) magic() bool {
	if m.distinct() != 9 {
		return false
	}
	tests := []func() int{m.d1, m.d2}
	for i := 0; i < len(m.a); i++ {
		ii := i
		tests = append(tests, func() int { return m.row(ii) })
		tests = append(tests, func() int { return m.col(ii) })
	}
	sum := tests[0]()
	for i := 1; i < len(tests); i++ {
		x := tests[i]()
		if x != sum {
			return false
		}
	}
	return true
}

func (m *matrix) distinct() int {
	n := make([]int, 9)
	for i := 0; i < len(m.a); i++ {
		for j := 0; j < len(m.a); j++ {
			x := m.a[i][j]
			n[x-1]++
		}
	}
	d := 0
	for _, c := range n {
		if c == 1 {
			d++
		}
	}
	return d
}

func (m *matrix) d1() int {
	sum := 0
	for i := 0; i < len(m.a); i++ {
		sum += m.a[i][i]
	}
	return sum
}

func (m *matrix) d2() int {
	sum := 0
	for i := 0; i < len(m.a); i++ {
		sum += m.a[i][len(m.a)-1-i]
	}
	return sum
}

func (m *matrix) row(i int) int {
	sum := 0
	for j := 0; j < len(m.a); j++ {
		sum += m.a[i][j]
	}
	return sum
}

func (m *matrix) col(j int) int {
	sum := 0
	for i := 0; i < len(m.a); i++ {
		sum += m.a[i][j]
	}
	return sum
}

func (m *matrix) edges() []*edge {
	var es []*edge
	for i := 0; i < len(m.a); i++ {
	jloop:
		for j := 0; j < len(m.a); j++ {
			for p := m.prev; p != nil; p = p.prev.prev {
				if i == p.i && j == p.j {
					continue jloop
				}
			}
			for x := 1; x <= 9; x++ {
				if m.a[i][j] == x {
					continue
				}
				e := &edge{i: i, j: j, x: x, prev: m}
				es = append(es, e)
			}
		}
	}
	return es
}

type edge struct {
	i, j, x int
	prev    *matrix
}

func (e *edge) next() *matrix {
	m := &matrix{
		a:    make([][]int, len(e.prev.a)),
		cost: e.prev.cost,
		prev: e,
	}
	for i, a := range e.prev.a {
		if e.i != i {
			m.a[i] = e.prev.a[i]
			continue
		}
		a2 := make([]int, len(a))
		copy(a2, a)
		m.a[i] = a2
	}
	old := m.a[e.i][e.j]
	m.a[e.i][e.j] = e.x
	var cost int
	if old < e.x {
		cost = e.x - old
	} else {
		cost = old - e.x
	}
	m.cost += cost
	m.prev = e
	return m
}

type mHeap []*matrix

func (h mHeap) Len() int {
	return len(h)
}

func (h mHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h mHeap) Less(i, j int) bool {
	gi := h[i].cost + h[i].heuristic()
	gj := h[j].cost + h[j].heuristic()
	return gi < gj
}

func (h *mHeap) Push(x interface{}) {
	*h = append(*h, x.(*matrix))
}

func (h *mHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func parse(r io.Reader) ([][]int, error) {
	p := parser.NewParser(r)
	values := make([][]int, 3)
	ln := func(i int) error {
		row, err := p.Ints(3)
		if err != nil {
			return err
		}
		values[i] = row
		return nil
	}
	for i := 0; i < 3; i++ {
		err := ln(i)
		if err != nil {
			return nil, fmt.Errorf("line %v: %v", i, err)
		}
	}
	return values, nil
}
