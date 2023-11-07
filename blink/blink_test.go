package blink

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var b = New()

func TestBlink_Put(t *testing.T) {
	b.Put("key", "value")

	value, ok := b.Get("key")
	assert.True(t, ok)
	assert.Equal(t, "value", value)
}
