package geom

import (
	"github.com/intdxdt/math"
)

type Point [3]float64
var NullPt = Point{nan, nan, nan}

//New Point from x, y values
func Pt(x, y float64) Point {
	return Point{x, y}
}

//New Point from x, y values
func PointXY(x, y float64) Point {
	return Pt(x, y)
}

//New Point from x, y, z values
func PointXYZ(x, y, z float64) Point {
	return Point{x, y, z}
}

//New constructor of Point
func CreatePoint(array []float64) Point {
	var pt = Point{}
	var n = math.MinInt(len(pt), len(array))
	for i := 0; i < n; i++ {
		pt[i] = array[i]
	}
	return pt
}

//create a new linestring from wkt string
//empty wkt will raise an exception
func PointFromWKT(wkt_geom string) Point {
	return CreatePoint(ReadWKT(wkt_geom).ToArray()[0][0][:])
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
	return NewLineString([]Point{*self, *self})
}

//As line strings
func (self *Point) AsLineStrings() []*LineString {
	return []*LineString{self.AsLineString()}
}
