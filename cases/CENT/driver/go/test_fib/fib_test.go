package main

import "testing"

func TestFib(t *testing.T) {
	fibs := GenFibonacci(45)

	if fibs[0] != 1 {
		t.Errorf("GenFibonacci(50) expect:%d got:%d", 1, fibs[0])
	}
}
