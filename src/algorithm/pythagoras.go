package algorithm

import "math"

func Pythagoras(ax, ay, bx, by float64) float64 {
	return math.Sqrt(math.Pow(ax-bx, 2) + math.Pow(ay-by, 2))
}
