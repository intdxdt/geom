package mbr

import "math"

//Expand to include other mbr
func (self *MBR) ExpandIncludeMBR(other MBR) *MBR {

	if other[x1] < self[x1] {
		self[x1] = other[x1]
	}

	if other[x2] > self[x2] {
		self[x2] = other[x2]
	}

	if other[y1] < self[y1] {
		self[y1] = other[y1]
	}
    
	if other[y2] > self[y2] {
		self[y2] = other[y2]
	}
	return self
}

//ExpandBy expands mbr by change in x and y
func (self *MBR) ExpandByDelta(dx, dy float64) *MBR {

	minx, miny := self[x1] - dx, self[y1] - dy
	maxx, maxy := self[x2] + dx, self[y2] + dy

	minx, maxx = math.Min(minx, maxx), math.Max(minx, maxx)
	miny, maxy = math.Min(miny, maxy), math.Max(miny, maxy)

	self[x1], self[y1] = minx, miny
	self[x2], self[y2] = maxx, maxy

	return self
}

//ExpandXY expands mbr to include x and y
func (self *MBR) ExpandIncludeXY(x_coord, y_coord float64) *MBR {

	if x_coord < self[x1] {
		self[x1] = x_coord
	}else if x_coord > self[x2] {
		self[x2] = x_coord
	}

	if y_coord < self[y1] {
		self[y1] = y_coord
	}else if y_coord > self[y2] {
		self[y2] = y_coord
	}

	return self
}

