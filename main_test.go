package main

import (
	"testing"

	go_greeting "github.com/pradiptaagus/go-greeting/v2"
)

func TestMain(t *testing.T) {
	result := go_greeting.SayHello("John Doe")
	if result != "Hello John Doe" {
		t.Errorf("Expected to be \"Hello John Doe\", got \"%s\"", result)
	}
}
