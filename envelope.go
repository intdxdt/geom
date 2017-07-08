package geom

import . "simplex/geom/mbr"

//envelope of linestring
func (self *LineString) Envelope() *MBR {
	return self.bbox.MBR.Clone()
}

//bounding box of linestring
func (self *LineString) BBox() *MBR {
	return self.Envelope()
}

//envelope of point
func (self *Point) Envelope() *MBR {
	return NewMBR(self[x], self[y], self[x], self[y])
}

//bounding box of point
func (self *Point) BBox() *MBR {
	return self.Envelope()
}


//envelope of segment
func (self *Segment) Envelope() *MBR {
	return NewMBR(self.A[x], self.A[y], self.B[x], self.B[y])
}

//bounding box of segment 
func (self *Segment) BBox() *MBR {
	return self.Envelope()
}

//envelope of polygon
func (self *Polygon) Envelope() *MBR {
	return self.Shell.bbox.MBR.Clone()
}

//bounding box of linestring
func (self *Polygon) BBox() *MBR {
	return self.Envelope()
}
