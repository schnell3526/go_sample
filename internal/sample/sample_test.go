package sample

import "testing"

func TestHelloWorld(t *testing.T) {
	expected := "Hello, World!"
	if HelloWorld() != expected {
		t.Errorf("Expected %s, got %s", expected, HelloWorld())
	}
}
