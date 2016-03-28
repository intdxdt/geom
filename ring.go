package geom

import (
    "math"
)


type LinearRing struct {
    *LineString
}

//new linear ring
func NewLinearRing(coordinates []*Point) *LinearRing {
    coords := CloneCoordinates(coordinates)
    if len(coordinates) > 1 {
        if !IsRing(coordinates) {
            coords = append(coords, coordinates[0].Clone())
        }
    }
    return &LinearRing{NewLineString(coords)}
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

