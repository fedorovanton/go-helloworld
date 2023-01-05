package main

import (
	"learn/examples"
	"testing"
)

func TestMultiply(t *testing.T) {
	r := examples.Sum(5, 10)
	if r != 15 {
		t.Error("Ожидается 15, получено ", r)
	}
}
