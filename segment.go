package geom

import (
	"simplex/side"
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
func (self *Segment) SegSegIntersects(other *Segment) bool {
	return SegSegIntersects(self.A, self.B, other.A, other.B)
}

//do two lines intersect line segments a && b with
//vertices lna0, lna1 and lnb0, lnb1
func (self *Segment) SegSegIntersection(other *Segment) []*InterPoint {
	return SegSegIntersection(self.A, self.B, other.A, other.B)
}
