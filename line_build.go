package geom

import (
	"github.com/intdxdt/rtree"
)

//builds rtree index of chains
func (self *LineString) build_index() *LineString {
	if !self.index.IsEmpty() {
		self.index.Clear()
	}
	var data = make([]rtree.BoxObj, 0, len(self.chains))
	for i := range self.chains {
		data = append(data, &self.chains[i])
	}
	self.index.Load(data) //bulkload
	return self
}
