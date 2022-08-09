package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPoint_Trans(t *testing.T) {
	tests := []struct {
		point     *Point
		wantPoint *Point
	}{
		{
			point: &Point{
				X: -1.16,
				Y: -2.83,
			},
			wantPoint: &Point{
				X: 3.84,
				Y: 7.83,
			},
		},
		{
			point: &Point{
				X: 1.16,
				Y: 2.83,
			},
			wantPoint: &Point{
				X: 6.16,
				Y: 2.17,
			},
		},
		{
			point: &Point{
				X: -4.16,
				Y: 4.83,
			},
			wantPoint: &Point{
				X: 0.84,
				Y: 0.17,
			},
		},
		{
			point: &Point{
				X: 4.9,
				Y: -4.9,
			},
			wantPoint: &Point{
				X: 9.9,
				Y: 9.9,
			},
		},
		{
			point: &Point{
				X: 0,
				Y: 0,
			},
			wantPoint: &Point{
				X: 5,
				Y: 5,
			},
		},
	}

	for _, item := range tests {
		point, _ := item.point.Trans(10, 10)
		assert.Equal(t, item.wantPoint, point, "they should be equal")
	}
}

func TestPoint_Inside(t *testing.T) {
	tests := []Point{
		{
			X: 1, Y: 1,
		},
		{
			X: 2, Y: 1,
		},
		{
			X: 3, Y: 1,
		},
		{
			X: 4, Y: 1,
		},
		{
			X: 5, Y: 1,
		},
		{
			X: 1, Y: 2,
		},
		{
			X: 2, Y: 2,
		},
		{
			X: 3, Y: 2,
		},
		{
			X: 4, Y: 2,
		},
		{
			X: 1, Y: 3,
		},
		{
			X: 2, Y: 3,
		},
		{
			X: 3, Y: 3,
		},
		{
			X: 1, Y: 4,
		},
		{
			X: 2, Y: 4,
		},
		{
			X: 1, Y: 5,
		},
	}

	triangle := Triangle{
		Points: [3]*Point{
			&Point{
				X: 0,
				Y: 0,
			},
			&Point{
				X: 10,
				Y: 0,
			},
			&Point{
				X: 0,
				Y: 10,
			},
		},
	}

	for _, point := range tests {
		assert.Equal(t, true, triangle.IsInside(point))
	}

	assert.Equal(t, false, triangle.IsInside(Point{X: 8, Y: 8}))
}
