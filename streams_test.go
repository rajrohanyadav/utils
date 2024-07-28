package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestContains(t *testing.T) {
	tests := map[string]struct{
		val interface{}
		itr []interface{}
		expected bool
	}{
		"empty list": {
			val: 10,
			itr: []interface{}{},
			expected: false,
		},
		"contains element": {
			val: 10,
			itr: []interface{}{1,4,10},
			expected: true,
		},
		"does not contain": {
			val: 10,
			itr: []interface{}{1,4},
			expected: false,
		},
	}

	for tn, td := range tests {
		t.Run(tn, func(tt *testing.T) {
			assert.Equal(tt, td.expected, Contains(td.val, td.itr))
		})
	}
}

// TODO: check if TestMap can be type independent
func TestMap(t *testing.T) {
	tests := map[string]struct{
		itr []int
		apply func(int) int 
		expected []int 
	}{
		"empty list": {
			itr: []int{},
			expected: []int{}, 
		},
		"multiply by 2": {
			itr: []int{1,4,10},
			apply: func(i int) int {return i * 2},
			expected: []int{2,8,20},
		},
	}

	for tn, td := range tests {
		t.Run(tn, func(tt *testing.T) {
			assert.Equal(tt, td.expected, Map(td.apply, td.itr)) 
		})
	}
}

func TestFilter(t *testing.T) {
	tests := map[string]struct{
		filterFunc func(interface{}) bool
		itr []interface{}
		expected []interface{}
	}{
		"empty list": {
			filterFunc: func(interface{}) bool {return true},
			itr: []interface{}{},
			expected: []interface{}{},
		},
		"filter odd": {
			filterFunc: func(x interface{}) bool {return x.(int)%2==1},
			itr: []interface{}{1,2,3,4,5,6},
			expected: []interface{}{1,3,5},
		},
	}
	
	for tn, td := range tests {
		t.Run(tn, func(tt *testing.T) {
			actual := Filter(td.filterFunc, td.itr)
			assert.Equal(tt, td.expected, actual)	
		})
	}
}

func TestReduce(t *testing.T) {}

func TestReverse(t *testing.T) {
	tests := map[string]struct{
		itr []interface{}
		expected []interface{} 
	}{
		"empty list": {
			itr: []interface{}{},
			expected: []interface{}{}, 
		},
		"short list": {
			itr: []interface{}{1,4,10},
			expected: []interface{}{10,4,1},
		},
		"long list": {
			itr: []interface{}{1,4,10,34,21,533,1253,62343},
			expected: []interface{}{62343,1253,533,21,34,10,4,1},
		},
	}

	for tn, td := range tests {
		t.Run(tn, func(tt *testing.T) {
			Reverse(td.itr)
			assert.Equal(tt, td.expected, td.itr) 
		})
	}
}
