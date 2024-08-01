package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMax(t *testing.T) {
	tests := map[string]struct{
		input []int
		expected int
	}{
		"one number": {
			input: []int{1},
			expected: 1,
		},
		"two numbers": {
			input: []int{1,2},
			expected: 2,
		},
		"more than two numbers": {
			input: []int{234, 523, 51, 532, 51, 533, 11, 0, 97},
			expected: 533,
		},
	}
	
	for tn, td := range tests {
		t.Run(tn, func(tt *testing.T) {
			assert.Equal(tt, td.expected, Max(td.input[0], td.input[1:]...))
		})
	}
}

func TestMin(t *testing.T) {
	tests := map[string]struct{
		input []int
		expected int
	}{
		"one number": {
			input: []int{1},
			expected: 1,
		},
		"two numbers": {
			input: []int{1,2},
			expected: 1,
		},
		"more than two numbers": {
			input: []int{234, 523, 51, 532, 51, 533, 11, 0, 97},
			expected: 0,
		},
	}
	
	for tn, td := range tests {
		t.Run(tn, func(tt *testing.T) {
			assert.Equal(tt, td.expected, Min(td.input[0], td.input[1:]...))
		})
	}
}
