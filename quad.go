package geom

import (
    . "simplex/geom/mbr"
    "golang.org/x/crypto/nacl/box"
)
const (
    minx = iota
    miny
    maxx
    maxy
)

func quad_relation(g, other Geometry){
    /*
     densify geometry by quadrant
    >>> from shapely.wkt import loads
    >>> poly_wkt = "POLYGON (( 580.1306271033712 749.058289279176, 550.6152553246721 720.079560623726, 567.2511921453935 710.4199844052428, 587.6436308288582 710.9566275284918, 602.1329951565832 716.8597018842316, 610.7192851285685 734.568924951451, 598.3764932938398 749.5949324024251, 580.1306271033712 749.058289279176 ))"
    >>> poly_complx_wkt = "POLYGON (( 460 682, 565 734, 680 719, 661 686, 647 641, 615 615, 571 626, 514 640, 460 682 ),( 556 686, 596 705, 612 679, 581 661, 556 686 ),( 495.36214270032923 675.6839401756312, 493.89462403951694 688.8916081229419, 510.5265021953896 690.8482996706916, 525.2016888035125 701.6101031833151, 546.7252958287595 712.3719066959386, 552.5953704720088 709.436869374314, 534.0068007683864 696.7183743139408, 524.2233430296377 689.3807810098793, 544.7686042810099 667.8571739846323, 567.7597299670691 655.1386789242591, 579.9890521405049 646.8227398463227, 563.8463468715696 639.4851465422613, 534.9851465422612 663.4546180021954, 514.9290581778265 682.0431877058178, 495.36214270032923 675.6839401756312 ))"
    >>> line_wkt = "LINESTRING ( 478.16843368604725 729.2024937189603, 509.2937348344935 798.9660997413398, 606.4261401425758 791.4530960158528, 510.5158680021953 722.4743902305162, 509.8303779577426 692.7107613380233, 551.1518984479212 682.5145419962909, 686.1289344127333 724.1864953347972, 727.1708428736174 668.561820791815, 661.9148765093305 654.7239453896818 )"

    >>> ln = loads(line_wkt)
    >>> const_obj = loads("POINT ( 453.0094189163016 742.5680104803437 )") //FFTFFTFFT
    >>> quad_relation(ln, const_obj)
    "FFTFFTFFT"

    >>> const_obj = loads("POINT ( 670.4391564889743 687.1652408103838 )") //TTTTFTTTT | TTTTFTTTT
    >>> quad_relation(ln, const_obj)
    "TTTTFTTTT"

    >>> const_obj = loads(poly_complx_wkt)  // FTFFTTFFF | FTFFTTFFF
    >>> quad_relation(ln, const_obj)
    "FTFFTTFFF"

     :g:geometry
     :other: other geometry
     @returns  :Array}
    */
    opts = quadrants(g, other)
    return _intersects(g, opts)}


func quadrants(g, other Geometry){
    /*
     * Get bounding box intersections
     @param g
     @param other
     @returns   Opt: {obj_bbox: [], const_bbox: [],  ext_lines : Opt}
     @private
    */
    var box = g.Envelope().Clone()
    var other_mbr = other.Envelope().Clone()

    box.ExpandIncludeMBR(other_mbr)
    box.ExpandByDelta(1e2, 1e2)

    var xs = []float64{box[minx], other_mbr[minx], other_mbr[maxx], box[maxx]}
    var ys = []float64{box[miny], other_mbr[miny], other_mbr[maxy], box[maxy]}
    mat = []
    for j :=0 ; j< len(ys) ; j++:
        for i in xrange(len(xs))
        mat.append([(xs[i], ys[j]) ])
    /*
        .(3,0).|.(3,1).|.(3,2).|.(3,3).
               nw      nn      ne
        .(2,0).|.(2,1).|.(2,2).|.(2,3).
               ww      ii      ee
        .(1,0).|.(1,1).|.(1,2).|.(1,3).
               sw      ss      se
        .(0,0).|.(0,1).|.(0,2).|.(0,3).
    //TODO: ii can be improved by changing ii to convex hull
    */
    ispoint = isinstance(other, Point)

    nw = Polygon((mat[2][0], mat[3][0], mat[3][1], mat[2][1], mat[2][0]))
    ne = Polygon((mat[2][2], mat[3][2], mat[3][3], mat[2][3], mat[2][2]))
    sw = Polygon((mat[0][0], mat[1][0], mat[1][1], mat[0][1], mat[0][0]))
    se = Polygon((mat[0][2], mat[1][2], mat[1][3], mat[0][3], mat[0][2]))

    if ispoint:
        nn = LineString((mat[2][1], mat[3][1]))
        ww = LineString((mat[1][0], mat[1][1]))
        ii = Point((mat[1][1],))
        ee = LineString((mat[1][2], mat[1][3]))
        ss = LineString((mat[0][1], mat[1][1]))
    else:
        nn = Polygon((mat[2][1], mat[3][1], mat[3][2], mat[2][2], mat[2][1]))
        ww = Polygon((mat[1][0], mat[2][0], mat[2][1], mat[1][1], mat[1][0]))
        ii = Polygon((mat[1][1], mat[2][1], mat[2][2], mat[1][2], mat[1][1]))
        ee = Polygon((mat[1][2], mat[2][2], mat[2][3], mat[1][3], mat[1][2]))
        ss = Polygon((mat[0][1], mat[1][1], mat[1][2], mat[0][2], mat[0][1]))

    return (
        nw, nn, ne,
        ww, ii, ee,
        sw, ss, se
    )}

//Expand oject intersections with extended bouding box intersections
func _intersects(obj, quads):
    qr = []
    relates = (obj.intersects(q) for q in quads)
    for r in relates:
        qr.append("T" if r else "F")
    return "".join(qr)

