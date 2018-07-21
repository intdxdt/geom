package geom

import "github.com/intdxdt/mbr"

//bounding box of linestring
func (self *LineString) BBox() *mbr.MBR {
	return &self.bbox.MBR
}

//bounding box of point
func (self Point) BBox() *mbr.MBR {
	var x, y = self[X], self[Y]
	var box = mbr.CreateMBR(x, y, x, y)
	return  &box
}

//bounding box of segment
func (self *Segment) BBox() *mbr.MBR {
	var box = mbr.CreateMBR(self.A[X], self.A[Y], self.B[X], self.B[Y])
	return &box
}

//bounding box of linestring
func (self *Polygon) BBox() *mbr.MBR {
	return &self.Shell.bbox.MBR
}
