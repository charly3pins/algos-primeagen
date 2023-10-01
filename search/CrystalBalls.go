package search

import "math"

// Given two crystal balls that will break if dropped from high enough
// distance, determine the exact spot in which it will break in the most
// optimized way.

func CrystalBalls(breaks []bool) int {
	jump := int(math.Floor(math.Sqrt(float64(len(breaks))))) // we jump Sqrt of N

	i := jump
	for ; i < len(breaks); i += jump {
		if breaks[i] { // when the 1st ball breaks
			break
		}
	}

	i -= jump // we go back only Sqrt of N

	for j := 0; j < jump && i < len(breaks)-1; j++ { // and walk linearly the array
		i++
		if breaks[i] { // until it breaks the 2nd ball
			return i
		}
	}

	return -1 // if not, return -1 as "not found"
}
