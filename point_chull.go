package geom

//ConvexHull computes slice of vertices as points forming convex hull
func (self Point) ConvexHull() *Polygon {
	return NewPolygon(Coordinates([]Point{self, self, self}))
}
