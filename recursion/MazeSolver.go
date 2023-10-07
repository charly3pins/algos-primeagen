package recursion

type Point struct {
	x int
	y int
}

func MazeSolver(maze []string, wall string, start Point, end Point) []Point {
	seen := make([][]bool, len(maze)) // init the seen matrix with the same dimensions as the maze
	for i := range seen {
		seen[i] = make([]bool, len(maze[0]))
	}

	var path []Point

	walk(maze, wall, start, end, seen, &path)

	return path
}

func walk(maze []string, wall string, curr Point, end Point, seen [][]bool, path *[]Point) bool {
	// Base cases
	// 1. off the map
	if curr.x < 0 || curr.x >= len(maze[0]) || curr.y < 0 || curr.y >= len(maze) {
		return false
	}

	// 2. it's a wall
	if string(maze[curr.y][curr.x]) == wall {
		return false
	}

	// 3. it's the end
	if curr.x == end.x && curr.y == end.y {
		*path = append(*path, end)
		return true
	}

	// 4. already visited
	if seen[curr.y][curr.x] {
		return false
	}

	// Recursive case
	// 1. pre condition
	seen[curr.y][curr.x] = true
	*path = append(*path, curr)

	// 2. recursion
	moves := []Point{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} // possible moves (up, down, left, right)

	for _, move := range moves {
		next := Point{curr.x + move.x, curr.y + move.y}

		if walk(maze, wall, next, end, seen, path) {
			return true
		}
	}

	// 3. post condition
	if len(*path) > 0 {
		*path = (*path)[:len(*path)-1] // remove the last element from the path
	}

	return false
}
