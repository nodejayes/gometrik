package vector

import (
	"github.com/nodejayes/gometrik/src/algorithm"
)

type Triangle struct {
	a []float64
	b []float64
	c []float64
}

func NewTriangle(a []float64, b []float64, c []float64) *Triangle {
	return &Triangle{a, b, c}
}

func (ctx *Triangle) GetA() []float64 {
	return ctx.a
}

func (ctx *Triangle) GetB() []float64 {
	return ctx.b
}

func (ctx *Triangle) GetC() []float64 {
	return ctx.c
}

func (ctx *Triangle) GetEdgeCB() float64 {
	return algorithm.Pythagoras(ctx.b[0], ctx.b[1], ctx.c[0], ctx.c[1])
}

func (ctx *Triangle) GetEdgeAC() float64 {
	return algorithm.Pythagoras(ctx.c[0], ctx.c[1], ctx.a[0], ctx.a[1])
}

func (ctx *Triangle) GetEdgeAB() float64 {
	return algorithm.Pythagoras(ctx.b[0], ctx.b[1], ctx.a[0], ctx.a[1])
}

func (ctx *Triangle) GetAngleAlpha() float64 {
	return algorithm.CosineTheorem(ctx.GetEdgeCB(), ctx.GetEdgeAC(), ctx.GetEdgeAB(), true)
}

func (ctx *Triangle) GetAngleBeta() float64 {
	return algorithm.CosineTheorem(ctx.GetEdgeAC(), ctx.GetEdgeAB(), ctx.GetEdgeCB(), true)
}

func (ctx *Triangle) GetAngleGamma() float64 {
	return algorithm.CosineTheorem(ctx.GetEdgeAB(), ctx.GetEdgeAC(), ctx.GetEdgeCB(), true)
}
