package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseObj(t *testing.T) {
	parse, err := LoadObjFile("./testdata/maze.obj", 10, 10)
	assert.Nil(t, err)

	world := parse.RenderWorld(100, 100)
	t.Logf("Output world\n%s", world)
}
