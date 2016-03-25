package point
import "github.com/intdxdt/simplex/geom/wkt"

//String creates a wkt string format of point
func (self *Point) String() string {
    return wkt.Write(
        wkt.NewWKTParserObj(wkt.Point, [][2]float64{self.ToArray()}),
    )
}
