package recursion

import (
	"testing"
)

func TestMazeSolver(t *testing.T) {
	maze := []string{
		"xxxxxxxxxx x",
		"x        x x",
		"x        x x",
		"x xxxxxxxx x",
		"x          x",
		"x xxxxxxxxxx",
	}

	mazeResult := []Point{
		{10, 0},
		{10, 1},
		{10, 2},
		{10, 3},
		{10, 4},
		{9, 4},
		{8, 4},
		{7, 4},
		{6, 4},
		{5, 4},
		{4, 4},
		{3, 4},
		{2, 4},
		{1, 4},
		{1, 5},
	}

	result := MazeSolver(maze, "x", Point{10, 0}, Point{1, 5})
	if !equalPaths(result, mazeResult) {
		t.Errorf("Expected path: %v, but got: %v", mazeResult, result)
	}
}

func equalPaths(path1, path2 []Point) bool {
	if len(path1) != len(path2) {
		return false
	}

	for i := range path1 {
		if path1[i] != path2[i] {
			return false
		}
	}

	return true
}
