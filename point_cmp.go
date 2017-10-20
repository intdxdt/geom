package geom

import (
	"github.com/intdxdt/math"
)

//compare points as items - x | y ordering
func PointCmp(a , b interface{}) int {
	self := a.(*Point)
	other := b.(*Point)
	d := self[X] - other[X]
	if math.FloatEqual(d, 0.0) {
		d = self[Y] - other[Y]
	}

	if math.FloatEqual(d, 0.0) {
		return 0
	} else if d < 0 {
		return -1
	}
	return 1
}
