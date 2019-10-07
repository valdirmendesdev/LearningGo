package main

import "testing"

func TestSum(t *testing.T) {
	total := Sum(2, 2)
	if total != 4 {
		t.Errorf("Should be 4! Actual value %d", total)
	}
}
