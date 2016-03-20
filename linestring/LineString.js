

/*
 description line intersect polygon rings
 param line
 param rings
 returns {*}
 private
 */
proto._line_inter_poly = func (line, rings) {

  var shell = rings[0], i
  bln = line._intersects(shell)
  if !bln {
    //if false, check if shell contains line
    var bln = shell.contains(line)
    var boolhole = false
    //inside shell, does it touch hole boundary ?
    for (i = 1 bln && !boolhole && i < len(rings) ++i) {
      boolhole = line._intersects(rings[i])
    }
    var boolcontains = false
    //inside shell but does not touch the boundary of holes
    if bln && !boolhole {//check if completely contained in hole
      for (i = 1 !boolcontains && i < len(rings) ++i) {
        boolcontains = rings[i].contains(line)
      }
    }
    bln = bln && !boolcontains
  }
  return bln
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
  var bln = false, dist = -1, _dist
  var i, q, lnrange, ibox, qbox, qrng
  var box = qbox = self.mbr.intersection(other.mbr)
  var query = self._searchbox(box)
  var inrange = self.index.search(query)
  if _.is_empty(inrange) {
    //go bruteforce
    dist = self._brutedist(other)
    bln = true
  }
  for (i = 0 !bln && i < len(inrange) ++i) {
    //cur self box
    ibox = inrange[i]
    //search ln using ibox
    query = self._searchbox(ibox)
    lnrange = other.index.search(query)
    if _.is_empty(lnrange) {
      //go bruteforce
      dist = self._brutedist(other)
      bln = true
    }
    for (q = 0 !bln && q < len(lnrange) ++q) {
      qbox = lnrange[q]
      qrng = ibox.intersection(qbox)
      var xor_segs = true //segments when nothing is in range of qrng
      self._segsinrange(selfsegs, qrng, ibox.i, ibox.j, false, xor_segs)
      other._segsinrange(othersegs, qrng, qbox.i, qbox.j, false, xor_segs)

      _dist = self._segseg_mindist(selfsegs, othersegs)
      dist = dist < 0 ? _dist : math.min(_dist, dist)
      if dist == 0.0 {
        bln = true
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
  var bln = false
  var dist = -1, _dist
  for (var a = 0 !bln && a < len(segsa) ++a) {
    for (var b = 0 !bln && b < len(segsb) ++b) {
      bln = segment.intersects(
        segsa[a][0], segsa[a][1],
        segsb[b][0], segsb[b][1]
      )
      if bln {
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
  var dist = -1, bln = false, i, j
  for (i = 0 !bln && i < len(ln) - 1 ++i) {
    for (j = 0 !bln && j < len(ln2) - 1 ++j) {
      var _dist = segment.seg2seg(ln[i], ln[i + 1], ln2[j], ln2[j + 1])
      dist = dist < 0 ? _dist : math.min(_dist, dist)
      if _dist == 0.0 {
        bln = true
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
  var bln = true, bcomplx, chain, inters, jbox, qbox
  var ln1 = [], ln2 = [], ptlist = [], i, j

  for (i = 0 bln && i < self.len(chains) ++i) {
    chain = self.chains[i]
    inters = self.index.search(self._searchbox(chain))

    for (j = 0 bln && j < len(inters) ++j) {
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
      bcomplx && (bln = false)
    }
  }
  return bln
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
