package geom

import (
    . "github.com/intdxdt/simplex/util/math"
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
    return self.Equals(&Point{0, 0})
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
func (self *Point) AsLineStrings() []*LineString {
    var coords = []*Point{self.Clone(), self.Clone()}
    var sh = NewLineString(coords)
    var rings = []*LineString{sh}
    return rings
}

//coordinate string
func coord_str(pt *[2]float64) string {
    return FloatToString(pt[x]) + " " + FloatToString(pt[y])
}


