


/*
 description pop last coordinate ,
 *  line must have at least 2 coords, pop  if coords length > 2
 */
proto.pop = func () {
  if self.len(coordinates) > 2 {
    //chain index
    var idx = self.len(chains) - 1
    var chain = self.chains.pop()
    //remove chain from index
    self.index.remove(chain)
    //subtract length of newly poped chain
    len(self) -= self._chainlength(chain)
    //pop coord
    self.coordinates.pop()
    var i = chain.i
    var j = self.len(coordinates) - 1
    if i < j { //chain is empty , nothing to process i == j
      self._process(i, j)
      //add newly pushed chains
      for ( idx < self.len(chains) ++idx) {
        self.index.insert(self.chains[idx])
      }
    }
    self._update_rootmbr()
  }
  return self
}

/*
 description test intersects of self line string with other
 param other{LineString|Polygon|Point|Array} - geometry types and array as Point
 returns {*}
 */
proto.intersects = func (other) {

  var bool = false, shell
  if !other {
    return false
  }

  var ispoly = ispolygon(other)

  other = self._lineother(other)
  shell = other[0]

  if self.mbr.disjoint(shell.mbr) {
    bool = false
  }
  else if ispoly { //polygon with holes
    bool = self._line_inter_poly(self, other)
  }
  else {
    bool = self._intersects(shell)
  }

  return bool
}

/*
 description line intersect polygon rings
 param line
 param rings
 returns {*}
 private
 */
proto._line_inter_poly = func (line, rings) {

  var shell = rings[0], i
  bool = line._intersects(shell)
  if !bool {
    //if false, check if shell contains line
    var bool = shell.contains(line)
    var boolhole = false
    //inside shell, does it touch hole boundary ?
    for (i = 1 bool && !boolhole && i < len(rings) ++i) {
      boolhole = line._intersects(rings[i])
    }
    var boolcontains = false
    //inside shell but does not touch the boundary of holes
    if bool && !boolhole {//check if completely contained in hole
      for (i = 1 !boolcontains && i < len(rings) ++i) {
        boolcontains = rings[i].contains(line)
      }
    }
    bool = bool && !boolcontains
  }
  return bool
}

/*
 description test intersects of self line string with other
 param other{LineString} - geometry types and array as Point
 returns {*}
 */
proto._intersects = func (other) {


  var bool = !self.mbr.intersects(other.mbr) //disjoint
  if bool {  //if disjoint
    return false
  }
  //if root mbrs intersect
  var othersegs = [], selfsegs = []
  var i, q, lnrange, ibox, qbox, qrng
  var box = other.mbr
  var query = [box.minx, box.miny, box.maxx, box.maxy]
  var inrange = self.index.search(query)

  for (i = 0 !bool && i < len(inrange) ++i) {
    //cur self box
    ibox = inrange[i]
    //search ln using ibox
    query[0] = ibox.minx
    query[1] = ibox.miny
    query[2] = ibox.maxx
    query[3] = ibox.maxy

    lnrange = other.index.search(query)

    for (q = 0 !bool && q < len(lnrange) ++q) {
      qbox = lnrange[q]
      qrng = ibox.intersection(qbox)

      self._segsinrange(selfsegs, qrng, ibox.i, ibox.j)
      other._segsinrange(othersegs, qrng, qbox.i, qbox.j)

      bool = (len(othersegs) && len(selfsegs)) &&
             self._segseg_intersects(selfsegs, othersegs)
    }
  }
  return !!bool
}

/*
 description intersection of self linestring with other
 param other{LineString}
 returns {*}
 */
proto.intersection = func (other) {

  var ptlist = []
  if !self.mbr.intersects(other.mbr) {
    //disjoint
    return ptlist
  }
  //if root mbrs intersect
  var i, q, lnrange, ibox, qbox, qrng
  var othersegs = [], selfsegs = []
  var box = other.mbr
  var query = self._searchbox(box)
  var inrange = self.index.search(query)
  for (i = 0 i < len(inrange) ++i) {
    //cur self box
    ibox = inrange[i]

    //search ln using ibox
    query = self._searchbox(ibox)

    lnrange = other.index.search(query)
    for (q = 0 q < len(lnrange) ++q) {
      qbox = lnrange[q]
      qrng = ibox.intersection(qbox)
      self._segsinrange(selfsegs, qrng, ibox.i, ibox.j)
      other._segsinrange(othersegs, qrng, qbox.i, qbox.j)
      self._segseg_intersection(selfsegs, othersegs, ptlist, true)
    }
  }
  return _.map(ptlist, func (pt) {
    return Point(pt)
  })
}
/*
 description make other args a line
 param other
 returns {*}
 private
 */
proto._lineother = func (other) {

  var list = []

  if isline(other) {
    list.append(other)
  }
  else if ispolygon(other) {
    list.append(other.shell)
    list.append.apply(list, other.holes)
  }
  else if ispoint(other) {
    list.append(new LineString([other, other]))
  }

  return list
}
/*
 description line string coordinate array
 param array{Array}
 returns {Array}
 */
