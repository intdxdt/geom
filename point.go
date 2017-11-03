package geom

import (
	"github.com/intdxdt/math"
)

type Point [3]float64

//New Point from x, y values
func NewPointXY(x, y float64) *Point {
	return &Point{x, y}
}

//New Point from x, y, z values
func NewPointXYZ(x, y, z float64) *Point {
	return &Point{x, y, z}
}

//New constructor of Point
func NewPoint(array []float64) *Point {
	pt := &Point{}
	n := math.MinInt(len(*pt), len(array))
	for i := 0; i < n; i++ {
		pt[i] = array[i]
	}
	return pt
}

//create a new linestring from wkt string
//empty wkt will raise an exception
func NewPointFromWKT(wkt_geom string) *Point {
	return NewPoint(ReadWKT(wkt_geom).ToArray()[0][0][:])
}

//Is point zero in 2d - origin
func (self *Point) IsZero() bool {
	return math.FloatEqual(self[X], 0.0) && math.FloatEqual(self[Y], 0.0)
}

//is null
func (self *Point) IsNull() bool {
	return math.IsNaN(self[X]) || math.IsNaN(self[Y])
}

//X gets the x coordinate of a point same as point[0]
func (self *Point) X() float64 {
	return self[X]
}

//Y gets the y coordinate of a point , same as wktreg[1]
func (self *Point) Y() float64 {
	return self[Y]
}

//Z gets the z coordinate of a point , same as wktreg[2]
func (self *Point) Z() float64 {
	return self[Z]
}

//As line strings
func (self *Point) AsLineString() *LineString {
	return NewLineString([]*Point{self, self})
}

//As line strings
func (self *Point) AsLineStrings() []*LineString {
	return []*LineString{self.AsLineString()}
}

//coordinate string
func coord_str(pt *[2]float64) string {
	return math.FloatToString(pt[X]) + " " + math.FloatToString(pt[Y])
}
