package geom

import "github.com/intdxdt/mbr"

//bounding box of linestring
func (self *LineString) BBox() *mbr.MBR {
	return self.bbox.MBR
}

//bounding box of point
func (self Point) BBox() *mbr.MBR {
	var x, y = self[X], self[Y]
	return mbr.New(x, y, x, y)
}

//bounding box of segment
func (self *Segment) BBox() *mbr.MBR {
	return mbr.New(self.A[X], self.A[Y], self.B[X], self.B[Y])
}

//bounding box of linestring
func (self *Polygon) BBox() *mbr.MBR {
	return self.Shell.bbox.MBR
}
