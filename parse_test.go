package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseObj(t *testing.T) {
	parse, err := LoadObjFile("./testdata/maze.obj", 10, 10)
	assert.Nil(t, err)
	for _, triangle := range parse.Triangles {
		fmt.Println(triangle)
	}
	str := parse.RenderWorld(100, 100)
	fmt.Println(str)
}
