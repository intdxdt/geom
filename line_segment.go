package geom

import (
	"github.com/intdxdt/sset"
)

//segments in range
//xor - altenate segments if nothing is in range of box
func (self *LineString) segsInrange(seglist *[]int, minx, miny, maxx, maxy float64, i, j int) {
	*seglist = (*seglist)[:0]
	var a, b *Point
	for ; i < j; i++ {
		a, b = self.Coordinates.Pt(i), self.Coordinates.Pt(i+1)
		if intersectsBounds(minx, miny, maxx, maxy, a, b) {
			*seglist = append(*seglist, self.Coordinates.Idxs[i], self.Coordinates.Idxs[i+1])
		}
	}
}

//Segment - Segment intersection of slice of arrays
func (self *LineString) segsegIntersection(a_coords []Point, segsa []int, b_coords []Point, segsb []int, ptset *sset.SSet) {
	var na, nb = len(segsa), len(segsb)
	var a0, a1 *Point
	var coord []InterPoint
	for a := 0; a < na; a += 2 {
		a0, a1 = &a_coords[segsa[a]], &a_coords[segsa[a+1]]
		for b := 0; b < nb; b += 2 {
			coord = SegSegIntersection(a0, a1, &b_coords[segsb[b]], &b_coords[segsb[b+1]])
			for idx := range coord {
				ptset.Add(coord[idx].Point)
			}
		}
	}
}
