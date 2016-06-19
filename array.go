package geom


//ToArray converts Point to [2]float64
func (self *Point) ToArray() []float64 {
    var coords = make([]float64, len(self))
    copy(coords, self[:])
    return coords
}

//Coordinates returns a copy of linestring coordinates
func (self *LineString) ToArray() [][]float64 {
    n := len(self.coordinates)
    clone := make([][]float64, n)
    for i := 0; i < n; i++ {
        clone[i] = self.coordinates[i].ToArray()
    }
    return clone
}

//As point array
func AsPointArray(array [][2]float64) []*Point {
    var coords = make([]*Point, len(array))
    for i := range array {
        coords[i] = NewPoint(array[i][:])
    }
    return coords
}



