package index

import (
	"fmt"
	"sync"
	"github.com/intdxdt/mbr"
	"github.com/intdxdt/geom/mono"
	"github.com/intdxdt/math"
)

var kobjPool = sync.Pool{
	New: func() interface{} {
		return new(KObj)
	},
}

//KObj instance struct
type KObj struct {
	node     *node
	MBR      *mbr.MBR
	IsItem   bool
	Distance float64
}

func (kobj *KObj) GetNode() *mono.MBR {
	return kobj.node.item
}

//String representation of knn object
func (kobj *KObj) String() string {
	return fmt.Sprintf("%v -> %v", kobj.node.bbox.String(), kobj.Distance)
}

//Compare - cmp interface
func kobjCmp(a interface{}, b interface{}) int {
	var self, other = a.(*KObj), b.(*KObj)
	var dx = self.Distance - other.Distance
	var r = 1
	if dx == 0 || math.Abs(dx) < math.EPSILON {
		r = 0
	} else if dx < 0 {
		r = -1
	}
	return r
}
