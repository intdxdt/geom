package geom
import . "github.com/intdxdt/simplex/geom/wkt"

//linestring as string
func (self *LineString) String() string {
    return WriteWKT(
        NewWKTParserObj(GeoType_LineString, self.ToArray()),
    )
}
