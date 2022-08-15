package main

import (
	"errors"
	"math"
	"strconv"
	"strings"
)

const (
	xIndex int = 0
	yIndex int = 2

	xMax float64 = 10
	yMax float64 = 10
)

type Point struct {
	X, Y float64
}

type Triangle struct {
	Points [3]*Point
}

func LoadIndices(items ...string) ([]int, error) {
	indices := make([]int, 0)
	for _, item := range items {
		_indices := strings.Split(item, "/")
		i, err := strconv.Atoi(_indices[0])
		if err != nil {
			return nil, err
		}
		indices = append(indices, i)
	}

	return indices, nil
}

func LoadPoint(coordinates ...string) (*Point, error) {
	xp, err := strconv.ParseFloat(coordinates[xIndex], 64)
	if err != nil {
		return nil, err
	}

	yp, err := strconv.ParseFloat(coordinates[yIndex], 64)
	if err != nil {
		return nil, err
	}

	return &Point{
		X: xp,
		Y: yp,
	}, nil
}

func (p *Point) Trans() (*Point, error) {
	xHalfLen := xMax / 2
	yHalfLen := yMax / 2

	if p.X > xHalfLen || p.X < -xHalfLen {
		return nil, errors.New("x out of bounds")
	}
	if p.Y > yHalfLen || p.Y < -yHalfLen {
		return nil, errors.New("y out of bounds")
	}
	if p.X < 0 {
		p.X = (xHalfLen*100 - math.Abs(p.X)*100) / 100
	} else {
		p.X = (xHalfLen*100 + math.Abs(p.X)*100) / 100
	}

	if p.Y > 0 {
		p.Y = (yHalfLen*100 - math.Abs(p.Y)*100) / 100
	} else {
		p.Y = (yHalfLen*100 + math.Abs(p.Y)*100) / 100
	}

	return p, nil
}

func (t *Triangle) IsInside(p Point) bool {
	var a, b, c float64

	ma := Point{X: p.X - t.Points[0].X, Y: p.Y - t.Points[0].Y}
	mb := Point{X: p.X - t.Points[1].X, Y: p.Y - t.Points[1].Y}
	mc := Point{X: p.X - t.Points[2].X, Y: p.Y - t.Points[2].Y}

	/*向量叉乘*/
	a = ma.X*mb.Y - ma.Y*mb.X
	b = mb.X*mc.Y - mb.Y*mc.X
	c = mc.X*ma.Y - mc.Y*ma.X

	if (a <= 0 && b <= 0 && c <= 0) || (a > 0 && b > 0 && c > 0) {
		return true
	}
	return false
}
