package main

import "testing"

func TestHelloWorld(t *testing.T) {
	if helloworld() != "Hello World!! ff" {
		t.Fatal("Test fail")
	}
}
