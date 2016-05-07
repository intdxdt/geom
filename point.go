package geom

import (
    . "simplex/util/math"
)

type Point [2]float64

//New Point from x, y values
func NewPointXY(x, y float64) *Point {
    return &Point{x, y}
}

//New constructor of Point
func NewPoint(array []float64) *Point {
    pt := &Point{0.0, 0.0}
    if len(array) == 1 {
        pt[x] = array[x]
    } else if len(array) >= 2 {
        pt[x], pt[y] = array[x], array[y]
    }
    return pt
}

//create a new linestring from wkt string
//empty wkt will raise an exception
func NewPointFromWKT(wkt_geom string) *Point {
    return NewPoint(
        ReadWKT(wkt_geom).ToArray()[0][0][:],
    )
}

//Is point zero - origin
func (self *Point) IsZero() bool {
    return FloatEqual(self[x], 0.0) && FloatEqual(self[y], 0.0)
}


//X gets the x coordinate of a point same as point[0]
func (self *Point) X() float64 {
    return self[x]
}

//Y gets the y coordinate of a point , same as self[1]
func (self *Point) Y() float64 {
    return self[y]
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


