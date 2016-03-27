package geom

import (
    "math"
)

//ConvexHull computes slice of vertices as points forming convex hull
func (self *Point ) ConvexHull() []*Point {
    x, y := self[x], self[y]
    return []*Point{{x, y}, {x, y}, {x, y}, {x, y}}
}

//Bbox bounding box
func (self *Point) BBox() []*Point {
    return []*Point{
        {self[x], self[y]},
        {self[x], self[y]},
    }
}

func (self *Point) IsNull() bool {
    return math.IsNaN(self[x]) || math.IsNaN(self[y])
}