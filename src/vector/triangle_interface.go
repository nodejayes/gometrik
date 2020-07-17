package vector

type ITriangle interface {
	A() []float64
	B() []float64
	C() []float64
	EdgeCB() float64
	EdgeAC() float64
	EdgeAB() float64
	AngleAlpha() float64
	AngleBeta() float64
	AngleGamma() float64
	Diameter() float64
	Radius() float64
	Scope() float64
	Area() float64
}
