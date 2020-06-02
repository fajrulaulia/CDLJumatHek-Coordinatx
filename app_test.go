package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSomething(t *testing.T) {
	output := Slice("2 4 1 2 5 3-3 2 3 1")
	assert.Equal(t, output, 6, "they should be equal")
	output = Slice("1 1-1 1")
	assert.Equal(t, output, 1, "they should be equal")

}
