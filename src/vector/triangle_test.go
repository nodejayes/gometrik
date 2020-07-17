package vector

import (
	"github.com/onsi/gomega"
	"testing"
)

var A = []float64{0, 0}
var B = []float64{5, 1}
var C = []float64{3, 2}

func TestNewTriangle(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	data := NewTriangle(A, B, C)
	g.Expect(data).ToNot(gomega.BeNil())
}

func TestTriangle_GetA(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	data := NewTriangle(A, B, C)
	g.Expect(data).ToNot(gomega.BeNil())
	g.Expect(data.GetA()).To(gomega.Equal(A))
}

func TestTriangle_GetB(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	data := NewTriangle(A, B, C)
	g.Expect(data).ToNot(gomega.BeNil())
	g.Expect(data.GetB()).To(gomega.Equal(B))
}

func TestTriangle_GetC(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	data := NewTriangle(A, B, C)
	g.Expect(data).ToNot(gomega.BeNil())
	g.Expect(data.GetC()).To(gomega.Equal(C))
}

func TestTriangle_GetEdgeAB(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	data := NewTriangle(A, B, C)
	g.Expect(data).ToNot(gomega.BeNil())
	g.Expect(data.GetEdgeAB()).To(gomega.Equal(5.0990195135927845))
}

func TestTriangle_GetEdgeAC(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	data := NewTriangle(A, B, C)
	g.Expect(data).ToNot(gomega.BeNil())
	g.Expect(data.GetEdgeAC()).To(gomega.Equal(3.605551275463989))
}

func TestTriangle_GetEdgeCB(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	data := NewTriangle(A, B, C)
	g.Expect(data).ToNot(gomega.BeNil())
	g.Expect(data.GetEdgeCB()).To(gomega.Equal(2.23606797749979))
}

func TestTriangle_GetAngleAlpha(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	data := NewTriangle(A, B, C)
	g.Expect(data).ToNot(gomega.BeNil())
	g.Expect(data.GetAngleAlpha()).To(gomega.Equal(22.3801350519596))
}

func TestTriangle_GetAngleBeta(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	data := NewTriangle(A, B, C)
	g.Expect(data).ToNot(gomega.BeNil())
	g.Expect(data.GetAngleBeta()).To(gomega.Equal(37.8749836510982))
}

func TestTriangle_GetAngleGamma(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	data := NewTriangle(A, B, C)
	g.Expect(data).ToNot(gomega.BeNil())
	g.Expect(data.GetAngleGamma()).To(gomega.Equal(119.7448812969422))
}

func TestTriangle_AngleSumMustHave180Degrees(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	data := NewTriangle(A, B, C)
	g.Expect(data).ToNot(gomega.BeNil())
	g.Expect(data.GetAngleGamma() + data.GetAngleBeta() + data.GetAngleAlpha()).To(gomega.Equal(float64(180)))
}
