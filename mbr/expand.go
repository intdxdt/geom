package mbr

import "math"

//Expand to include other mbr
func (self *MBR) Expand(other MBR) *MBR {

	if other.ll[x] < self.ll[x] {
		self.ll[x] = other.ll[x]
	}
	if other.ur[x] > self.ur[x] {
		self.ur[x] = other.ur[x]
	}
	if other.ll[y] < self.ll[y] {
		self.ll[y] = other.ll[y]
	}
	if other.ur[y] > self.ur[y] {
		self.ur[y] = other.ur[y]
	}
	return self
}

//ExpandBy expands mbr by change in x and y
func (self *MBR) ExpandBy(dx, dy float64) *MBR {

	minx, miny := self.ll[x] - dx, self.ll[y] - dy
	maxx, maxy := self.ur[x] + dx, self.ur[y] + dy

	minx, maxx = math.Min(minx, maxx), math.Max(minx, maxx)
	miny, maxy = math.Min(miny, maxy), math.Max(miny, maxy)

	self.ll[x], self.ll[y] = minx, miny
	self.ur[x], self.ur[y] = maxx, maxy

	return self
}

//ExpandXY expands mbr to include x and y
func (self *MBR) ExpandXY(x_coord, y_coord float64) *MBR {

	if x_coord < self.ll[x] {
		self.ll[x] = x_coord
	}else if x_coord > self.ur[x] {
		self.ur[x] = x_coord
	}

	if y_coord < self.ll[y] {
		self.ll[y] = y_coord
	}else if y_coord > self.ur[y] {
		self.ur[y] = y_coord
	}

	return self
}

