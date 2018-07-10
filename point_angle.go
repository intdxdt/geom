package geom

import (
	"github.com/intdxdt/math"
)

//Compute angle at point
func (self *Point) AngleAtPoint(a, b *Point) float64 {
	var sa, sb = a.Sub(self[X], self[Y]), b.Sub(self[X], self[Y])
	return math.Abs(math.Atan2(
		CrossProduct(sa, sb), DotProduct(sa, sb),
	))
}
