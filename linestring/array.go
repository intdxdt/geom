package linestring

import (
    "github.com/intdxdt/simplex/geom/point"
)

//ToArray convert as slice
func (self *LineString) ToArray() []*point.Point {
    n := len(self.coordinates)
    clone := make([]*point.Point, n )
    for i:=0 ; i < n ; i++{
        clone[i] = self.coordinates[i].Clone()
    }
    return clone
}