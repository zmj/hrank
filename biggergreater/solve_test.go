package main

import (
	"hrank/test"
	"testing"
)

func Test0(t *testing.T) {
	test.Test(t, solve, "example0", `ba
no answer
hegf
dhkc
hcdk`)
}

func Test1(t *testing.T) {
	test.Test(t, solve, "example1", `lmon
no answer
no answer
acbd
abdc
fedcbabdc`)
}
