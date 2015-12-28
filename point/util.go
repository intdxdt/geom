package point

import "fmt"

//Convex_hull computes slice of vertices as points forming convex hull
func (self *Point ) Convex_hull(pt Point) []Point {
	return []Point{New(pt[:]), New(pt[:]), New(pt[:]), New(pt[:])}
}

//Bbox bounding box
func (self *Point) Bbox() []Point {
	return []Point{New(self[:]), New(self[:])}
}

//Wkt creates a wkt string from point
func (self *Point) Wkt() string {
	return fmt.Sprintf("POINT (%v %v)", self[x], self[y])
}


