package geom

import (
    . "github.com/intdxdt/simplex/geom/point"
)

//number of vertices
func (self *LineString) LenVertices() int {
    return len(self.coordinates)
}

//vertex at given index
func (self *LineString) VertexAt(i int) *Point{
    return self.coordinates[i]
}

//Coordinates returns a copy of linestring coordinates
func (self *LineString) Coordinates() []*Point {
    n := len(self.coordinates)
    clone := make([]*Point, n)
    for i := 0; i < n; i++ {
        clone[i] = self.coordinates[i].Clone()
    }
    return clone
}
