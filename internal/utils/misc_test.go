package utils

import (
	"slices"
	"testing"
)

func TestUniqueIntSlice(t *testing.T) {
	tests := []struct {
		input    []int
		expected []int
	}{
		{[]int{}, []int{}},
		{[]int{1}, []int{1}},
		{[]int{2, 1}, []int{2, 1}},
		{[]int{1, 2, 1}, []int{1, 2}},
		{[]int{3, 3, 3, 4, 4, 5, 1, 3}, []int{3, 4, 5, 1}},
	}

	for _, test := range tests {
		if result := UniqueIntSlice(test.input); slices.Equal(result, test.expected) == false {
			t.Errorf("UniqueIntSlice(%v) = %v, expected %v", test.input, result, test.expected)
		}
	}
}
