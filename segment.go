package geom

import (
	"simplex/side"
	"github.com/intdxdt/math"
	"github.com/intdxdt/segs"
)

type Segment struct {
	A  *Point
	B  *Point
	ln *LineString
}

//New Segment constructor
func NewSegment(a, b *Point) *Segment {
	return &Segment{A: a, B: b}
}

//WKT
func (self *Segment) WKT() string {
	return self.AsLineString().WKT()
}

//Segment as line string
func (self *Segment) AsLineString() *LineString {
	if self.ln == nil {
		self.ln = NewLineString([]*Point{self.A, self.B})
	}
	return self.ln
}

//Side of pt to segement
func (self *Segment) SideOf(pt *Point) *side.Side {
	return pt.SideOf(self.A, self.B)
}

//do two lines intersect line segments a && b with
//vertices lna0, lna1, lnb0, lnb1
func (self *Segment) SegSegIntersects(other *Segment, extln bool) bool {
	return segs.Intersects(
		self.A[:], self.B[:], other.A[:], other.B[:],
	)
}

//do two lines intersect line segments a && b with
//vertices lna0, lna1 and lnb0, lnb1
func (self *Segment) SegSegIntersection(other *Segment, extln bool) ([]*Point, bool) {
	var coords = make([]*Point, 0)
	var pts = segs.Intersection(
		self.A[:], self.B[:], other.A[:], other.B[:],
	)
	for _, pt := range pts {
		coords = append(coords, NewPoint(pt))
	}
	return coords, len(coords) > 0
}

//clamp to zero or one
func snap_to_zero_or_one(v float64) float64 {
	if math.FloatEqual(v, 0.0) {
		v = 0.0
	} else if math.FloatEqual(v, 1.0) {
		v = 1.0
	}
	return v
}
