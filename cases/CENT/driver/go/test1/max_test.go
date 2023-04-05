package main

import (
	"testing"
)

func TestMax1(t *testing.T) {
	a, b := 1, 2
	if Max(a, b) != 2 {
		t.Errorf("Max(%d, %d) expect:%d got:%d", a, b, 2, Max(a, b))
	}
}

func TestMax2(t *testing.T) {
	a, b := 1, 0
	if Max(a, b) != 1 {
		t.Errorf("Max(%d, %d) expect:%d got:%d", a, b, 2, Max(a, b))
	}
}

func TestMax3(t *testing.T) {
	a, b := 3, 2
	if Max(a, b) != 3 {
		t.Errorf("Max(%d, %d) expect:%d got:%d", a, b, 2, Max(a, b))
	}
}

func TestMax4(t *testing.T) {
	a, b := 3, 5
	if Max(a, b) != 5 {
		t.Errorf("Max(%d, %d) expect:%d got:%d", a, b, 2, Max(a, b))
	}
}
