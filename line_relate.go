package geom

import (
		"github.com/intdxdt/sset"
	"github.com/intdxdt/geom/mono"
)

//intersection of self linestring with other
func (self *LineString) linearIntersection(other *LineString) []Point {
	var ptlist []Point
	var ptset = sset.NewSSet(ptCmp)

	if self.bbox.Disjoint(&other.bbox.MBR) {
		return ptlist //disjoint
	}

	//if root mbrs intersect
	//var i, q, lnrange, ibox, qbox, qrng
	//var ok bool
	//var qrng mbr.MBR
	//var qbox, ibox *mono.MBR
	var selfsegs  []*Segment
	var othersegs []*Segment
	var lnrange   []*mono.MBR
	var inrange = self.index.Search(other.bbox.MBR)

	for i := 0; i < len(inrange); i++ {
		//cur self box
		var ibox = inrange[i]
		//search ln using ibox
		lnrange = other.index.Search(ibox.MBR)
		for q := 0; q < len(lnrange); q++ {
			var qbox = lnrange[q]
			var qrng, ok = ibox.MBR.Intersection(&qbox.MBR)

			if ok {
				self.segsInrange(&selfsegs, &qrng,   ibox.I, ibox.J)
				other.segsInrange(&othersegs, &qrng, qbox.I, qbox.J)
				self.segsegIntersection(selfsegs, othersegs, ptset)
			}
		}
	}

	var vals = ptset.Values()
	var pts = make([]Point, 0, len(vals))
	for i := range vals {
		pts = append(pts, vals[i].(Point))
	}
	return pts
}

//Checks if line intersects other line
//other{LineString} - geometry types and array as Point
func (self *LineString) intersectsLinestring(other *LineString) bool {
	var bln = false
	//if root mbrs intersect
	var othersegs []*Segment
	var selfsegs  []*Segment
	var lnrange   []*mono.MBR
	//var qrng *mbr.MBR
	//var qbox, ibox *mono.MBR
	var inrange = self.index.Search(other.bbox.MBR)

	for i := 0; !bln && i < len(inrange); i++ {
		//search ln using ibox
		var ibox = inrange[i]
		lnrange = other.index.Search(ibox.MBR)

		for q := 0; !bln && q < len(lnrange); q++ {

			var qbox = lnrange[q]
			var qrng, _ = ibox.Intersection(&qbox.MBR)

			self.segsInrange(&selfsegs, &qrng,   ibox.I, ibox.J)
			other.segsInrange(&othersegs, &qrng, qbox.I, qbox.J)

			if len(othersegs) > 0 && len(selfsegs) > 0 {
				bln = self.segseg_intersects(selfsegs, othersegs)
			}
		}
	}
	return bln
}

//line intersect polygon rings
func (self *LineString) intersects_polygon(lns []*LineString) bool {
	var bln, intersects_hole, in_hole bool
	var rings = make([]*LinearRing, 0, len(lns))
	for i := range lns {
		rings = append(rings, &LinearRing{lns[i]})
	}
	var shell = rings[0]

	bln = self.Intersects(shell.LineString)
	//if false, check if shell contains line
	if !bln {

		bln = shell.containsLine(self)
		//inside shell, does it touch hole boundary ?
		for i := 1; bln && !intersects_hole && i < len(rings); i++ {
			intersects_hole = self.Intersects(rings[i].LineString)
		}
		//inside shell but does not touch the boundary of holes
		if bln && !intersects_hole {
			//check if completely contained in hole
			for i := 1; !in_hole && i < len(rings); i++ {
				in_hole = rings[i].containsLine(self)
			}
		}
		bln = bln && !in_hole
	}
	return bln
}

// Tests whether a collection of segments from line a and line b intersects
// TODO:Improve O(n^2) - although expects few number of segs from index selection
func (self *LineString) segseg_intersects(segsa []*Segment, segsb []*Segment) bool {
	var bln = false
	var na, nb = len(segsa), len(segsb)
	for a := 0; !bln && a < na; a++ {
		for b := 0; !bln && b < nb; b++ {
			bln = segsa[a].SegSegIntersects(segsb[b])
		}
	}
	return bln
}
