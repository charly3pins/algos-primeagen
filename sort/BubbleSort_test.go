package sort

import (
	"fmt"
	"reflect"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	array := []int{9, 3, 7, 4, 69, 420, 42}
	ordered := []int{3, 4, 7, 9, 42, 69, 420}

	testCases := []struct {
		array []int
		want  []int
	}{
		{array, ordered},
	}
	for k, tt := range testCases {
		t.Run(fmt.Sprintf("test for bubble sort %d", k), func(t *testing.T) {
			if got := BubbleSort(tt.array); !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("Case %d: Wanted %v and obtained %v", k, tt.want, got)
			}
		})
	}
}
