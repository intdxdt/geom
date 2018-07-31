package geom

type Polygon struct {
	Shell *LinearRing
	Holes []*LinearRing
}

//New polygon from points
func NewPolygon(coordinates ...[]Point) *Polygon {
	var coords = make([]Coords, 0, len(coordinates))
	for i := range coordinates {
		coords = append(coords, Coordinates(coordinates[i]))
	}
	var rings = shells(coords)
	return NewPolygonFromRings(rings...)
}

//New polygon from points
func NewPolygonFromCoords(coordinates ...Coords) *Polygon {
	var rings = shells(coordinates)
	return NewPolygonFromRings(rings...)
}

//New Polygon from rings
func NewPolygonFromRings(rings ...*LinearRing) *Polygon {
	var shell *LinearRing
	var holes = make([]*LinearRing, 0)
	shell = rings[0]
	if len(rings) > 1 {
		holes = rings[1:]
	}
	return &Polygon{Shell: shell, Holes: holes}
}

//create a new linestring from wkt string
//empty wkt are not allowed
func NewPolygonFromWKT(wkt string) *Polygon {
	var coords = readWKT(wkt, GeoTypePolygon).ToCoordinates()
	return NewPolygonFromCoords(coords...)
}

//get geometry type
func (self *Polygon) Type() GeoType {
	return GeoType(GeoTypePolygon)
}

//get geometry interface
func (self *Polygon) Geometry() Geometry {
	return self
}

//ConvexHull computes slice of vertices as points forming convex hull
func (self *Polygon) ConvexHull() *Polygon {
	return NewPolygonFromCoords(ConvexHull(self.Shell.Coordinates))
}

//As line strings
func (self *Polygon) AsLinearRings() []*LinearRing {
	var rings = make([]*LinearRing, 0, len(self.Holes)+1)
	rings = append(rings, self.Shell)
	for i := range self.Holes {
		rings = append(rings, self.Holes[i])
	}
	return rings
}

//polygon shells
func shells(coords []Coords) []*LinearRing {
	var rings = make([]*LinearRing, 0, len(coords))
	for i := range coords {
		rings = append(rings, NewLinearRing(coords[i]))
	}
	return rings
}
