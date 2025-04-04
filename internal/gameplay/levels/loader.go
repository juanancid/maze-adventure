package levels

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

func Load(levelNumber int) (*Level, error) {
	// TODO: Remove this temporary hack
	if levelNumber > 2 {
		levelNumber = 0
	}

	filename := fmt.Sprintf("internal/gameplay/levels/configs/level_%02d.yaml", levelNumber)
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var lvl Level
	if err := yaml.Unmarshal(data, &lvl); err != nil {
		return nil, err
	}

	return &lvl, nil
}
