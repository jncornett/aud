package quantize

import (
	"fmt"
	"testing"

	"github.com/jncornett/aud"
	"github.com/stretchr/testify/assert"
)

func TestTo8BitUnsigned(t *testing.T) {
	tests := []struct {
		expected, input aud.Sample
	}{
		{expected: 0, input: -1},
		{expected: 128, input: 0},
		{expected: 255, input: 1},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.input), func(t *testing.T) {
			assert.Equal(t, tt.expected, To8BitUnsigned(tt.input))
		})
	}
}

func TestTo16BitSigned(t *testing.T) {
	tests := []struct {
		expected, input aud.Sample
	}{
		{expected: -32768, input: -1},
		{expected: 0, input: 0},
		{expected: 32767, input: 1},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.input), func(t *testing.T) {
			assert.Equal(t, tt.expected, To16BitSigned(tt.input))
		})
	}
}

func TestTo24BitSigned(t *testing.T) {
	tests := []struct {
		expected, input aud.Sample
	}{
		{expected: -16777216, input: -1},
		{expected: 0, input: 0},
		{expected: 16777215, input: 1},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.input), func(t *testing.T) {
			assert.Equal(t, tt.expected, To24BitSigned(tt.input))
		})
	}
}
