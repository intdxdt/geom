package geom

//intersection of self linestring with other
func (self *LineString) linear_intersects_rb(other *LineString) bool {
	var bln bool
	redblueLineSegmentIntersection(self, other, func(r, b int) bool {
		bln = true
		return bln
	})
	return bln
}

func (self *LineString) linear_intersection_rb(other *LineString) []Point {
	var results []Point
	var sa, sb, oa, ob *Point
	var pts []InterPoint
	redblueLineSegmentIntersection(self, other, func(i, j int) bool {
		sa, sb = self.Coordinates.Pt(i), self.Coordinates.Pt(i+1)
		oa, ob = other.Coordinates.Pt(j), other.Coordinates.Pt(j+1)
		pts = SegSegIntersection(sa, sb, oa, ob)
		for i := range pts {
			results = append(results, pts[i].Point)
		}
		return false
	})
	return results
}

//intersection of self linestring with other
func (self *LineString) linearIntersection(other *LineString) []Point {
	return self.linear_intersection_rb(other)
}

//Checks if line intersects other line
//other{LineString} - geometry types and array as Point
func (self *LineString) linearIntersects(other *LineString) bool {
	return self.linear_intersects_rb(other)
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
