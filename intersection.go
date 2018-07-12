package geom

//Checks if pt intersection other geometry
func (pt *Point) Intersection(other Geometry) []Point {
	var res []Point
	//checks for non-geometry types
	if IsNullGeometry(other) {
		return res
	}

	if p, ok := IsPoint(other); ok {
		if pt.Equals2D(p) {
			res = append(res, *pt)
		}
	} else if ln, ok := IsLineString(other); ok {
		res = pt.AsLineString().Intersection(ln)
	} else if seg, ok := IsSegment(other); ok {
		res = pt.AsLineString().Intersection(seg)
	} else if ply, ok := IsPolygon(other); ok {
		res = pt.AsLineString().Intersection(ply)
	}
	return res
}

//Segment intersection other geometry
func (self *Segment) Intersection(other Geometry) []Point {
	return self.AsLineString().Intersection(other)
}

//Checks if pt intersection other geometry
func (self *LineString) Intersection(other Geometry) []Point {
	var res []Point
	//checks for non-geometry types
	if IsNullGeometry(other) {
		return res
	}

	if pt, ok := IsPoint(other); ok {
		res = self.linear_intersection(pt.AsLineString())
	} else if seg, ok := IsSegment(other); ok {
		res = self.linear_intersection(seg.AsLineString())
	} else if ln, ok := IsLineString(other); ok {
		res = self.linear_intersection(ln)
	} else if ply, ok := IsPolygon(other); ok {
		res = self.intersection_polygon_rings(ply.AsLinearRings())
	}

	return res
}

//Checks if pt intersection other geometry
func (self *Polygon) Intersection(other Geometry) []Point {
	var res []Point
	//checks for non-geometry types
	if IsNullGeometry(other) {
		return res
	}

	if pt, ok := IsPoint(other); ok {
		ln := pt.AsLineString()
		res = ln.Intersection(self)
	} else if seg, ok := IsSegment(other); ok {
		res = seg.Intersection(self)
	} else if ln, ok := IsLineString(other); ok {
		res = ln.Intersection(self)
	} else if ply, ok := IsPolygon(other); ok {
		ptset := NewPtSet()
		//other intersect self
		lns := ply.AsLinear()
		for _, ln := range lns {
			pts := ln.Intersection(self)
			for _, p := range pts {
				ptset.Add(p)
			}
		}

		//self intersects other
		lns = self.AsLinear()
		for _, ln := range lns {
			pts := ln.Intersection(other)
			for _, p := range pts {
				ptset.Add(p)
			}
		}

		pts := ptset.Values()
		for _, p := range pts {
			res = append(res, p.(Point))
		}
	}

	return res
}

//line intersect polygon rings
func (self *LineString) intersection_polygon_rings(rings []*LinearRing) []Point {
	var shell = rings[0]
	var ptset = NewPtSet()

	var res []Point
	var selfBox = self.BBox()
	var shellBox = shell.BBox()
	var bln = selfBox.Intersects(&shellBox)

	if bln {
		spts := self.linear_intersection(shell.LineString)
		for idx := range spts {
			ptset.Add(spts[idx])
		}
		//inside shell, does it touch hole boundary ?
		for i := 1; i < len(rings); i++ {
			hpts := self.linear_intersection(rings[i].LineString)
			for idx := range hpts {
				ptset.Add(hpts[idx])
			}
		}
		//check for all vertices
		for idx := range self.coordinates {
			var pt = self.coordinates[idx]
			if shell.contains_point(&pt) {
				inhole := false
				for i := 1; !inhole && i < len(rings); i++ {
					inhole = rings[i].contains_point(&pt)
				}
				if !inhole {
					ptset.Add(pt)
				}
			}
		}
		vals := ptset.Values()
		for _, pt := range vals {
			res = append(res, pt.(Point))
		}
	}
	return res
}
