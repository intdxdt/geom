package geom

//ConvexHull computes slice of vertices as points forming convex hull
func (self *Polygon ) ConvexHull() *Polygon {
    return NewPolygon(ConvexHull(self.Shell.coordinates))
}
