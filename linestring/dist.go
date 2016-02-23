package linestring


//Length of line
func (self *LineString) Length() float64 {
    var dist float64
    for i := range self.coordinates {
        if i == 0 {
            dist = 0
        } else {
            dist += self.coordinates[i].Distance(self.coordinates[i - 1])
        }
    }
    return dist;
}
