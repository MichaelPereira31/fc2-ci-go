package main

import "testing"

func TestSum(t *testing.T) {
	total := Sum(10, 20)
	
	if total != 30 {
		t.Errorf("Result should be 30, got %d", total)
	}
}

func Sum(a int, b int) int {
	return a + b
}
}