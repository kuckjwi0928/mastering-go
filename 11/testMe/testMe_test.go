package testMe

import "testing"

func Test_s1(t *testing.T) {
	if s1("123456789") != 9 {
		t.Error(`s1("123456789") != 9`)
	}
	if s1("") != 0 {
		t.Error(`s1("") != 0`)
	}
}

func Test_f1(t *testing.T) {
	if f1(0) != 0 {
		t.Error(`f1(0) != 0`)
	}
	if f1(1) != 1 {
		t.Error(`f1(1) != 1`)
	}
	if f1(10) != 55 {
		t.Error(`f1(10) != 55`)
	}
}
