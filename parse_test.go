package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseObj(t *testing.T) {
	parse, err := LoadObjFile("./testdata/maze.obj", 100, 100)
	assert.Nil(t, err)

	_ = parse.InitWorld()

	parse.SetStart(6, 6)
	parse.SetEnd(5, 50)

	world := parse.Render()
	t.Logf("Output world\n%s", world)
}
