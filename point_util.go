package geom

import (
    "math"
)

//ConvexHull computes slice of vertices as points forming convex hull
func (self *Point ) ConvexHull() []*Point {
    x, y := self[x], self[y]
    return []*Point{{x, y}, {x, y}, {x, y}, {x, y}}
}

//is null
func (self *Point) IsNull() bool {
    return math.IsNaN(self[x]) || math.IsNaN(self[y])
}