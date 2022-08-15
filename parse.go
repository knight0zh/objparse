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

	Blank uint8 = 0
	Block uint8 = 1
	Start uint8 = 8
	End   uint8 = 9
)

var (
	parts = regexp.MustCompile("[ \t]+")
	enter = "\n"
)

type ObjParse struct {
	xLen, yLen int
	pix        []uint8
	Triangles  []*Triangle
}

func LoadObjFile(filename string, xLen, yLen int) (p *ObjParse, err error) {
	buf, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var (
		points  = make(map[int]*Point)
		indices = make([][]int, 0)
	)

	p = &ObjParse{
		xLen:      xLen,
		yLen:      yLen,
		pix:       make([]uint8, 0),
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
			_point, err := point.Trans()
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

func (p *ObjParse) InitWorld() error {
	for y := 0; y < p.yLen; y++ {
		for x := 0; x < p.xLen; x++ {
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
				p.pix = append(p.pix, Block)
			} else {
				p.pix = append(p.pix, Blank)

			}
		}
	}
	return nil
}

func (p *ObjParse) pixOffset(x, y int) int {
	return y*p.xLen + x
}

func (p *ObjParse) SetStart(x, y int) {
	p.pix[p.pixOffset(x, y)] = Start
}

func (p *ObjParse) SetEnd(x, y int) {
	p.pix[p.pixOffset(x, y)] = End
}

func (p *ObjParse) Render() (out string) {
	for i, v := range p.pix {
		if (i+1)%p.xLen == 0 {
			out += "\n"
		} else {
			out += p.TransOut(v)
		}
	}
	return
}

func (p *ObjParse) TransOut(str uint8) (out string) {
	switch str {
	case Blank:
		out = "."
	case Block:
		out = "X"
	case Start:
		out = "F"
	case End:
		out = "T"
	}

	return
}

func (p *ObjParse) RenderWorld() (s string) {
	for y := 0; y < p.yLen; y++ {
		for x := 0; x < p.xLen; x++ {
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
