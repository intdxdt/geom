package geom

//length of line
func (self *LineString) Length() float64 {
	return self.len(0, self.Coordinates.Len()-1)
}

//length of line from index i to j
func (self *LineString) len(i, j int) float64 {
	var dist float64
	if j < i {
		i, j = j, i
	}
	for ; i < j; i++ {
		dist += self.Coordinates.Pt(i).Magnitude(
			self.Coordinates.Pt(i + 1),
		)
	}
	return dist
}
