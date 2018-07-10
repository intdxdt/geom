package geom

//ConvexHull computes slice of vertices as points forming convex hull
func (self *Point) ConvexHull() *Polygon {
	var pt = *self
	return NewPolygon([]Point{pt, pt, pt})
}
