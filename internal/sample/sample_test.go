package sample

import "testing"

func TestHelloWorld(t *testing.T) {
	expected := "Hello, World!"
	if HelloWorld() != expected {
		t.Errorf("Expected %s, got %s", expected, HelloWorld())
	}
}

func TestAdd(t *testing.T) {
	expected := 3
	if Add(1, 2) != expected {
		t.Errorf("Expected %d, got %d", expected, Add(1, 2))
	}
}

func TestSubtrac(t *testing.T) {
	expected := -1
	if Subtract(1, 2) != expected {
		t.Errorf("Expected %d, got %d", expected, Subtract(1, 2))
	}
}
