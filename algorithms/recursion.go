package algorithms

type Point struct {
	X int
	Y int
}

var posbilePaths = []Point{
	{0, 1},
	{0, -1},
	{1, 0},
	{-1, 0},
}

func walk(maze []string, wall byte, curr, end Point, seen [][]bool, path *[]Point) bool {
	if curr.X < 0 || curr.X >= len(maze[0]) {
		return false
	}
	if curr.Y < 0 || curr.Y >= len(maze) {
		return false
	}
	if maze[curr.Y][curr.X] == wall {
		return false
	}
	if curr.X == end.X && curr.Y == end.Y {
		*path = append(*path, curr)
		return true
	}
	if seen[curr.Y][curr.X] {
		return false
	}
	seen[curr.Y][curr.X] = true
	*path = append(*path, curr)
	for i := 0; i < len(posbilePaths); i++ {
		if walk(maze, wall, Point{curr.X + posbilePaths[i].X, curr.Y + posbilePaths[i].Y}, end, seen, path) {
			return true
		}
	}
	*path = (*path)[:len(*path)-1]
	return false
}
func MazeSolver(maze []string, wall byte, start, end Point) []Point {
	seen := make([][]bool, len(maze))
	for i := range maze {
		seen[i] = make([]bool, len(maze[i]))
	}
	path := []Point{}

	walk(maze, wall, start, end, seen, &path)

	return path
}
