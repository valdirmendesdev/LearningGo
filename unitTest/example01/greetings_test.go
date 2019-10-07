package main

import (
	"testing"

	"gotest.tools/assert"
)

func TestSayHello(t *testing.T) {
	greeting := SayHello("Valdir")
	assert.Assert(greeting, "Hello, Valdir")
}
