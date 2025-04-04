package levels

type Level struct {
	Maze   MazeConfig   `yaml:"maze"`
	Player PlayerConfig `yaml:"player"`
	Exit   ExitConfig   `yaml:"exit"`
}

type MazeConfig struct {
	Width    int `yaml:"width"`
	Height   int `yaml:"height"`
	CellSize int `yaml:"cell_size"`
}

type PlayerConfig struct {
	Size int `yaml:"size"`
}

type ExitConfig struct {
	Position Coordinate `yaml:"position"`
}

type Coordinate struct {
	X int `yaml:"x"`
	Y int `yaml:"y"`
}
