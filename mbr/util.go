package mbr

func (self *MBR) AsArray() []float64 {
	return []float64{self[x1], self[y1], self[x2], self[y2]}
}
func (self *MBR) AsPolyArray() [][2]float64 {
	lx, ly := self[x1], self[y1]
	ux, uy := self[x2], self[y2]
	return [][2]float64{{lx, ly}, {lx, uy}, {ux, uy}, {ux, ly}, {lx, ly}}
}

func (self *MBR) Width() float64 {
	return self[x2] - self[x1]
}

func (self *MBR) Height() float64 {
	return self[y2] - self[y1]
}

//Area  of polygon
func (self *MBR) Area() float64 {
	return self.Height() * self.Width()
}

func (self *MBR) IsPoint() bool {
	return self.Height() == 0.0 && self.Width() == 0.0;
}

//Translate mbr  by change in x and y
func (self *MBR) Translate(dx, dy float64) *MBR {
	return NewMBR(
		self[x1]+dx, self[y1]+dy,
		self[x2]+dx, self[y2]+dy,
	)
}

func (self *MBR) Center() []float64 {
	return []float64{
		(self[x1] + self[x2]) / 2.0,
		(self[y1] + self[y2]) / 2.0,
	}
}
