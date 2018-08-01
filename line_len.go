package geom

import "github.com/intdxdt/geom/mono"

//length of line
func (self *LineString) Length() float64 {
	return self.len(0, self.Coordinates.Len()-1)
}

//compute length of chain
func (self *LineString) chainLength(chain *mono.MBR) float64 {
	return self.len(chain.I, chain.J)
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
