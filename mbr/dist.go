package mbr

import "math"

func distance_dxdy(self *MBR, other *MBR) (float64, float64) {
	dx := 0.0
	dy := 0.0
	//find closest edge by x
	if (self.ur[x] < other.ll[x]) {
		dx = other.ll[x] - self.ur[x]
	}else if (self.ll[x] > other.ur[x]) {
		dx = self.ll[x] - other.ur[x]
	}
	//find closest edge by y
	if (self.ur[y] < other.ll[y]) {
		dy = other.ll[y] - self.ur[y]
	}else if (self.ll[y] > other.ur[y]) {
		dy = self.ll[y] - other.ur[y]
	}
	return dx, dy
}

//Distance computes the distance between two mbrs
func (self *MBR) Distance(other MBR) float64 {

	if self.Intersects(other) {
		return 0.0
	}
	dx, dy := distance_dxdy(self, &other)
	return math.Hypot(dx, dy)
}

//DistanceSquare computes the distance squared between mbrs
func (self *MBR) DistanceSquare(other MBR) float64 {

	if self.Intersects(other) {
		return 0.0
	}
	dx, dy := distance_dxdy(self, &other)
	return (dx * dx) + (dy * dy)
}

