package main

import (
	"hrank/test"
	"testing"
)

func Test0(t *testing.T) {
	test.Test(t, solve, "example0", "7")
}

func Test1(t *testing.T) {
	test.Test(t, solve, "example1", "4")
}

func Test3(t *testing.T) {
	test.Test(t, solve, "example3", "20")
}
