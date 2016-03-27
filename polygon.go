package geom


type Polygon struct {
    Shell *LinearRing
    Holes []*LinearRing
}

// description Polygon geometry
// param [opts]{Object} - {monosize : 1, bucketsize}
func NewPolygon(coordinates ...[]*Point) *Polygon {
    var shell *LinearRing
    var rings = shells(coordinates)
    var holes = make([]*LinearRing, 0)
    shell = rings[0]
    if len(rings) > 1 {
        holes = rings[1:]
    }
    return &Polygon{shell, holes}
}

//As line strings
func (self *Polygon) AsLineStrings() []*LineString{
    var n = len(self.Holes) + 1
    var sh = self.Shell.LineString
    var rings = make([]*LineString, n)
    rings[0] = sh
    for i := 1; i < n; i++ {
        rings[i] = self.Holes[i].LineString
    }
    return rings
}

//polygon shells
func shells(coords [][]*Point) []*LinearRing {
    var n = len(coords)
    var rings = make([]*LinearRing, n)
    for i := 0; i < n; i++ {
        rings[i] = NewLinearRing(coords[i])
    }
    return rings
}



