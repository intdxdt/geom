package geom

type Polygon struct {
	Shell *LinearRing
	Holes []*LinearRing
}

//New polygon from points
func NewPolygon(coordinates ...[]Point) *Polygon {
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
	var array = readWKT(wkt, GeoTypePolygon).ToArray()
	var pts [][]Point
	for i := range array {
		pts = append(pts, AsPointArray(array[i]))
	}
	return NewPolygon(pts...)
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
	return NewPolygon(ConvexHull(self.Shell.coordinates))
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
func shells(coords [][]Point) []*LinearRing {
	var rings = make([]*LinearRing, 0, len(coords))
	for i := range coords {
		rings = append(rings, NewLinearRing(coords[i]))
	}
	return rings
}
