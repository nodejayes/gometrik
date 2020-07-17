package algorithm

import (
	"github.com/nodejayes/gometrik/src/core"
	"math"
)

func CosineTheorem(edge1, edge2, edge3 float64, inDegrees bool) float64 {
	a2 := math.Pow(edge1, 2)
	b2 := math.Pow(edge2, 2)
	c2 := math.Pow(edge3, 2)
	bc2 := -(2 * edge2 * edge3)
	cosAlpha := (a2 - (b2 + c2)) / bc2
	res := math.Acos(cosAlpha)
	if inDegrees {
		return core.RadiansToDegrees(res)
	}
	return res
}
