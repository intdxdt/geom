package geom

import (
    . "github.com/intdxdt/simplex/geom/mbr"
    . "github.com/intdxdt/simplex/geom/point"
    "math"
)

type MonoMBR struct {
    *MBR
    i int
    j int
}

//clone  mono mbr
func (box *MonoMBR) Clone() *MonoMBR {
    clone_mbr := box.MBR.Clone()
    return &MonoMBR{clone_mbr, box.i, box.j}
}

//clone  mono mbr
func ( box *MonoMBR) BBox() *MBR {
    return box.MBR
}

//update mono chain index
func ( box *MonoMBR) update_index(i, j int) {
    box.i = i
    box.j = j
}

//new monotone mbr
func new_mono_mbr(box  *MBR) *MonoMBR {
    return &MonoMBR{box, null, null}
}


//build xymonotone chain, perimeter length,
//monotone build starts from i and ends at j, designed for
//appending new points to the end of line
//param [i]{number} - start index
//param [j]{number} - end index
func (self *LineString) process_chains(i, j int) {
    var dx, dy float64
    var v0, v1 *Point
    var cur_x, cur_y, prev_x, prev_y int
    var mono *MonoMBR

    prev_x, prev_y = null, null
    mono_limit := self.monosize

    if j == 0 || j >= len(self.coordinates) {
        j = len(self.coordinates) - 1
    }

    v0 = self.coordinates[i]
    box := NewMBR(v0[x], v0[y], v0[x], v0[y])

    self.bbox = new_mono_mbr(box)
    box = box.Clone()
    mono = new_mono_mbr(box)

    self.xy_monobox(mono, i, i)
    self.chains = append(self.chains, mono)

    var mono_size = 0
    for i = i + 1; i <= j; i += 1 {
        v0 = self.coordinates[i - 1]
        v1 = self.coordinates[i]
        dx = v1[x] - v0[x]
        dy = v1[y] - v0[y]

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
            self.xy_monobox(mono, i, null)
        } else {
            mono_size = 1

            prev_x, prev_y = cur_x, cur_y
            p0, p1 := self.coordinates[i - 1], self.coordinates[i]
            box := NewMBR(p0[x], p0[y], p1[x], p1[y])

            mono = new_mono_mbr(box)
            self.xy_monobox(mono, i - 1, i)
            self.chains = append(self.chains, mono)
        }
    }
}

//compute bbox of x or y mono chain
func (self *LineString) xy_monobox(mono *MonoMBR, i, j int) {
    if i != null {
        mono.ExpandIncludeXY(self.coordinates[i][x], self.coordinates[i][y])
        if j == null {
            mono.j = i
        } else {
            mono.i, mono.j = i, j
        }

        self.bbox.ExpandIncludeMBR(mono.MBR)
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


