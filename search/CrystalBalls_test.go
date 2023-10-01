package search

import (
	"fmt"
	"math"
	"math/rand"
	"testing"
)

func TestCrystalBalls(t *testing.T) {
	index := int(math.Floor(rand.Float64() * 1000))
	arrayBreak := []bool{}
	arrayNoBreak := []bool{}
	for i := 0; i < 1000; i++ {
		arrayBreak = append(arrayBreak, false)
		arrayNoBreak = append(arrayNoBreak, false)
	}
	for i := index; i < 1000; i++ {
		arrayBreak[i] = true
	}
	testCases := []struct {
		array []bool
		want  int
	}{
		{arrayBreak, index},
		{arrayNoBreak, -1},
	}
	for k, tt := range testCases {
		t.Run(fmt.Sprintf("test for crystal balls %d", k), func(t *testing.T) {
			if got := CrystalBalls(tt.array); got != tt.want {
				t.Fatalf("Case %d: Wanted %v and obtained %v", k, tt.want, got)
			}
		})
	}
}
