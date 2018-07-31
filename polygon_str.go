package geom

import "bytes"

//polygon as  string
func (self *Polygon) String() string {
	var buf bytes.Buffer
	var aslines = self.AsLinear()
	var n = len(aslines) - 1
	buf.WriteString("[")
	for i, ln := range aslines {
		buf.WriteString(ln.String())
		if i < n {
			buf.WriteString(",")
		}
	}
	buf.WriteString("]")
	return buf.String()
}

//polygon as  string
func (self *Polygon) WKT() string {
	var rings = make([]Coords, 0, len(self.Holes)+1)
	rings = append(rings, self.Shell.Coordinates)
	for i := range self.Holes {
		rings = append(rings, self.Holes[i].Coordinates)
	}
	return WriteWKT(NewWKTParserObj(GeoTypePolygon, rings...))
}
