package geom

//linestring as string
func (self *LineString) String() string {
    return WriteWKT(
        NewWKTParserObj(GeoType_LineString, self.ToArray()),
    )
}
