package geom

//get copy of coordinates of linestring
func (self *LineString)Coordinates() []*Point {
    return CloneCoordinates(self.coordinates)
}

//checks if a point is a ring , by def every point is a ring
// which concides on itself
func (self *Point) IsRing() bool {
    return true
}

//Checks if line string is a ring
func (self *LineString) IsRing() bool {
    return IsRing(self.coordinates)
}

//Checks if polygon is a ring - default to true since all polygons are closed ring(s)
func (self *Polygon) IsRing() bool {
    return true
}


//------------------------------------------------------------------------------
//Is coordinates a ring : P0 == Pn
func IsRing(coordinates []*Point) bool {
    if len(coordinates) < 2 {
        return false
    }
    return coordinates[0].Equals(
        coordinates[len(coordinates) - 1],
    )
}

//Coordinates returns a copy of linestring coordinates
func CloneCoordinates(coordinates []*Point) []*Point {
    n := len(coordinates)
    clone := make([]*Point, n)
    for i := 0; i < n; i++ {
        clone[i] = coordinates[i].Clone()
    }
    return clone
}


//linear search if point is a member of list of points
func InCoordinates(coords []*Point, pt *Point) bool {
    bln := false
    n := len(coords)
    for i := 0; !bln && i < n; i++ {
        bln = pt.Equals(coords[i])
    }
    return bln
}
