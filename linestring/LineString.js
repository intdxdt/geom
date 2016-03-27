





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
