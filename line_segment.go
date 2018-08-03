package geom

import (
	"github.com/intdxdt/mbr"
	"github.com/intdxdt/sset"
)

//segments in range
//xor - altenate segments if nothing is in range of box
func (self *LineString) segsInrange(seglist *[]*Segment, box *mbr.MBR, i, j int) {
	*seglist = (*seglist)[:0]
	var a, b *Point
	for ; i < j; i++ {
		a, b = self.Coordinates.Pt(i), self.Coordinates.Pt(i+1)
		if box.IntersectsBounds(a[:], b[:]) {
			*seglist = append(*seglist, NewSegment(self.Coordinates, self.Coordinates.Idxs[i], self.Coordinates.Idxs[i+1]))
		}
	}
}

//Segment - Segment intersection of slice of arrays
func (self *LineString) segsegIntersection(segsa, segsb []*Segment, ptset *sset.SSet) {
	var na, nb = len(segsa), len(segsb)
	for a := 0; a < na; a++ {
		for b := 0; b < nb; b++ {
			var coord = segsa[a].SegSegIntersection(segsb[b])
			if len(coord) > 0 {
				for _, pt := range coord {
					ptset.Add(pt.Point)
				}
			}
		}
	}

}
