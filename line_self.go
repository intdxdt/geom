package geom

import (
	"strconv"
	"github.com/intdxdt/mbr"
	"github.com/intdxdt/sset"
)

//get geometry type
func (self *LineString) Type() *geoType {
	return new_geoType(GeoType_LineString)
}

//get geometry interface
func (self *LineString) Geometry() Geometry {
	return self
}

func (self *LineString) IsComplex() bool {
	return !self.IsSimple()
}

//computes whether linestring is simple
func (self *LineString) IsSimple() bool {
	var bln = true //, bcomplx, chain, inters, jbox, qbox
	var ptset = sset.NewSSet(PointCmp)
	var cache = make(map[string]bool)
	var ln1 = make([]*Segment, 0)
	var ln2 = make([]*Segment, 0)

	var query *mbr.MBR
	var isring = self.IsRing()
	var v_i, v_j = 0, len(self.coordinates) - 1

	for i := 0; bln && i < len(self.chains); i++ {
		chain := self.chains[i]
		query = chain.BBox()
		inters := self.index.Search(query)

		for j := 0; bln && j < len(inters); j++ {
			jbox := inters[j].GetItem().(*MonoMBR)

			ckey := self.cashe_key(chain, jbox)

			if cache[ckey] || jbox.MBR.Equals(chain.MBR) {
				continue //already checked || already monotone
			}

			self.cashe_ij(cache, chain, jbox, true)
			qbox, ok := chain.MBR.Intersection(jbox.MBR)

			if !(ok) || (chain.j == jbox.i) || (chain.i == jbox.j) {
				continue //non overlapping || contiguous
			}

			if (isring && v_i == chain.i && v_j == jbox.j) ||
				(isring && v_j == chain.j && v_i == jbox.i) {
				continue //start and end vertex for closed ring
			}

			self.segs_inrange(&ln1, qbox, chain.i, chain.j)
			self.segs_inrange(&ln2, qbox, jbox.i, jbox.j)
			self.segseg_intersection(ln1, ln2, ptset, false)

			bcomplx := (chain.j != jbox.i && ptset.Size() > 0) ||
				(chain.j == jbox.i && ptset.Size() > 1) //len(ptlist) > 1))
			if bcomplx {
				bln = false
			}
		}
	}
	return bln
}

//cache box ij keys
func (self *LineString) cashe_ij(cashe map[string]bool,
	box1, box2 *MonoMBR, rev bool) {
	var a = strconv.Itoa(box1.i) + "_" + strconv.Itoa(box1.j)
	var b = strconv.Itoa(box2.i) + "_" + strconv.Itoa(box2.j)

	cashe[a+"-"+b] = true
	if rev {
		cashe[b+"-"+a] = true
	}
}

//cache key
func (self *LineString) cashe_key(box1, box2 *MonoMBR) string {
	var a = strconv.Itoa(box1.i) + "_" + strconv.Itoa(box1.j)
	var b = strconv.Itoa(box2.i) + "_" + strconv.Itoa(box2.j)

	return a + "-" + b
}

// self intersection coordinates
func (self *LineString) SelfIntersection() []*Point {

	var ptset  = sset.NewSSet(InterPointCmp)
	var tmpset = sset.NewSSet(PointCmp)
	var cache = make(map[string]bool)
	var ln1 = make([]*Segment, 0)
	var ln2 = make([]*Segment, 0)

	var ckey string
	var query *mbr.MBR
	var chain *MonoMBR

	var isring = self.IsRing()
	var vi, vj = 0, len(self.coordinates) - 1

	for i := 0; i < len(self.chains); i++ {
		chain = self.chains[i]
		query = chain.BBox()
		inters := self.index.Search(query)

		for j := 0; j < len(inters); j++ {
			jbox := inters[j].GetItem().(*MonoMBR)
			ckey = self.cashe_key(chain, jbox)

			if cache[ckey] || jbox.MBR.Equals(chain.MBR) {
				continue //already checked || already monotone
			}

			self.cashe_ij(cache, chain, jbox, true)
			qbox, ok := chain.MBR.Intersection(jbox.MBR)

			ispoint := qbox.IsPoint()
			iscontiguous := (chain.j == jbox.i) || (chain.i == jbox.j)
			is_end_contiguous := vi == chain.i && vj == jbox.j

			if !ok || (ispoint && iscontiguous) {
				continue //non overlapping || contiguous
			}

			if ispoint && isring && is_end_contiguous {
				continue //start and end vertex for closed ring
			}

			self.segs_inrange(&ln1, qbox, chain.i, chain.j)
			self.segs_inrange(&ln2, qbox, jbox.i, jbox.j)
			self.segseg_intersection(ln1, ln2, tmpset, false)

			if !tmpset.IsEmpty() {
				for _, v := range tmpset.Values() {
					intpt := &InterPoint{Pt: v.(*Point),
						I: chain.i, J: chain.j, K: jbox.i, L: jbox.j}
					ptset.Add(intpt)
				}
			}
			//overlaping contiguous mono bboxes
			if !ispoint && (iscontiguous || (isring && is_end_contiguous)) {
				var pt = &InterPoint{Pt: nil, I: chain.i, J: chain.j, K: jbox.i, L: jbox.j}
				if isring && is_end_contiguous {
					pt.Pt = self.coordinates[chain.i]
					ptset.Remove(pt)
				} else if iscontiguous {
					pt.Pt = self.coordinates[chain.j]
					ptset.Remove(pt)
				}
			}
		}
	}

	var pt_inters = sset.NewSSet(PointCmp)
	for _, p := range ptset.Values() {
		pt_inters.Add(p.(*InterPoint).Pt)
	}

	var inters = make([]*Point, 0)
	for _, p := range pt_inters.Values() {
		inters = append(inters, p.(*Point))
	}

	return inters
}
