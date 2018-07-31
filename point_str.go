package geom

import (
	"bytes"
	"strconv"
	"github.com/intdxdt/math"
)

//String creates a wkt string format of point
func (self Point) WKT() string {
	return WriteWKT(NewWKTParserObj(GeoTypePoint, Coordinates([]Point{self})))
}

//String creates a wkt string format of point
func (self *Point) String() string {
	var buf bytes.Buffer
	buf.WriteString("[")
	buf.WriteString(strconv.FormatFloat(self[X], 'f', -1, 64) + ", ")
	buf.WriteString(strconv.FormatFloat(self[Y], 'f', -1, 64) + ", ")
	buf.WriteString(strconv.FormatFloat(self[Z], 'f', -1, 64))
	buf.WriteString("]")

	return buf.String()
}

//coordinate string
func coordStr(pt []float64) string {
	return math.FloatToString(pt[X]) + " " + math.FloatToString(pt[Y])
}
