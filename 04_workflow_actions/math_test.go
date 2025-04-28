package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	actual := Add(1, 2)
	expected := 3
	assert.Equal(t, expected, actual)
}
