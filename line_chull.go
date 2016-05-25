package geom

//ConvexHull computes slice of vertices as points forming convex hull
func (self *LineString ) ConvexHull() *Polygon {
    return NewPolygon(ConvexHull(self.coordinates))
}