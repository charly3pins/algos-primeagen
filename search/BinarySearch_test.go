package search

import (
	"fmt"
	"testing"
)

func TestBinarySearch(t *testing.T) {
	array := []int{1, 3, 4, 69, 71, 81, 90, 99, 420, 1337, 69420}
	testCases := []struct {
		needle int
		want   bool
	}{
		{69, true},
		{1336, false},
		{69420, true},
		{69421, false},
		{1, true},
		{0, false},
	}
	for _, tt := range testCases {
		t.Run(fmt.Sprintf("test for needle %d", tt.needle), func(t *testing.T) {
			if got := BinarySearch(array, tt.needle); got != tt.want {
				t.Fatalf("Needle %d: Wanted %v and obtained %v", tt.needle, tt.want, got)
			}
		})
	}
}
