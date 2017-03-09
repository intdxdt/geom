package geom

//number of vertices
func (self *LineString) LenVertices() int {
	return len(self.coordinates)
}

//vertex at given index
func (self *LineString) VertexAt(i int) *Point {
	return self.coordinates[i]
}
