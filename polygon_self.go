package geom
//get geometry type
func (self *Polygon) Type() *geoType{
    return new_geoType(GeoType_Polygon)
}
