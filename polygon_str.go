package geom

import "bytes"

//polygon as  string
func (self *Polygon) String() string {
	//NewWKTParserObj
	var aslines = self.AsLinear()
	var buf bytes.Buffer
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
	//NewWKTParserObj
	self_holes := self.Holes
	rings := make([][][]float64, len(self_holes)+1)

	rings[0] = CoordinatesAsFloat2D(self.Shell.coordinates)
	for i := 0; i < len(self_holes); i++ {
		rings[i+1] = CoordinatesAsFloat2D(self_holes[i].coordinates)
	}

	return WriteWKT(
		NewWKTParserObj(GeoType_Polygon, rings...),
	)
}
