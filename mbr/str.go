package mbr

import (
	point "github.com/intdxdt/simplex/geom/point"
)

func (self *MBR) String() string {
	ll, ur := self.ll, self.ur
	ul, lr := point.Point{ll[x], ur[y]}, point.Point{ur[x], ll[y]}

	return "POLYGON ((" +
		ll.String() + ", " +
		ul.String() + ", " +
		ur.String() + ", " +
		lr.String() + ", " +
		ll.String() +
	"))"
}