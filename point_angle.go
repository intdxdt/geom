package geom

import (
	"github.com/intdxdt/math"
)

//Compute angle at point
func (self *Point) AngleAtPoint(a, b *Point) float64 {
	var ax, ay = a.Sub(self[X], self[Y])
	var bx, by = b.Sub(self[X], self[Y])
	return math.Abs(math.Atan2(
		CrossProduct(ax, ay, bx, by), DotProduct(ax, ay, bx, by),
	))
}
