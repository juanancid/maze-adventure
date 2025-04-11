package levels

type Level struct {
	Number int
	Maze   MazeConfig   `yaml:"maze"`
	Player PlayerConfig `yaml:"player"`
	Exit   ExitConfig   `yaml:"exit"`
}

type MazeConfig struct {
	Cols int `yaml:"cols"`
	Rows int `yaml:"rows"`
}

type PlayerConfig struct {
	Size int `yaml:"size"`
}

type ExitConfig struct {
	Position Coordinate `yaml:"position"`
	Size     int        `yaml:"size"`
}

type Coordinate struct {
	X int `yaml:"x"`
	Y int `yaml:"y"`
}
