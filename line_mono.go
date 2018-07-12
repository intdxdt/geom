package geom

import (
	"github.com/intdxdt/mbr"
	"github.com/intdxdt/math"
)

type MonoMBR struct {
	mbr.MBR
	i int
	j int
}

//clone  mono mbr
func (box *MonoMBR) BBox() mbr.MBR {
	return box.MBR
}

//update mono chain index
func (box *MonoMBR) update_index(i, j int) {
	box.i = i
	box.j = j
}

//new monotone mbr
func new_mono_mbr(box mbr.MBR) MonoMBR {
	return MonoMBR{box, null, null}
}

//build xymonotone chain, perimeter length,
//monotone build starts from i and ends at j, designed for
//appending new points to the end of line
func (self *LineString) process_chains(i, j int) *LineString {
	var dx, dy float64
	var v0, v1 *Point
	var cur_x, cur_y, prev_x, prev_y int
	var mono_limit = self.monosize

	prev_x, prev_y = null, null

	if j == 0 || j >= len(self.coordinates) {
		j = len(self.coordinates) - 1
	}

	v0 = &self.coordinates[i]
	var box = mbr.CreateMBR(v0[X], v0[Y], v0[X], v0[Y])

	self.bbox = new_mono_mbr(box)
	var mono = new_mono_mbr(box)

	self.xy_monobox(&mono, i, i)
	self.chains = append(self.chains, mono)
	var m_index = len(self.chains) - 1

	var mono_size = 0
	for i = i + 1; i <= j; i += 1 {
		v0 = &self.coordinates[i-1]
		v1 = &self.coordinates[i]
		dx = v1[X] - v0[X]
		dy = v1[Y] - v0[Y]

		self.length += math.Hypot(dx, dy)

		cur_x = xy_sign(dx)
		cur_y = xy_sign(dy)

		if prev_x == null {
			prev_x = cur_x
		}

		if prev_y == null {
			prev_y = cur_y
		}

		//((cur_x + prev_x > 0) && (prev_y + cur_y > 0))
		mono_size += 1
		if prev_x == cur_x && prev_y == cur_y && mono_size <= mono_limit {
			self.xy_monobox(&self.chains[m_index], i, null)
		} else {
			mono_size = 1

			prev_x, prev_y = cur_x, cur_y
			p0, p1 := self.coordinates[i-1], self.coordinates[i]
			box := mbr.CreateMBR(p0[X], p0[Y], p1[X], p1[Y])

			mono = new_mono_mbr(box)
			self.xy_monobox(&mono, i-1, i)
			self.chains = append(self.chains, mono)
			 m_index = len(self.chains) - 1
		}
	}
	return self
}

//compute bbox of x or y mono chain
func (self *LineString) xy_monobox(mono *MonoMBR, i, j int) {
	if i != null {
		mono.ExpandIncludeXY(self.coordinates[i][X], self.coordinates[i][Y])
		if j == null {
			mono.j = i
		} else {
			mono.i, mono.j = i, j
		}

		self.bbox.ExpandIncludeMBR(&mono.MBR)
		if self.bbox.i == null {
			self.bbox.i, self.bbox.j = mono.i, mono.j
		} else {
			if mono.j > self.bbox.j {
				self.bbox.j = mono.j
			}
		}
	}
}

//find the sign of value -1, 0 , 1
func xy_sign(v float64) int {
	var i int
	if v > 0 {
		i = 1
	} else if v < 0 {
		i = -1
	}
	return i
}
