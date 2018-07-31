package geom

import (
	"bytes"
)

//linestring as string
func (self *LineString) String() string {
	var buf bytes.Buffer
	var coords = self.Coordinates
	var n = coords.Len() - 1
	buf.WriteString("[")
	for i := range coords.Idxs {
		buf.WriteString(coords.Pt(i).String())
		if i < n {
			buf.WriteString(",")
		}
	}
	buf.WriteString("]")
	return buf.String()
}

func (self *LineString) WKT() string {
	//var coords = CoordinatesAsFloat2D()
	return WriteWKT(NewWKTParserObj(GeoTypeLineString, self.Coordinates))
}
