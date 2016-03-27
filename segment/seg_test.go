package segment

import (
    "testing"
    . "github.com/franela/goblin"
    . "github.com/intdxdt/simplex/geom/point"
    . "github.com/intdxdt/simplex/util/math"
)

func TestSegment(t *testing.T) {
    g := Goblin(t)

    g.Describe("Segment", func() {
        g.It("should test segment intersection", func() {
            a := NewPointXY(0, 0)
            b := NewPointXY(-3, 4)
            c := NewPointXY(1.5, -2)
            d := NewPointXY(-1.5, 2)
            e := NewPointXY(0.5, 3)
            //f := &Point{-2, -2}
            gk := &Point{-1.5, -2.5}
            h := &Point{0.484154648492778, -0.645539531323704}
            i := &Point{0.925118053504632, -1.233490738006176}
            k := &Point{2, 2}
            n := &Point{1, 5}

            seg_ab := NewSegment(a, b)
            seg_de := &Segment{d, e}

            seg_cd := &Segment{c, d}
            seg_gkh := &Segment{gk, h}
            seg_hi := &Segment{h, i}
            seg_ak := &Segment{a, k}
            seg_kn := &Segment{k, n}

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

func TestSegDist(t *testing.T) {
    g := Goblin(t)
    g.Describe("SegSeg and SegToPoint", func() {
        g.It("should test segment seg2seg", func() {
            var a = &Point{-0.8, -2.6};
            var b = &Point{-1, 1}; // {6, 2}; {5, 2}
            var c = &Point{-2, 3};
            var d = &Point{7, -3};

            var expects = Round(1.1094003924504583, 12);
            seg_ab := &Segment{a, b}
            seg_ba := &Segment{b, a}
            seg_cd := &Segment{c, d}
            seg_dc := &Segment{d, c}

            g.Assert(Round(seg_ab.Distance(seg_cd), 12)).Equal(expects)
            g.Assert(Round(seg_ab.Distance(seg_dc), 12)).Equal(expects)
            g.Assert(Round(seg_ba.Distance(seg_cd), 12)).Equal(expects)
            g.Assert(Round(seg_cd.Distance(seg_ab), 12)).Equal(expects)

            a = &Point{16.82295, 10.44635}
            b = &Point{28.99656, 15.76452}
            on_ab := &Point{25.32, 14.16}

            tpoints := []*Point{
                {30., 0.},
                {15.78786, 25.26468},
                {-2.61504, -3.09018},
                {28.85125, 27.81773},
                a, b, on_ab,
            }

            t_dists := []float64{14.85, 13.99, 23.69, 12.05, 0.00, 0.00, 0.00}
            tvect := &Segment{A: a, B: b}
            dists := make([]float64, len(tpoints))

            for i, tp := range tpoints {
                dists[i] = tvect.DistanceToPoint(tp)
            }

            for i, _ := range tpoints {
                g.Assert(Round(dists[i], 2)).Equal(Round(t_dists[i], 2))
            }
        });

    })

}

