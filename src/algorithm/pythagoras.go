package algorithm

import "math"

func Pythagoras(ax, ay, bx, by float64) float64 {
	return PythagorasLength(ax-bx, ay-by)
}

func PythagorasLength(l1, l2 float64) float64 {
	return math.Sqrt((l1 * l1) + (l2 * l2))
}
