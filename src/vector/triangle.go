package vector

import "math"

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
	return ctx.pythagoras(ctx.b[0], ctx.b[1], ctx.c[0], ctx.c[1])
}

func (ctx *Triangle) GetEdgeAC() float64 {
	return ctx.pythagoras(ctx.c[0], ctx.c[1], ctx.a[0], ctx.a[1])
}

func (ctx *Triangle) GetEdgeAB() float64 {
	return ctx.pythagoras(ctx.b[0], ctx.b[1], ctx.a[0], ctx.a[1])
}

func (ctx *Triangle) pythagoras(ax, ay, bx, by float64) float64 {
	return math.Sqrt(math.Pow(ax-bx, 2) + math.Pow(ay-by, 2))
}
