package main

import (
	"fmt"
	"math"
	"math/rand"
)

type Point struct{ x, y float64 }

func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.x-p.x, q.y-p.y)
}

type Path []Point

func (path Path) Distance() float64 {
	sum := 0.0
	for i := range path {
		if i > 0 {
			sum += path[i-1].Distance(path[i])
		}
	}
	return sum
}

func main() {
	var sides int
	fmt.Scan(&sides)

	var max int = 100.0
	var min int = -100.0
	var ope = max - min
	var p Path

	for i := 0; i < sides; i++ {
		var point = Point{float64(rand.Intn(ope) + min), float64(rand.Intn(ope) + min)}
		p = append(p, point)
	}

	fmt.Printf("Generating a [%v] sides figure\n", sides)
	fmt.Println("Figure's vertices")
	for i := 0; i < sides; i++ {
		fmt.Printf("( %v, %v)\n", p[i].x, p[i].y)
	}
	fmt.Println(p.Distance())
}

