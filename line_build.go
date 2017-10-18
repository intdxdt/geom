package geom

import (
	"github.com/intdxdt/rtree"
)

//builds rtree index of chains
func (self *LineString) build_index() {
	if !self.index.IsEmpty() {
		self.index.Clear()
	}
	data := make([]rtree.BoxObj, len(self.chains))
	for i := range self.chains {
		data[i] = self.chains[i]
	}
	self.index.Load(data) //bulkload
}
