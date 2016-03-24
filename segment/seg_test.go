package segment

import (
    "testing"
    . "github.com/franela/goblin"
    . "github.com/intdxdt/simplex/geom/point"
)

func TestSegment(t *testing.T) {
    g := Goblin(t)

    g.Describe("Segment", func() {
        g.It("should test segment intersection", func() {
            a := &Point{0, 0}
            b := &Point{-3, 4}
            c := &Point{1.5, -2}
            d := &Point{-1.5, 2}
            e := &Point{0.5, 3}
            //f := &Point{-2, -2}
            gk := &Point{-1.5, -2.5}
            h := &Point{0.484154648492778, -0.645539531323704}
            i := &Point{0.925118053504632, -1.233490738006176}
            k := &Point{2, 2}
            n := &Point{1,5}

            seg_ab := &Segment{a, b}
            seg_de := &Segment{d, e}

            seg_cd := &Segment{c, d}
            seg_gkh := &Segment{gk, h}
            seg_hi := &Segment{h, i}
            seg_ak := &Segment{a , k}
            seg_kn := &Segment{k , n}

            pts, ok := seg_ab.Intersection(seg_de, false)
            g.Assert(ok).Equal(true)
            g.Assert(pts[0]).Equal(&Point{-1.5, 2})

            pts, ok = seg_ab.Intersection(seg_cd, false)
            g.Assert(ok).Equal(true)
            g.Assert(pts[0]).Equal(&Point{-1.5, 2})
            g.Assert(pts[1]).Equal(&Point{0.0, 0.0})

            pts, ok = seg_gkh.Intersection(seg_cd, false)
            g.Assert(ok).Equal(true)
            g.Assert(len(pts)).Equal(1) //at h

            pts, ok = seg_hi.Intersection(seg_cd, false)
            g.Assert(ok).Equal(true)
            g.Assert(len(pts)).Equal(2) //at h, i

            pts, ok = seg_hi.Intersection(seg_ab, false)
            g.Assert(seg_hi.Intersects(seg_ab, false)).Equal(ok)
            g.Assert(ok).Equal(false)
            g.Assert(len(pts)).Equal(0) //empty

            pts, ok = seg_ak.Intersection(seg_kn, false)
            g.Assert(seg_ak.Intersects(seg_kn, false)).Equal(ok)
            g.Assert(ok).Equal(true)
            g.Assert(len(pts)).Equal(1)//at k
            g.Assert(pts[0]).Equal(k) //k
        })
    })

}