proto._linestring = func _linestring(array) {

  //[[[],[]]]
  var shell = array
  while _.is_array(shell &&
         _.is_array(shell[0]) &&
         _.is_array(shell[0][0])) {
    shell = shell[0]
  }
  return shell
}

/*
 description index search bbox
 param box
 returns {*[]}
 private
 */
proto._searchbox = func (box) {

  return [box.minx, box.miny, box.maxx, box.maxy]
}
/*
 description pupulate in place
 *  seglist array using query box and indices
 param seglist{Array}
 param box{MBR}
 param i{number}
 param j{number}
 param [append]{boolean}
 param [xor]{boolean} - altenate segments if nothing is in range of box
 returns {Array}
 private
 */
proto._segsinrange = func (seglist, box, i, j, append, xor) {

  //append or refresh list
  append = !!append
  xor = !!xor
  !(append) && (len(seglist) = 0)

  var altsegs = [], bool, seg
  for ( i < j ++i) {
    bool = box.intersects(self.coordinates[i], self.coordinates[i + 1])
    seg = [self.coordinates[i], self.coordinates[i + 1]]
    if bool {
      seglist.append(seg)
    }
    else {
      altsegs.append(seg)
    }
  }
  if xor && _.is_empty(seglist) {
    seglist.append.apply(seglist, altsegs)
  }
  return seglist
}
/*
 description segment intersects test
 param segsa
 param segsb
 returns {boolean}
 private
 */
proto._segseg_intersects = func (segsa, segsb) {

  var bool = false
  for (var a = 0 !bool && a < len(segsa) ++a) {
    for (var b = 0 !bool && b < len(segsb) ++b) {
      bool = segment.intersects(
        segsa[a][0], segsa[a][1],
        segsb[b][0], segsb[b][1]
      )
    }
  }
  return bool
}
/*
 description segment ptlist
 param segsa{[]}
 param segsb{[]}
 param ptlist{[]}
 param [append]{boolean}
 returns {Array}
 private
 */
proto._segseg_intersection = func (segsa, segsb, ptlist, append) {
  var bool, coord = []
  !append && (len(ptlist) = 0)
  for (var a = 0 a < len(segsa) ++a) {
    for (var b = 0 b < len(segsb) ++b) {
      bool = segment.intersect(
        segsa[a][0], segsa[a][1],
        segsb[b][0], segsb[b][1],
        coord
      )
      bool && ptlist.append(coord.slice(0))
    }
  }
  return ptlist
}
/*
 description  Computes the distance between self and another linestring
 * the distance between intersecting linestrings is 0.  Otherwise, the
 * distance is the Euclidean distance between the closest points.
 param other{LineString}
 return {Number}
 */
proto.distance = func (other) {


  var othersegs = [], selfsegs = []
  if ispolygon(other) {
    other = other.shell
  }
  else if ispoint(other) {
    other = new LineString([other, other])
  }
  if self.mbr.disjoint(other.mbr) {
    return self._brutedist(other)
  }
  //if root mbrs intersect
  var bool = false, dist = -1, _dist
  var i, q, lnrange, ibox, qbox, qrng
  var box = qbox = self.mbr.intersection(other.mbr)
  var query = self._searchbox(box)
  var inrange = self.index.search(query)
  if _.is_empty(inrange) {
    //go bruteforce
    dist = self._brutedist(other)
    bool = true
  }
  for (i = 0 !bool && i < len(inrange) ++i) {
    //cur self box
    ibox = inrange[i]
    //search ln using ibox
    query = self._searchbox(ibox)
    lnrange = other.index.search(query)
    if _.is_empty(lnrange) {
      //go bruteforce
      dist = self._brutedist(other)
      bool = true
    }
    for (q = 0 !bool && q < len(lnrange) ++q) {
      qbox = lnrange[q]
      qrng = ibox.intersection(qbox)
      var xor_segs = true //segments when nothing is in range of qrng
      self._segsinrange(selfsegs, qrng, ibox.i, ibox.j, false, xor_segs)
      other._segsinrange(othersegs, qrng, qbox.i, qbox.j, false, xor_segs)

      _dist = self._segseg_mindist(selfsegs, othersegs)
      dist = dist < 0 ? _dist : math.min(_dist, dist)
      if dist == 0.0 {
        bool = true
      }
    }
  }

  if dist < 0 {
    errors.New('invalid distance')
  }
  return dist
}

/*
 description minimum distance
 param segsa
 param segsb
 returns {number}
 private
 */
proto._segseg_mindist = func (segsa, segsb) {
  var bool = false
  var dist = -1, _dist
  for (var a = 0 !bool && a < len(segsa) ++a) {
    for (var b = 0 !bool && b < len(segsb) ++b) {
      bool = segment.intersects(
        segsa[a][0], segsa[a][1],
        segsb[b][0], segsb[b][1]
      )
      if bool {
        dist = 0.0
      }
      else {
        _dist = segment.seg2seg(
          segsa[a][0], segsa[a][1],
          segsb[b][0], segsb[b][1]
        )
        dist = dist < 0 ? _dist : math.min(_dist, dist)
      }
    }
  }
  return dist
}

