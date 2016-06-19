package geom

import (
    "strconv"
    "bytes"
)


//String creates a wkt string format of point
func (self *Point) WKT() string {
    coords := self.ToArray()
    array := [2]float64{coords[x], coords[y]}
    return WriteWKT(
        NewWKTParserObj(GeoType_Point, [][2]float64{array}),
    )
}

//String creates a wkt string format of point
func (self *Point) String() string {
    var buf  bytes.Buffer
    buf.WriteString("[")
    buf.WriteString(strconv.FormatFloat(self[x], 'f', -1, 64) + ", ")
    buf.WriteString(strconv.FormatFloat(self[y], 'f', -1, 64) + ", ")
    buf.WriteString(strconv.FormatFloat(self[z], 'f', -1, 64))
    buf.WriteString("]")

    return buf.String()
}
