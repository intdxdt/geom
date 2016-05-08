package geom

import "sort"

//coordinates interable of points
type Coordinates []*Point

//Len for sort interface
func (o Coordinates) Len() int {
    return len(o)
}

//Swap for sort interface
func (o Coordinates) Swap(i, j int) {
    o[i], o[j] = o[j], o[i]
}

//pop chain from chainl list
func (self Coordinates) Pop() (*Point, Coordinates) {
    var v *Point
    var n int
    if len(self) == 0 {
        return nil, self
    }
    n = len(self) - 1
    v, self[n] = self[n], nil
    return v, self[:n]
}


//XBoxes type  for  x sorting of boxes
type XYCoordinates struct{
    Coordinates
}

//lexical sort of x and y coordinates
func (o XYCoordinates) Less(i, j int) bool {
    if o.Coordinates[i][x] < o.Coordinates[j][x] {
        return true
    } else if o.Coordinates[i][x] == o.Coordinates[j][x] {
        if o.Coordinates[i][y] < o.Coordinates[j][y]{
            return true
        }
    }
    return false
}

//Inplace Lexicographic sort
func (self XYCoordinates ) Sort(){
    sort.Sort(self)
}

//XCoordinates type  for  x sorting of boxes
type XCoordinates struct{
    Coordinates
}

//Less sorts boxes by y
func (self  XCoordinates) Less(i, j int) bool {
    return self.Coordinates[i][x] < self.Coordinates[j][x]

}

//Inplace sort by x
func (self XCoordinates ) Sort(){
    sort.Sort(self)
}



//YCoordinates type for y sorting of boxes
type YCoordinates struct{
    Coordinates
}


//Less sorts boxes by y
func (o *YCoordinates) Less(i, j int) bool {
    return o.Coordinates[i][y] < o.Coordinates[j][y]

}

//Inplace sort by y
func (self *YCoordinates ) Sort(){
    sort.Sort(self)
}


//get copy of coordinates of linestring
func (self *LineString) Coordinates() []*Point {
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



