package geom

import (
	"math"
	"simplex/cart2d"
)

//Compute angle at point
func (self *Point) AngleAtPoint(a, b *Point) float64 {
	sa,sb := a.Sub(self), b.Sub(self)
	return math.Abs(math.Atan2(
		cart2d.CrossProduct(sa, sb),
		cart2d.DotProduct(sa, sb),
	))
}
