package algorithm

import (
	"github.com/nodejayes/gometrik/src/core"
	"math"
)

func CosineTheorem(edge1, edge2, edge3 float64, inDegrees bool) float64 {
	a2 := edge1 * edge1
	b2 := edge2 * edge2
	c2 := edge3 * edge3
	bc2 := -(2 * edge2 * edge3)
	cosAlpha := (a2 - (b2 + c2)) / bc2
	res := math.Acos(cosAlpha)
	if inDegrees {
		return core.RadiansToDegrees(res)
	}
	return res
}
