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
func NewPolygonFromWKT(wkt_geom string) *Polygon {
	var array = readWKT(wkt_geom, GeoTypePolygon).ToArray()
	var pts [][]Point
	for i := range array {
		pts = append(pts, AsPointArray(array[i]))
	}
	return NewPolygon(pts...)
}

//As line strings
func (self *Polygon) AsLinearRings() []*LinearRing {
	var rings = make([]*LinearRing, len(self.Holes)+1)
	rings[0] = self.Shell
	for i := 0; i < len(self.Holes); i++ {
		rings[i+1] = self.Holes[i]
	}
	return rings
}

//polygon shells
func shells(coords [][]Point) []*LinearRing {
	var n = len(coords)
	var rings = make([]*LinearRing, n)
	for i := 0; i < n; i++ {
		rings[i] = NewLinearRing(coords[i])
	}
	return rings
}
