package geom

import (
	"github.com/intdxdt/side"
)

type Segment struct {
	Coords Coords
	ln     *LineString
}

////New Segment constructor
func NewSegment(coordinates Coords, i, j int) *Segment {
	coordinates.Idxs = []int{i, j}
	return &Segment{Coords: coordinates}
}

//New Segment constructor
func NewSegmentAB(a, b Point) *Segment {
	return &Segment{Coords: Coordinates([]Point{a, b})}
}

//WKT
func (self *Segment) WKT() string {
	return self.AsLineString().WKT()
}

//Segment as line string
func (self *Segment) AsLineString() *LineString {
	if self.ln == nil {
		self.ln = NewLineString(self.Coords)
	}
	return self.ln
}

//Side of pt to segement
func (self *Segment) SideOf(pt *Point) *side.Side {
	return pt.SideOf(self.A(), self.B())
}

//do two lines intersect line segments a && b with
//vertices lna0, lna1, lnb0, lnb1
func (self *Segment) SegSegIntersects(other *Segment) bool {
	return SegSegIntersects(self.A(), self.B(), other.A(), other.B())
}

//do two lines intersect line segments a && b with
//vertices lna0, lna1 and lnb0, lnb1
func (self *Segment) SegSegIntersection(other *Segment) []InterPoint {
	return SegSegIntersection(self.A(), self.B(), other.A(), other.B())
}

func (self *Segment) A() *Point{
	return self.Coords.Pt(0)
}

func (self *Segment) B() *Point{
	return self.Coords.Pt(1)
}
