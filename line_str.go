package geom

import (
	"bytes"
)

//linestring as string
func (self *LineString) String() string {
	var buf bytes.Buffer
	var coords = self.coordinates
	var n = len(coords) - 1
	buf.WriteString("[")
	for i, p := range coords {
		buf.WriteString(p.String())
		if i < n {
			buf.WriteString(",")
		}
	}
	buf.WriteString("]")
	return buf.String()
}

func (self *LineString) WKT() string {
	var coords = CoordinatesAsFloat2D(self.coordinates)
	return WriteWKT(NewWKTParserObj(GeoTypeLineString, coords))
}
