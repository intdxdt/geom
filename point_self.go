package geom

// Type - get geometry type
func (self Point) Type() GeoType {
	return GeoType(GeoTypePoint)
}

// Geometry - get geometry interface
func (self Point) Geometry() Geometry {
	return self
}
