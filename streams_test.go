package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestContains(t *testing.T) {
	tests := []struct{
		testName string
		val interface{}
		itr []interface{}
		expected bool
	}{
		{
			testName: "empty list",
			val: 10,
			itr: []interface{}{},
			expected: false,
		},
		{
			testName: "contains element",
			val: 10,
			itr: []interface{}{1,4,10},
			expected: true,
		},
		{
			testName: "does not contain",
			val: 10,
			itr: []interface{}{1,4},
			expected: false,
		},
	}

	for _, test := range tests {
		t.Run(test.testName, func(tt *testing.T) {
			assert.Equal(tt, test.expected, Contains(test.val, test.itr))
		})
	}
}

// TODO: check if TestMap can be type independent
func TestMap(t *testing.T) {
	tests := []struct{
		testName string
		itr []int
		apply func(int) int 
		expected []int 
	}{
		{
			testName: "empty list",
			itr: []int{},
			expected: []int{}, 
		},
		{
			testName: "multiply by 2",
			itr: []int{1,4,10},
			apply: func(i int) int {return i * 2},
			expected: []int{2,8,20},
		},
	}

	for _, test := range tests {
		t.Run(test.testName, func(tt *testing.T) {
			assert.Equal(tt, test.expected, Map(test.apply, test.itr)) 
		})
	}
}

func TestFilter(t *testing.T) {}

func TestReduce(t *testing.T) {}

func TestReverse(t *testing.T) {
	tests := []struct{
		testName string
		itr []interface{}
		expected []interface{} 
	}{
		{
			testName: "empty list",
			itr: []interface{}{},
			expected: []interface{}{}, 
		},
		{
			testName: "short list",
			itr: []interface{}{1,4,10},
			expected: []interface{}{10,4,1},
		},
		{
			testName: "long list",
			itr: []interface{}{1,4,10,34,21,533,1253,62343},
			expected: []interface{}{62343,1253,533,21,34,10,4,1},
		},
	}

	for _, test := range tests {
		t.Run(test.testName, func(tt *testing.T) {
			Reverse(test.itr)
			assert.Equal(tt, test.expected, test.itr) 
		})
	}
}
