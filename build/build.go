package build

// Maze is a 2D array of type rune
type Maze [][]rune

// Init initializes the 2d array
func Init(x, y int) Maze {
	m := make([][]rune, y)
	for i := range m {
		m[i] = make([]rune, x)
	}
	return m
}

// Build builds the maze
func (m *Maze) Build(rawMap string) {
	x := 0
	y := 0
	for _, r := range rawMap {
		if r == '\n' {
			x = 0
			y++
		} else {
			(*m)[y][x] = r
			x++
		}
	}
}
