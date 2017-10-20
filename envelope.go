package geom

import "github.com/intdxdt/mbr"

//envelope of linestring
func (self *LineString) Envelope() *mbr.MBR {
	return self.bbox.MBR.Clone()
}

//bounding box of linestring
func (self *LineString) BBox() *mbr.MBR {
	return self.Envelope()
}

//envelope of point
func (self *Point) Envelope() *mbr.MBR {
	return mbr.NewMBR(self[X], self[Y], self[X], self[Y])
}

//bounding box of point
func (self *Point) BBox() *mbr.MBR {
	return self.Envelope()
}


//envelope of segment
func (self *Segment) Envelope() *mbr.MBR {
	return mbr.NewMBR(self.A[X], self.A[Y], self.B[X], self.B[Y])
}

//bounding box of segment 
func (self *Segment) BBox() *mbr.MBR {
	return self.Envelope()
}

//envelope of polygon
func (self *Polygon) Envelope() *mbr.MBR {
	return self.Shell.bbox.MBR.Clone()
}

//bounding box of linestring
func (self *Polygon) BBox() *mbr.MBR {
	return self.Envelope()
}
