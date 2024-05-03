package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxInt(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected int
	}{
		{
			name:     "first number larger",
			input:    []int{5, 4},
			expected: 5,
		},
		{
			name:     "second number larger",
			input:    []int{3, 4},
			expected: 4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, MaxInt(tt.input[0], tt.input[1]))
		})
	}
}
