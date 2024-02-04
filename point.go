package geom

import (
	"github.com/intdxdt/math"
)

type Point [3]float64

var NullPt = Point{nan, nan, nan}

// Pt - New Point from x, y values
func Pt(x, y float64) Point {
	return Point{x, y}
}

// PointXY - New Point from x, y values
func PointXY(x, y float64) Point {
	return Pt(x, y)
}

// PointXYZ - New Point from x, y, z values
func PointXYZ(x, y, z float64) Point {
	return Point{x, y, z}
}

// CreatePoint Point
func CreatePoint(array []float64) Point {
	var pt = Point{}

	var n = math.Min(len(pt), len(array))
	for i := 0; i < n; i++ {
		pt[i] = array[i]
	}
	return pt
}

// PointFromWKT - create a new linestring from wkt string
// empty wkt will raise an exception
func PointFromWKT(wkt string) Point {
	return *ReadWKT(wkt, GeoTypePoint).ToCoordinates()[0].Pt(0)
}

func PointsFromArray(array [][]float64) []Point {
	var pts = make([]Point, 0, len(array))
	for i := range array {
		pts = append(pts, CreatePoint(array[i]))
	}
	return pts
}

func PointsFromArray2D(array [][2]float64) []Point {
	var pts = make([]Point, 0, len(array))
	for i := range array {
		pts = append(pts, Pt(array[i][0], array[i][1]))
	}
	return pts
}

// Clone - clone point
func (self *Point) Clone() *Point {
	return &Point{self[X], self[Y], self[Z]}
}

// IsZero - Is point zero in 2d - origin
func (self *Point) IsZero() bool {
	return IsZero(self[X], self[Y])
}

// IsNull - is null
func (self *Point) IsNull() bool {
	return IsNull(self[X], self[Y])
}

// AsLineString - As line strings
func (self Point) AsLineString() *LineString {
	return NewLineString(Coordinates([]Point{self, self}))
}

// AsLineStrings - As line strings
func (self *Point) AsLineStrings() []*LineString {
	return []*LineString{self.AsLineString()}
}
