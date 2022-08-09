package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"
)

const (
	Vertex = "v"
	Face   = "f"
)

var (
	parts = regexp.MustCompile("[ \t]+")
	enter = "\n"
)

type ObjParse struct {
	Triangles []*Triangle
}

func LoadObjFile(filename string, xMax, yMax float64) (p *ObjParse, err error) {
	buf, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var (
		points  = make(map[int]*Point)
		indices = make([][]int, 0)
	)

	p = &ObjParse{
		Triangles: make([]*Triangle, 0),
	}

	lines := strings.Split(string(buf), enter)
	for n, line := range lines {
		line = strings.TrimSpace(line)
		if len(line) == 0 {
			continue
		}
		if strings.HasPrefix(line, "#") {
			continue
		}
		items := parts.Split(line, -1)

		switch items[0] {
		case Vertex:
			if len(items) != 4 {
				err = errors.New(fmt.Sprintf("%s:%d: invalid count of vertex or face\n", filename, n+1))
				break
			}
			point, err := LoadPoint(items[1:]...)
			if err != nil {
				err = fmt.Errorf("%s:%d: invalid float: %w\n", filename, n+1, err)
				break
			}
			_point, err := point.Trans(xMax, yMax)
			if err != nil {
				err = fmt.Errorf("%s:%d: invalid trans: %w\n", filename, n+1, err)
				break
			}
			points[n] = _point
		case Face:
			if len(items) != 4 {
				err = errors.New(fmt.Sprintf("%s:%d: invalid count of vertex or face\n", filename, n+1))
				break
			}
			_indices, err := LoadIndices(items[1:]...)
			if err != nil {
				err = fmt.Errorf("%s:%d: invalid indices: %w\n", filename, n+1, err)
				break
			}
			indices = append(indices, _indices)

		}
	}

	for _, index := range indices {
		if len(index) != 3 {
			continue
		}
		p.Triangles = append(p.Triangles, &Triangle{
			Points: [3]*Point{
				points[index[0]],
				points[index[1]],
				points[index[2]],
			},
		})
	}

	return p, err
}

func (p *ObjParse) RenderWorld(xLen, zLen int) (s string) {
	for y := 0; y < zLen; y++ {
		for x := 0; x < xLen; x++ {
			isBlock := true
			for _, triangle := range p.Triangles {
				if triangle.IsInside(Point{
					X: float64(x) / 10,
					Y: float64(y) / 10,
				}) {
					isBlock = false
					break
				}
			}
			if isBlock {
				s += "X"
			} else {
				s += "."
			}
		}
		s += "\n"
	}
	return
}
