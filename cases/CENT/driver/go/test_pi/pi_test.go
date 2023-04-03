package main

import "testing"

func TestPi(t *testing.T) {
	pi := Pi(1000)
	if pi != 3.141592653589793 {
		t.Errorf("Pi(100) expect:%f got:%f", 3.141592653589793, pi)
	}
}
