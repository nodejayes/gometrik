package core

import (
	"github.com/onsi/gomega"
	"testing"
)

func TestDegreesToRadians(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	g.Expect(DegreesToRadians(12.0)).To(gomega.Equal(0.20943951023931953))
}

func TestRadiansToDegrees(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	g.Expect(RadiansToDegrees(12.0)).To(gomega.Equal(327.5493541569879))
}
