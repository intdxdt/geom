package geom

import "math"

//Compute angle at point
func (self *Point) AngleAtPoint(a, b *Point) float64 {
    da, db, dab := self.Distance(a), self.Distance(b), a.Distance(b)
    // keep product units small to avoid overflow
    return math.Acos(
        ((da / db) * 0.5) +
        ((db / da) * 0.5) -
        ((dab / db) * (dab / da) * 0.5))
}
