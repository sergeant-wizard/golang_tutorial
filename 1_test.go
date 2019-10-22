package main

import (
	"testing"
)

func TestSqrt(t *testing.T) {
	sqrt2 := Sqrt(2)
	if !isClose(sqrt2*sqrt2, 2) {
		t.Errorf("Wrong Answer")
	}
}
