package main

import (
	"testing"
)

func TestToRoman(t *testing.T) {
	_, err1 := ToRoman(0)
	if err1 != ErrOutOfRange {
		t.Errorf("ToRoman(0) expect error:%v got:%v", ErrOutOfRange, err1)
	}

	roman2, err2 := ToRoman(1)
	if err2 != nil {
		t.Errorf("ToRoman(1) expect nil error, got:%v", err2)
	}
	if roman2 != "I" {
		t.Errorf("ToRoman(1) expect:%s got:%s", "I", roman2)
	}
}
