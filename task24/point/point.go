package point

import "math"

// инкапсуляций достигается с помощью использования нижнего регистра для первого символа поля структуры
type Point struct {
	x, y float64
}

func (p *Point) DistanceTo(a *Point) float64 {
	return math.Sqrt(math.Pow(p.x-a.x, 2) + math.Pow(p.y-a.y, 2))
}

func NewPoint(x float64, y float64) *Point {
	return &Point{x: x, y: y}
}