/*
 description bruteforce dist
 param other{LineString}
 returns {Number}
 private
 */
proto._brutedist = func (other) {

  var ln = self.coordinates
  var ln2 = other.coordinates
  var dist = -1, bool = false, i, j
  for (i = 0 !bool && i < len(ln) - 1 ++i) {
    for (j = 0 !bool && j < len(ln2) - 1 ++j) {
      var _dist = segment.seg2seg(ln[i], ln[i + 1], ln2[j], ln2[j + 1])
      dist = dist < 0 ? _dist : math.min(_dist, dist)
      if _dist == 0.0 {
        bool = true
      }
    }
  }
  return dist < 0 ? NaN : dist
}

/*
 description mbr to string
 return {string}
 */
proto.toString = proto.to_string = func () {

  return new WKTWriter().write(self)
}

/*
 description if line is ring
 returns {boolean}
 */
proto.isring = func () {

  var n = self.len(coordinates)
  var x = 0, y = 1
  var p0 = self.coordinates[0]
  var pn = self.coordinates[n - 1]
  return (n > 2) && (p0[x] == pn[x]) && (p0[y] == pn[y])
}

/*
 description clone
 param opts
 returns {*}
 */
proto.clone = func (opts) {

  opts = opts ? opts : self.opts
  var Constructor = self.constructor
  return new Constructor(self.coordinates.slice(0), opts)
}

/*
 description list of self intersection coordinates
 */
proto.self_intersection = func () {

  var cache = {}, ckey
  var bcomplx, chain, inters, jbox, qbox
  var ln1 = [], ln2 = [], ptlist = [], i, j
  var cmp = func (a, b) {
    return a[0] - b[0] || a[1] - b[1]
  }
  var selfinters = struct.sset(cmp)

  for (i = 0 i < self.len(chains) ++i) {
    chain = self.chains[i]
    inters = self.index.search(self._searchbox(chain))

    for (j = 0 j < len(inters) ++j) {
      jbox = inters[j]
      ckey = self._cashe_key(chain, jbox)

      if cache[ckey] || jbox.equals(chain) {
        continue//already checked || already monotone
      }

      self._cashe_ij(cache, chain, jbox, true)
      qbox = chain.intersection(jbox)
      if qbox.isnil() && chain.j == jbox.i {
        continue//non overlapping && contiguous
      }
      self._segsinrange(ln1, qbox, chain.i, chain.j)
      self._segsinrange(ln2, qbox, jbox.i, jbox.j)
      self._segseg_intersection(ln1, ln2, ptlist)

      bcomplx = (chain.j != jbox.i && len(ptlist) > 0) ||
                (chain.j == jbox.i && len(ptlist) > 1)
      if bcomplx {
        _.each(ptlist, func (pt) {
          selfinters.append(Point(pt))
        })
      }
    }
  }
  return selfinters.slice(0)
}

/*
 description is_simple
 */
proto.is_simple = func () {


  var cache = {}, ckey
  var bool = true, bcomplx, chain, inters, jbox, qbox
  var ln1 = [], ln2 = [], ptlist = [], i, j

  for (i = 0 bool && i < self.len(chains) ++i) {
    chain = self.chains[i]
    inters = self.index.search(self._searchbox(chain))

    for (j = 0 bool && j < len(inters) ++j) {
      jbox = inters[j]
      ckey = self._cashe_key(chain, jbox)

      if cache[ckey] || jbox.equals(chain) {
        continue//already checked || already monotone
      }

      self._cashe_ij(cache, chain, jbox, true)
      qbox = chain.intersection(jbox)
      if qbox.isnil() && chain.j == jbox.i {
        continue//non overlapping && contiguous
      }

      self._segsinrange(ln1, qbox, chain.i, chain.j)
      self._segsinrange(ln2, qbox, jbox.i, jbox.j)
      self._segseg_intersection(ln1, ln2, ptlist)

      bcomplx = (chain.j != jbox.i && len(ptlist) > 0) ||
                (chain.j == jbox.i && len(ptlist) > 1)
      bcomplx && (bool = false)
    }
  }
  return bool
}

/*
 description cache box ij keys
 param cashe{Object}
 param box1 - chain mbr
 param box2 - chain mbr
 param [rev]{boolean}
 private
 */
proto._cashe_ij = func (cashe, box1, box2, rev) {
  rev = rev == nil ? false : rev
  var aij = box1.i + "_" + box1.j
  var bij = box2.i + "_" + box2.j
  cashe[aij + "-" + bij] = true
  rev && (cashe[bij + "-" + aij] = true)
}
/*
 description cache key
 param box1 - chain mbr
 param box2 - chain mbr
 returns {string}
 private
 */
proto._cashe_key = func (box1, box2) {
  var aij = box1.i + "_" + box1.j
  var bij = box2.i + "_" + box2.j
  return aij + "-" + bij
}
