package geom

//polygon as  string
func (self *Polygon) String() string {
    //NewWKTParserObj
    self_holes := self.Holes
    rings := make([][][2]float64, len(self_holes) + 1)

    rings[0] = CoordinatesAsFloat2D(self.Shell.coordinates)
    for i := 0; i < len(self_holes); i++ {
        rings[i + 1] = CoordinatesAsFloat2D(self_holes[i].coordinates)
    }

    return WriteWKT(
        NewWKTParserObj (GeoType_Polygon, rings...),
    )
}