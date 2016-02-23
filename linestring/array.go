package linestring

import (
    p "github.com/intdxdt/simplex/geom/point"
)

//ToArray convert as slice
func (self *LineString) ToArray() []p.Point {
    clone := make([]p.Point, len(self.coordinates))
    copy(clone, self.coordinates)
    return clone
}