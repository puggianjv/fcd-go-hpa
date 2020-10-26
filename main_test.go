package main

import "testing"

func TestDoSqrt(t *testing.T) {
	result := doSqrt()
	expected := 1000001
	if result != expected {
		t.Errorf("doSqrt() %d; want %d", result, expected)
	}
}
