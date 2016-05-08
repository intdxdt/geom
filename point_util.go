package geom

//ConvexHull computes slice of vertices as points forming convex hull
func (self *Point ) ConvexHull() []*Point {
    x, y := self[x], self[y]
    return []*Point{{x, y}, {x, y}, {x, y}, {x, y}}
}

