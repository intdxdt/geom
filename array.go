package geom

//ToCoordinates converts Point to [2]float64
func (self *Point) ToArray() []float64 {
	return (*self)[:]
}

//Coords returns a copy of linestring Coords
func (self *LineString) ToArray() [][]float64 {
	var n = self.Coordinates.Len()
	var clone = make([][]float64, 0, n)
	for i := 0; i < n; i++ {
		clone = append(clone, self.Pt(i).ToArray())
	}
	return clone
}

//As point array
func AsCoordinates(array [][]float64) Coords {
	var coords = make([]Point, 0, len(array))
	for i := range array {
		coords = append(coords, CreatePoint(array[i]))
	}
	return Coordinates(coords)
}
