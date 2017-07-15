package geom

import (
	"simplex/util/math"
)

//compare points as items - x | y ordering
func PointCmp(a , b interface{}) int {
	self := a.(*Point)
	other := b.(*Point)
	d := self[x] - other[x]
	if math.FloatEqual(d, 0.0) {
		d = self[y] - other[y]
	}

	if math.FloatEqual(d, 0.0) {
		return 0
	} else if d < 0 {
		return -1
	}
	return 1
}
