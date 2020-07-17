package vector

type ITriangle interface {
	GetA() []float64
	GetB() []float64
	GetC() []float64
	GetEdgeCB() float64
	GetEdgeAC() float64
	GetEdgeAB() float64
	GetAngleAlpha() float64
	GetAngleBeta() float64
	GetAngleGamma() float64
}
