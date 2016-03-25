package linestring
import "github.com/intdxdt/simplex/geom/wkt"

//linestring as string
func (self *LineString) String() string {
    return wkt.Write(
        wkt.NewWKTParserObj(wkt.LineString, self.ToArray()),
    )
}
