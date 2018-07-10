package geom

//length of line
func (self *LineString) Length() float64 {
	return self.len(0, len(self.coordinates)-1)
}

//compute length of chain
func (self *LineString) chain_length(chain *MonoMBR) float64 {
	return self.len(chain.i, chain.j)
}

//length of line from index i to j
func (self *LineString) len(i, j int) float64 {
	var dist float64
	if j < i {
		i, j = j, i
	}
	for ; i < j; i++ {
		dist += self.coordinates[i].Magnitude(&self.coordinates[i+1])
	}
	return dist
}
