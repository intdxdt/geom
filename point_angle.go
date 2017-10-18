package geom

import (
	"github.com/intdxdt/cart"
	"github.com/intdxdt/math"
)

//Compute angle at point
func (self *Point) AngleAtPoint(a, b *Point) float64 {
	sa, sb := a.Sub(self), b.Sub(self)
	return math.Abs(math.Atan2(
		cart.CrossProduct(sa, sb),
		cart.DotProduct(sa, sb),
	))
}
