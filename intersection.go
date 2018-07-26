package geom

import "github.com/intdxdt/sset"

//Checks if pt intersection other geometry
func (pt Point) Intersection(other Geometry) []Point {
	var res []Point
	//checks for non-geometry types
	if IsNullGeometry(other) {
		return res
	}

	other.Type().IsPoint()

	if other.Type().IsPoint() {
		var p = CastAsPoint(other)
		if pt.Equals2D(&p) {
			res = append(res, pt)
		}
	} else if other.Type().IsLineString() {
		res = pt.AsLineString().Intersection(other.(*LineString))
	} else if other.Type().IsSegment() {
		res = pt.AsLineString().Intersection(other.(*Segment))
	} else if other.Type().IsPolygon() {
		res = pt.AsLineString().Intersection(other.(*Polygon))
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

	if other.Type().IsPoint() {
		var pt = CastAsPoint(other)
		res = self.linearIntersection(pt.AsLineString())
	} else if other.Type().IsSegment() {
		res = self.linearIntersection(CastAsSegment(other).AsLineString())
	} else if other.Type().IsLineString() {
		res = self.linearIntersection(CastAsLineString(other))
	} else if other.Type().IsPolygon() {
		res = self.intersectionPolygonRings(CastAsPolygon(other).AsLinearRings())
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

	if other.Type().IsPoint() {
		res = CastAsPoint(other).AsLineString().Intersection(self)
	} else if other.Type().IsSegment() {
		res = other.(*Segment).Intersection(self)
	} else if other.Type().IsLineString() {
		res = other.(*LineString).Intersection(self)
	} else if other.Type().IsPolygon() {
		var ptset = sset.NewSSet(ptCmp)
		//other intersect self
		var lns = other.(*Polygon).AsLinear()
		for _, ln := range lns {
			var pts = ln.Intersection(self)
			for i := range pts {
				ptset.Add(pts[i])
			}
		}

		//self intersects other
		lns = self.AsLinear()
		for _, ln := range lns {
			var pts = ln.Intersection(other)
			for i := range pts {
				ptset.Add(pts[i])
			}
		}

		var pts = ptset.Values()
		for i := range pts {
			res = append(res, pts[i].(Point))
		}
	}

	return res
}

//line intersect polygon rings
func (self *LineString) intersectionPolygonRings(rings []*LinearRing) []Point {
	var shell = rings[0]
	var ptset = sset.NewSSet(ptCmp)
	var res []Point
	var bln = self.bbox.MBR.Intersects(&shell.bbox.MBR)

	if bln {
		spts := self.linearIntersection(shell.LineString)
		for idx := range spts {
			ptset.Add(spts[idx])
		}
		//inside shell, does it touch hole boundary ?
		for i := 1; i < len(rings); i++ {
			hpts := self.linearIntersection(rings[i].LineString)
			for idx := range hpts {
				ptset.Add(hpts[idx])
			}
		}
		//check for all vertices
		for idx := range self.coordinates {
			var pt = self.coordinates[idx]
			if shell.containsPoint(&pt) {
				inhole := false
				for i := 1; !inhole && i < len(rings); i++ {
					inhole = rings[i].containsPoint(&pt)
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
