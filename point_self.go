package geom

//get geometry type
func (self *Point) Type() *geoType{
    return new_geoType(GeoType_Point)
}
