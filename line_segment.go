package geom

import (
	"github.com/intdxdt/mbr"
	"github.com/intdxdt/sset"
)

//segments in range
//xor - altenate segments if nothing is in range of box
func (self *LineString) segs_inrange(seglist *[]*Segment, box *mbr.MBR, i, j int) {
	*seglist = (*seglist)[:0]

	for ; i < j; i++ {
		inters := box.IntersectsBounds(
			self.coordinates[i][:],
			self.coordinates[i+1][:],
		)
		if inters {
			*seglist = append(*seglist, &Segment{
				A: self.coordinates[i],
				B: self.coordinates[i+1],
			})
		}
	}
}

//Segment - Segment intersection of slice of arrays
func (self *LineString) segseg_intersection(segsa, segsb []*Segment,
	ptset *sset.SSet, extend bool) {
	if !extend {
		ptset.Empty()
	}
	na, nb := len(segsa), len(segsb)
	for a := 0; a < na; a++ {
		for b := 0; b < nb; b++ {
			coord, ok := segsa[a].SegSegIntersection(segsb[b], false)
			if ok {
				for _, pt := range coord {
					ptset.Add(pt)
				}
			}
		}
	}
}
