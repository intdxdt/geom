package geom

import (
	"github.com/intdxdt/side"
)

type VBits uint8

const (
	OtherB VBits = 1 << iota // 1 << 0 == 0001
	OtherA                   // 1 << 1 == 0010
	SelfB                    // 1 << 2 == 0100
	SelfA                    // 1 << 3 == 1000
)
const InterX VBits = 0

const (
	SelfMask  = SelfA | SelfB
	OtherMask = OtherA | OtherB
)


type Segment struct {
	Coords Coords
	ln     *LineString
}

//New Segment constructor
func NewSegment(coordinates Coords, i, j int) *Segment {
	coordinates.Idxs = []int{coordinates.Idxs[i], coordinates.Idxs[j]}
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

func (self *Segment) A() *Point {
	return self.Coords.Pt(0)
}

func (self *Segment) B() *Point {
	return self.Coords.Pt(1)
}
