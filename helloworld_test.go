package main

import "testing"

func TestHelloWorldReturn(t *testing.T) {
	output := SayHello()
	if output != "Hello World!" {
		t.Error("Expected Hello World!, got ", output)
	}
}
