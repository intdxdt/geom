package point

import 	"math"


//Distance computes distance between two points
func (self *Point ) Distance(pt Point) float64 {
	return math.Hypot(self[x] - pt[x], self[y] - pt[y])
}
