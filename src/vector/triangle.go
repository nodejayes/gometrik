package vector

import (
	"github.com/nodejayes/gometrik/src/algorithm"
	"github.com/nodejayes/gometrik/src/core"
	"math"
)

type Triangle struct {
	a []float64
	b []float64
	c []float64
}

func NewTriangle(a []float64, b []float64, c []float64) *Triangle {
	return &Triangle{a, b, c}
}

func (ctx *Triangle) A() []float64 {
	return ctx.a
}

func (ctx *Triangle) B() []float64 {
	return ctx.b
}

func (ctx *Triangle) C() []float64 {
	return ctx.c
}

func (ctx *Triangle) EdgeCB() float64 {
	return algorithm.Pythagoras(ctx.b[0], ctx.b[1], ctx.c[0], ctx.c[1])
}

func (ctx *Triangle) EdgeAC() float64 {
	return algorithm.Pythagoras(ctx.c[0], ctx.c[1], ctx.a[0], ctx.a[1])
}

func (ctx *Triangle) EdgeAB() float64 {
	return algorithm.Pythagoras(ctx.b[0], ctx.b[1], ctx.a[0], ctx.a[1])
}

func (ctx *Triangle) AngleAlpha() float64 {
	return algorithm.CosineTheorem(ctx.EdgeCB(), ctx.EdgeAC(), ctx.EdgeAB(), true)
}

func (ctx *Triangle) AngleBeta() float64 {
	return algorithm.CosineTheorem(ctx.EdgeAC(), ctx.EdgeAB(), ctx.EdgeCB(), true)
}

func (ctx *Triangle) AngleGamma() float64 {
	return algorithm.CosineTheorem(ctx.EdgeAB(), ctx.EdgeAC(), ctx.EdgeCB(), true)
}

func (ctx *Triangle) Diameter() float64 {
	return ctx.EdgeAB() / math.Sin(ctx.AngleGamma())
}

func (ctx *Triangle) Radius() float64 {
	return ctx.Diameter() / 2
}

func (ctx *Triangle) Scope() float64 {
	return ctx.EdgeAB() + ctx.EdgeAC() + ctx.EdgeCB()
}

func (ctx *Triangle) HeightFrom(point string) float64 {
	switch point {
	case "A":
		return ctx.EdgeAB() * math.Sin(core.DegreesToRadians(ctx.AngleBeta()))
	case "B":
		return ctx.EdgeCB() * math.Sin(core.DegreesToRadians(ctx.AngleGamma()))
	case "C":
		return ctx.EdgeAC() * math.Sin(core.DegreesToRadians(ctx.AngleAlpha()))
	default:
		return 0.0
	}
}

func (ctx *Triangle) Area() float64 {
	return 0.0
}
