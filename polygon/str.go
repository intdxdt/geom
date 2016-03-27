package geom

import . "github.com/intdxdt/simplex/geom/wkt"


//polygon as  string
func (self *Polygon) String() string {
    //NewWKTParserObj
    self_holes := self.Holes
    rings := make([][][2]float64, len(self_holes) + 1)

    rings[0] = self.Shell.ToArray()
    for i := 0; i < len(self_holes); i++ {
        rings[i + 1] = self_holes[i].ToArray()
    }

    return WriteWKT(
        NewWKTParserObj (GeoType_Polygon, rings...),
    )
}