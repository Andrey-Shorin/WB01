package main

import (
	"fmt"
	"math"
)

type Point struct {
	x, y float64
}

func NewPoint(x, y float64) Point {
	return Point{
		x: x,
		y: y,
	}
}

func dist(p1 Point, p2 Point) float64 {
	return math.Sqrt(math.Pow((p1.x-p2.x), 2) + math.Pow((p1.y-p2.y), 2))
}

func main() {
	pointOne := NewPoint(17.1, 4.4)
	pointTwo := NewPoint(8.1, 4.5)
	fmt.Println(dist(pointOne, pointTwo))
}
