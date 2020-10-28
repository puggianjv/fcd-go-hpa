package main

import "testing"

func TestGreeting(t *testing.T) {
	result := greeting("hello")
	expected := "<b>hello</b>"
	if result != expected {
		t.Errorf("greeting('hello') %s; want '%s'", result, expected)
	}
}

func TestDoSqrt(t *testing.T) {
	result := doSqrt()
	expected := 1000000
	if result != expected {
		t.Errorf("doSqrt() %d; want '%d'", result, expected)
	}
}
