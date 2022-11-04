package triangleslib

import (
	"math"
	"sort"
)

// IsTriangle можно ли построить треугольник по координатам?
func IsTriangle(x1, y1, x2, y2, x3, y3 int) bool {
	return !(x3*(y2-y1)-y3*(x2-x1) == x1*y2-x2*y1)
}

// Square Натйи площадь треугольника
func Square(x1, y1, x2, y2, x3, y3 int) float64 {
	a := calcLength(x1, y1, x2, y2)
	b := calcLength(x2, y2, x3, y3)
	c := calcLength(x3, y3, x1, y1)

	p := (a + b + c) / 2
	s := math.Sqrt(p * (p - a) * (p - b) * (p - c))
	return s
}

// IsRightTriangle Является ли треугольник прямоугольным
func IsRightTriangle(x1, y1, x2, y2, x3, y3 int) bool {
	a := calcLength(x1, y1, x2, y2)
	b := calcLength(x2, y2, x3, y3)
	c := calcLength(x3, y3, x1, y1)
	var sides []float64
	sides = make([]float64, 3)
	sides[0] = a
	sides[1] = b
	sides[2] = c
	sort.Float64s(sides)

	if math.Round(math.Pow(sides[2], 2)) == math.Round((math.Pow(sides[0], 2) + math.Pow(sides[1], 2))) {
		return true
	} else {
		return false
	}
}

// Вычислить расстояние между точками
func calcLength(x1, y1, x2, y2 int) float64 {
	xdiff := float64(x2 - x1)
	ydiff := float64(y2 - y1)
	return math.Sqrt(math.Pow(xdiff, 2) + math.Pow(ydiff, 2))
}
