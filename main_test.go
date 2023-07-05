package main

import (
	"testing"
)

func TestPutToken01(t *testing.T) {
	b := &Board{
		tokens: [9]int{0, 0, 0, 0, 0, 0, 0, 0, 0},
	}
	b.put(1, 1, "o")
	if b.get(1, 1) != "o" {
		t.Errorf("....")
	}
}
