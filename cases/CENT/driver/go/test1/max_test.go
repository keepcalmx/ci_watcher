package main

import (
	"testing"
)

func TestMax(t *testing.T) {
	a, b := 1, 2
	if Max(a, b) != 2 {
		t.Errorf("Max(%d, %d) expect:%d got:%d", a, b, 2, Max(a, b))
	}
}
