package geom

import (
    . "github.com/intdxdt/simplex/geom/linestring"
    . "github.com/intdxdt/simplex/geom/point"
    "math"
)
const (
    x =iota
    y
)

type LinearRing struct {
    *LineString
}

//area of linear ring
func (self *LinearRing) Area() float64 {
    var n = self.LenVertices()
    var a, b *Point
    var area = 0.0
    b = self.VertexAt(n - 1)
    for i := 0; i < n; i++ {
        a = b
        b = self.VertexAt(i)
        area += a[y] * b[x] - a[x] * b[y]
    }
    return math.Abs(area * 0.5)
}



//new linear ring
func NewLinearRing(coordinates []*Point) *LinearRing {
    n := len(coordinates)
    var coords = make([]*Point, n)
    coords = clone_coords(coords, coordinates)
    if len(coordinates) > 1 {
        pt_0 := coordinates[0]
        pt_n := coordinates[n - 1]
        if !pt_0.Equals(pt_n) {
            coords = append(coords, pt_0.Clone())
        }
    }
    return &LinearRing{NewLineString(coords)}
}

//clone coordinates
func clone_coords(dst, src []*Point) []*Point {
    for i := range src {
        dst[i] = src[i].Clone()
    }
    return dst
}
