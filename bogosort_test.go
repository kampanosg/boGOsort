package bogosort_test

import (
	"testing"

	bogosort "github.com/kampanosg/boGOsort"
)

func TestIsOrdered(t *testing.T) {
	tests := []struct {
		name     string
		slice    []int
		expected bool
	}{
		{name: "success when empty empty slice", slice: []int{}, expected: true},
		{name: "success when single element slice", slice: []int{1}, expected: true},
		{name: "success", slice: []int{1, 2, 3, 4, 5}, expected: true},
		{name: "failure when reverse order", slice: []int{5, 4, 3, 2, 1}, expected: false},
		{name: "success when duplicate elements", slice: []int{1, 2, 2, 3, 4, 5}, expected: true},
		{name: "failure", slice: []int{1, 2, 2, 1, 3, 4, 5}, expected: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tc := tt
			t.Parallel()
			if isSorted := bogosort.IsSorted(tc.slice); isSorted != tc.expected {
				t.Errorf("expected IsSorted(%v) to be %v, but got %v", tc.slice, tc.expected, isSorted)
			}
		})
	}
}
