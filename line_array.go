package geom

//Coordinates returns a copy of linestring coordinates
func (self *LineString) ToArray() [][2]float64 {
    n := len(self.coordinates)
    clone := make([][2]float64, n)
    for i := 0; i < n; i++ {
        clone[i] = self.coordinates[i].ToArray()
    }
    return clone
}