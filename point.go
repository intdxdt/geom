package geom

type Point [2]float64

//New Point from x, y values
func NewPointXY(x, y float64) *Point {
    return &Point{x, y}
}

//New constructor of Point
func NewPoint(array []float64) *Point {
    pt := &Point{0.0, 0.0}
    if len(array) == 1 {
        pt[x] = array[x]
    } else if len(array) >= 2 {
        pt[x], pt[y] = array[x], array[y]
    }
    return pt
}

//Is point zero - origin
func (self *Point) IsZero() bool {
    return self.Equals(&Point{0, 0})
}

//Clone point
func (self *Point) Clone() *Point {
    return NewPointXY(self[x], self[y])
}

//X gets the x coordinate of a point same as point[0]
func (self *Point) X() float64 {
    return self[x]
}

//Y gets the y coordinate of a point , same as self[1]
func (self *Point) Y() float64 {
    return self[y]
}

//ToArray converts Point to [2]float64
func (self *Point) ToArray() [2]float64 {
    return [2]float64{self[x], self[y]}
}

//As line strings
func (self *Point) AsLineStrings() []*LineString {
    var coords = []*Point{self.Clone(), self.Clone()}
    var sh = NewLineString(coords)
    var rings = []*LineString{sh}
    return rings
}

