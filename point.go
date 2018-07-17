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
func PointFromWKT(wkt string) Point {
	return CreatePoint(readWKT(wkt, GeoTypePoint).ToArray()[0][0][:])
}

//Is point zero in 2d - origin
func (self *Point) IsZero() bool {
	return IsZero(self[X], self[Y])
}

//is null
func (self *Point) IsNull() bool {
	return IsNull(self[X], self[Y])
}

//As line strings
func (self Point) AsLineString() *LineString {
	return NewLineString([]Point{self, self})
}

//As line strings
func (self *Point) AsLineStrings() []*LineString {
	return []*LineString{self.AsLineString()}
}
