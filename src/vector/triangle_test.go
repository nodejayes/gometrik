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

func TestTriangle_Method_A(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	data := NewTriangle(A, B, C)
	g.Expect(data).ToNot(gomega.BeNil())
	g.Expect(data.A()).To(gomega.Equal(A))
}

func TestTriangle_Method_B(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	data := NewTriangle(A, B, C)
	g.Expect(data).ToNot(gomega.BeNil())
	g.Expect(data.B()).To(gomega.Equal(B))
}

func TestTriangle_Method_C(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	data := NewTriangle(A, B, C)
	g.Expect(data).ToNot(gomega.BeNil())
	g.Expect(data.C()).To(gomega.Equal(C))
}

func TestTriangle_Method_EdgeAB(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	data := NewTriangle(A, B, C)
	g.Expect(data).ToNot(gomega.BeNil())
	g.Expect(data.EdgeAB()).To(gomega.Equal(5.0990195135927845))
}

func TestTriangle_Method_EdgeAC(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	data := NewTriangle(A, B, C)
	g.Expect(data).ToNot(gomega.BeNil())
	g.Expect(data.EdgeAC()).To(gomega.Equal(3.605551275463989))
}

func TestTriangle_Method_EdgeCB(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	data := NewTriangle(A, B, C)
	g.Expect(data).ToNot(gomega.BeNil())
	g.Expect(data.EdgeCB()).To(gomega.Equal(2.23606797749979))
}

func TestTriangle_Method_AngleAlpha(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	data := NewTriangle(A, B, C)
	g.Expect(data).ToNot(gomega.BeNil())
	g.Expect(data.AngleAlpha()).To(gomega.Equal(22.3801350519596))
}

func TestTriangle_Method_AngleBeta(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	data := NewTriangle(A, B, C)
	g.Expect(data).ToNot(gomega.BeNil())
	g.Expect(data.AngleBeta()).To(gomega.Equal(37.8749836510982))
}

func TestTriangle_Method_AngleGamma(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	data := NewTriangle(A, B, C)
	g.Expect(data).ToNot(gomega.BeNil())
	g.Expect(data.AngleGamma()).To(gomega.Equal(119.7448812969422))
}

func TestTriangle_AngleSumMustHave180Degrees(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	data := NewTriangle(A, B, C)
	g.Expect(data).ToNot(gomega.BeNil())
	g.Expect(data.AngleGamma() + data.AngleBeta() + data.AngleAlpha()).To(gomega.Equal(float64(180)))
}

func TestTriangle_Method_Diameter(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	data := NewTriangle(A, B, C)
	g.Expect(data).ToNot(gomega.BeNil())
	g.Expect(data.Diameter()).To(gomega.Equal(14.308947884271586))
}

func TestTriangle_Method_Radius(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	data := NewTriangle(A, B, C)
	g.Expect(data).ToNot(gomega.BeNil())
	g.Expect(data.Radius()).To(gomega.Equal(data.Diameter() / 2))
}

func TestTriangle_Scope(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	data := NewTriangle(A, B, C)
	g.Expect(data).ToNot(gomega.BeNil())
	g.Expect(data.Scope()).To(gomega.Equal(10.940638766556564))
}

func TestTriangle_HeightFrom(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	data := NewTriangle(A, B, C)
	g.Expect(data).ToNot(gomega.BeNil())
	g.Expect(data.HeightFrom("A")).To(gomega.Equal(3.1304951684997047))
	g.Expect(data.HeightFrom("B")).To(gomega.Equal(1.9414506867883026))
	g.Expect(data.HeightFrom("C")).To(gomega.Equal(1.3728129459672895))
	g.Expect(data.HeightFrom("")).To(gomega.Equal(0.0))
}
