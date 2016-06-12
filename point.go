package geom

import (
    . "simplex/util/math"
    "math"
)

type Point [3]float64

//New Point from x, y values
func NewPointXY(x, y float64) *Point {
    return &Point{x, y}
}

//New Point from x, y values
func NewPointXYZ(x, y, z float64) *Point {
    return &Point{x, y, z}
}

//New constructor of Point
func NewPoint(array []float64) *Point {
    pt := &Point{}
    if len(array) == 1 {
        pt[x] = array[x]
    } else if len(array) >= 2 {
        pt[x], pt[y] = array[x], array[y]
    } else if len(array) >= 3 {
        pt[z] = array[z]
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
    return FloatEqual(self[x], 0.0) && FloatEqual(self[y], 0.0)
}

//is null
func (self *Point) IsNull() bool {
    return math.IsNaN(self[x]) || math.IsNaN(self[y])
}

//X gets the x coordinate of a point same as point[0]
func (self *Point) X() float64 {
    return self[x]
}

//Y gets the y coordinate of a point , same as self[1]
func (self *Point) Y() float64 {
    return self[y]
}

//Z gets the z coordinate of a point , same as self[2]
func (self *Point) Z() float64 {
    return self[z]
}

//As line strings
func (self *Point) AsLineString() *LineString {
    return NewLineString([]*Point{self.Clone(), self.Clone()})
}

//As line strings
func (self *Point) AsLineStrings() []*LineString {
    return []*LineString{self.AsLineString()}
}

//coordinate string
func coord_str(pt *[2]float64) string {
    return FloatToString(pt[x]) + " " + FloatToString(pt[y])
}


