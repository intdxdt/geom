package geom

import (
	"fmt"
	"bytes"
	"github.com/intdxdt/math"
	"github.com/intdxdt/algor"
)

const ptSetN = 32
const ptLoad = 1000
const splitFactor = 2

//PtSet type
type PtSet struct {
	cmp     func(a, b interface{}) int
	list    *SubPtSet
	maxes   []interface{}
	offsets []int
	load    int
}

//NewPtSet Sorted Set
func NewPtSet() *PtSet {
	var cmp = ptCmp
	var ldN = ptLoad

	var maxCmp = createMaxCmp(cmp)
	var list = NewSubPtSet(maxCmp)

	return &PtSet{
		cmp:  cmp,
		list: list,
		load: math.MinInt(ldN, ptLoad),
	}
}

//Clone PtSet
func (s *PtSet) Clone() *PtSet {
	clone := NewPtSet()
	view  := s.list.DataView()
	for _, sb := range view {
		sub := sb.(*subset)
		clone.addSubset(sub.clone())
	}
	return clone
}

func (s *PtSet) Size() int {
	var n = 0
	if s.IsEmpty() {
		return n
	}
	var view = s.list.DataView()
	var sub = view[len(view)-1].(*subset)
	n = sub.offset + sub.size()
	return n
}

func (s *PtSet) IsEmpty() bool {
	return s.list.IsEmpty()
}

//First item in s
func (s *PtSet) First() interface{} {
	if !s.IsEmpty() {
		sub := s.list.Get(0).(*subset)
		return sub.set.First()
	}
	return nil
}

//Last Item in s
func (s *PtSet) Last() interface{} {
	if !s.IsEmpty() {
		sub := s.list.Get(-1).(*subset)
		return sub.set.Last()
	}
	return nil
}

//Get value at given index in O(lgN)
func (s *PtSet) Get(index int) interface{} {
	if index < 0 {
		index += s.Size()
	}

	idx := s.findSubsetByIndex(index)
	if idx >= 0 {
		sub := s.list.Get(idx).(*subset)
		i, j := sub.offset, sub.offset+sub.size()
		if index >= i && index < j {
			idx = index - i //i=offset
			val := sub.set.Get(idx)
			return val
		}
	}
	return nil
}

//Contains item for the presence of a value in the Array - O(2lgN)
func (s *PtSet) Contains(items ...interface{}) bool {
	if s.IsEmpty() {
		return false
	}
	view := s.list.DataView()

	var idx int
	var sub *subset
	var bln = true
	var n = len(items)

	for i := 0; bln && i < n; i++ {
		idx = s.findSubsetByMax(items[i])
		sub = view[idx].(*subset)
		bln = sub.set.Contains(items[i])
	}
	return bln
}

//IndexOf item for the presence of a value in the Array - O(2lgN)
func (s *PtSet) IndexOf(item interface{}) int {
	idx := -1
	if s.IsEmpty() {
		return idx
	}
	view := s.list.DataView()
	idx = s.findSubsetByMax(item)
	sub := view[idx].(*subset)
	idx = sub.set.IndexOf(item)
	if idx >= 0 {
		idx += sub.offset
	}
	return idx
}

//Values of the set
func (s *PtSet) Values() []interface{} {
	vals := make([]interface{}, 0)
	view := s.list.DataView()
	for i := 0; i < len(view); i++ {
		sub := view[i].(*subset)
		vals = append(vals, sub.vals()...)
	}
	return vals
}

//NextItem gets next given item in the sorted set
func (s *PtSet) NextItem(v interface{}) interface{} {
	if s.IsEmpty() {
		return nil
	}
	idx := s.IndexOf(v)
	n := s.Size() - 1

	var prev interface{} = nil
	if idx >= 0 && idx < n {
		prev = s.Get(idx + 1)
	}
	return prev
}

//PrevItem gets previous given item in the sorted s
func (s *PtSet) PrevItem(v interface{}) interface{} {
	if s.IsEmpty() {
		return nil
	}
	idx := s.IndexOf(v)
	n := s.Size() - 1

	var prev interface{} = nil
	if idx > 0 && idx <= n {
		prev = s.Get(idx - 1)
	}
	return prev
}

//Loop through items in the queue with a callback
// if callback returns bool. Break looping with callback
// return as false
func (s *PtSet) ForEach(fn func(interface{}, int) bool) {
	vals := s.Values()
	for i, v := range vals {
		if !fn(v, i) {
			break
		}
	}
}

//Filters items based on predicate : func (item Item, i int) bool
func (s *PtSet) Filter(fn func(interface{}, int) bool) []interface{} {
	var items = make([]interface{}, 0)
	s.ForEach(func(v interface{}, i int) bool {
		if fn(v, i) {
			items = append(items, v)
		}
		return true
	})
	return items
}

func (s *PtSet) String() string {
	var buffer bytes.Buffer
	view := s.Values()
	n := len(view) - 1

	buffer.WriteString("[")
	for i, o := range view {
		token := fmt.Sprintf("%v", o)
		if i < n {
			token += ", "
		}
		buffer.WriteString(token)
	}
	buffer.WriteString("]")
	return buffer.String()
}

func (s *PtSet) addSubset(sub *subset) *PtSet {
	s.list.Add(sub)
	return s
}

func (s *PtSet) allocSubset() *subset {
	var sub = NewSubPtSet(s.cmp, ptSetN)
	return &subset{set: sub, offset: 0}
}

func (s *PtSet) subOverflows(sub *subset) bool {
	return sub.size() > (s.load * 2)
}

func (s *PtSet) splitSub(sub *subset) []*subset {
	var subs = make([]*subset, 0)
	var view = sub.vals()
	var n = len(view)
	var end int
	var chunk []interface{}

	var size = (sub.size() + splitFactor) / splitFactor
	var offset = sub.offset

	for i := 0; i < n; i += size {
		end = i + size
		if end > n {
			end = n
		}

		chunk = view[i:end]
		if len(chunk) > 0 {
			sb := s.allocSubset()
			sb.addVals(chunk)

			sb.offset = offset
			offset += sb.size()

			subs = append(subs, sb)
		}
	}

	return subs
}

//Union - s union
func (s *PtSet) Union(other *PtSet) *PtSet {
	u := s.Clone()
	view := other.Values()
	for _, v := range view {
		u.Add(v)
	}
	return u
}

//Intersection - s intersection
func (s *PtSet) Intersection(other *PtSet) *PtSet {
	inter := NewPtSet()
	view := other.Values()
	for _, v := range view {
		if s.Contains(v) {
			inter.Add(v)
		}
	}
	return inter
}

//Difference- s difference
//items in s not contained in other
func (s *PtSet) Difference(other *PtSet) *PtSet {
	diff := NewPtSet()
	for _, v := range s.Values() {
		if !other.Contains(v) {
			diff.Add(v)
		}
	}
	return diff
}

//SymDifference - symmetric difference with between s and other
//new s with elements in either s or other but not both
func (s *PtSet) SymDifference(other *PtSet) *PtSet {
	return s.Difference(other).Union(
		other.Difference(s),
	)
}

func (s *PtSet) updateIndexAtSubset(sub *subset) *PtSet {
	var prev *subset
	var view = s.list.DataView()
	var idx = s.list.IndexOf(sub)

	if idx > 0 {
		prev = view[idx-1].(*subset)
		sub.offset = prev.offset + prev.size()
	} else if idx == 0 {
		sub.offset = 0
	}

	n := sub.offset + sub.size()
	for i := idx + 1; i < len(view); i++ {
		sub = view[i].(*subset)
		sub.offset = n
		n += sub.size()
	}
	return s
}


func (s *PtSet) findSubsetByIndex(index int) int {
	view := s.list.DataView()
	idx := algor.BS(view, index, offsetCmp, 0, len(view)-1)
	if idx < 0 {
		idx = -idx - 2
	}
	return idx
}

func (s *PtSet) findSubsetByMax(v interface{}) int {
	var view = s.list.DataView()
	var maxCmp = createMaxCmp(s.cmp)
	var n = len(view) - 1

	idx := algor.BS(view, v, maxCmp, 0, n)
	if idx < 0 {
		idx = -idx - 1
	}

	if idx > n {
		idx = n
	}
	return idx
}

func (s *PtSet) Add(v interface{}) *PtSet {
	var sub *subset
	if s.IsEmpty() {
		sub = s.allocSubset()
		sub.add(v)
		s.addSubset(sub)
	} else {
		view := s.list.DataView()
		idx := s.findSubsetByMax(v)
		sub = view[idx].(*subset)
		sub.add(v)
	}
	s.updateIndexAtSubset(sub)

	if s.subOverflows(sub) {
		subs := s.splitSub(sub)
		for _, sb := range subs {
			s.addSubset(sb)
		}
	}
	return s
}

//Extend PtSet given list of values as params
func (s *PtSet) Extend(values ...interface{}) *PtSet {
	for _, v := range values {
		s.Add(v)
	}
	return s
}

//Empty SubPtSet
func (s *PtSet) Empty() *PtSet {
	//given this snap short of the view
	//remove each item. Note : remove changes the original
	//view slice, but given this window in time, remove all subsets
	view := s.list.DataView()
	for i := 0; i < len(view); i++ {
		sub := view[i].(*subset)
		s.list.Remove(sub.max())
		sub.set.Empty()
	}
	return s
}

//Remove item from set
func (s *PtSet) Remove(items ...interface{}) *PtSet {
	if s.IsEmpty() {
		return s
	}

	var idx int
	var sub *subset
	var n = len(items)
	var prev_idx = n - 1
	for i := 0; i < n; i++ {
		idx = s.findSubsetByMax(items[i])
		//update prev index
		prev_idx = math.MinInt(prev_idx, idx-1)

		sub = s.list.Get(idx).(*subset)
		if sub.size() == 1 && s.cmp(sub.max(), items[i]) == 0 {
			s.list.Remove(sub)
		} else {
			sub.set.Remove(items[i])
		}
	}

	//can be -1 if idx == 0
	if prev_idx < 0 {
		prev_idx = 0
	}
	if !s.IsEmpty() {
		sub := s.list.Get(prev_idx).(*subset)
		s.updateIndexAtSubset(sub)
	}

	return s
}

//Pop item from the end of the sorted list
func (s *PtSet) Pop() interface{} {
	var val interface{}
	if s.IsEmpty() {
		return val
	}
	view := s.list.DataView()
	n := len(view)
	sub := view[n-1].(*subset)

	if sub.size() == 1 {
		val = sub.set.Get(-1)
		s.list.Remove(sub)
	} else {
		val = sub.set.Pop()
	}

	return val
}

//PopLeft item from the beginning of the sorted list
func (s *PtSet) PopLeft() interface{} {
	var val interface{}
	if s.IsEmpty() {
		return val
	}

	sub := s.list.Get(0).(*subset)

	if sub.size() == 1 {
		val = sub.set.Get(0)
		s.list.Remove(sub)
	} else {
		val = sub.set.PopLeft()
	}

	if !s.IsEmpty() {
		sub := s.list.Get(0).(*subset)
		s.updateIndexAtSubset(sub)
	}
	return val
}

func createMaxCmp(cmp func(a, b interface{}) int) func(a, b interface{}) int {
	return func(as, bs interface{}) int {
		ma, mb := as, bs
		a, ok := as.(*subset)
		if ok {
			ma = a.max()
		}

		b, ok := bs.(*subset)
		if ok {
			mb = b.max()
		}

		d := cmp(ma, mb)
		if d < 0 {
			return -1
		} else if d > 0 {
			return 1
		}
		return 0
	}
}

func offsetCmp(as, bs interface{}) int {
	var i, j int
	a, ok := as.(*subset)
	if ok {
		i = a.offset
	} else {
		i = as.(int)
	}

	b, ok := bs.(*subset)
	if ok {
		j = b.offset
	} else {
		j = bs.(int)
	}
	return i - j
}

//=====================================================================================================================
type subset struct {
	set    *SubPtSet
	offset int
}

func (sb *subset) clone() *subset {
	return &subset{set: sb.set.Clone(), offset: sb.offset}
}

func (sb *subset) max() interface{} {
	return sb.set.Get(-1)
}

func (sb *subset) add(v interface{}) {
	sb.set.Add(v)
}

func (sb *subset) size() int {
	return sb.set.Size()
}

func (sb *subset) vals() []interface{} {
	return sb.set.DataView()
}

func (sb *subset) addVals(vals []interface{}) {
	for _, v := range vals {
		sb.add(v)
	}
}

//=====================================================================================================================

//PtSet type
type SubPtSet struct {
	cmp      func(a, b interface{}) int
	base     []interface{}
	view     []interface{}
	i        int
	j        int
	initSize int
}

//New Sorted Set
func NewSubPtSet(cmp func(a, b interface{}) int, initSize ...int) *SubPtSet {
	var iSize = ptSetN
	if len(initSize) > 0 {
		iSize = math.MaxInt(1, initSize[0])
	}

	var base, view, i, j = initQue(iSize)
	return &SubPtSet{
		cmp:      cmp,
		base:     base,
		view:     view,
		i:        i,
		j:        j,
		initSize: iSize,
	}
}

//reveal underlying sorted slice of data view
func (s *SubPtSet) DataView() []interface{} {
	return s.view
}

//Clone PtSet
func (s *SubPtSet) Clone() *SubPtSet {
	var base = make([]interface{}, len(s.base))
	copy(base, s.base)
	var view = base[s.i:s.j]
	return &SubPtSet{
		cmp:  s.cmp,
		base: base,
		view: view,
		i:    s.i,
		j:    s.j,
	}
}

//Contains item for the presence of a value in the Array - O(lgN)
func (s *SubPtSet) Contains(items ...interface{}) bool {
	if s.IsEmpty() {
		return false
	}
	var bln = true
	var n = len(items)
	for i := 0; bln && i < n; i++ {
		bln = algor.BS(s.base, items[i], s.cmp, s.i, s.j-1) >= 0
	}
	return bln
}

//IndexOf item in the sorted s  - O(lgN)
func (s *SubPtSet) IndexOf(item interface{}) int {
	var idx = -1
	if s.IsEmpty() {
		return idx
	}
	idx = algor.BS(s.view, item, s.cmp, 0, s.len()-1)
	if idx < 0 {
		idx = -1
	}
	return idx
}

//Size of list
func (s *SubPtSet) Size() int {
	return s.len()
}

//First item in s
func (s *SubPtSet) First() interface{} {
	var r interface{}
	if !s.IsEmpty() {
		r = s.first()
	}
	return r
}

//Last Item in s
func (s *SubPtSet) Last() interface{} {
	var r interface{}
	if !s.IsEmpty() {
		r = s.last()
	}
	return r
}

//NextItem given item in the sorted s
func (s *SubPtSet) NextItem(v interface{}) interface{} {
	if s.IsEmpty() {
		return nil
	}
	var array = s.base
	var n = s.j - 1
	var rawIdx = algor.BS(array[:], v, s.cmp, s.i, n)

	var idx = rawIdx
	if idx < 0 {
		idx = -idx - 1
	}
	var next interface{}
	if rawIdx >= 0 && s.i <= idx && idx < n {
		next = array[idx+1]
	}
	return next
}

//PrevItem gets previous given item in the sorted s
func (s *SubPtSet) PrevItem(v interface{}) interface{} {
	if s.IsEmpty() {
		return nil
	}
	var array = s.base
	var n = s.j - 1
	var rawIdx = algor.BS(array[:], v, s.cmp, s.i, n)

	idx := rawIdx
	if idx < 0 {
		idx = -idx - 1
	}
	var prev interface{}
	if rawIdx >= 0 && s.i < idx && idx <= n {
		prev = array[idx-1]
	}
	return prev
}

//Filters items based on predicate : func (item Item, i int) bool
func (s *SubPtSet) Filter(fn func(interface{}, int) bool) []interface{} {
	var items = make([]interface{}, 0)
	s.ForEach(func(v interface{}, i int) bool {
		if fn(v, i) {
			items = append(items, v)
		}
		return true
	})
	return items
}

//Pop item from the end of the sorted list
func (s *SubPtSet) Pop() interface{} {
	var r interface{}
	if !s.IsEmpty() {
		r = s.pop()
	}
	return r
}

//PopLeft item from the beginning of the sorted list
func (s *SubPtSet) PopLeft() interface{} {
	var r interface{}
	if !s.IsEmpty() {
		r = s.popLeft()
	}
	return r
}

//Values in s
func (s *SubPtSet) Values() []interface{} {
	var values = make([]interface{}, s.len())
	copy(values, s.view)
	return values
}

//Empty SubPtSet
func (s *SubPtSet) Empty() *SubPtSet {
	s.clear()
	return s
}

//Extend PtSet given list of values as params
func (s *SubPtSet) Extend(values ...interface{}) *SubPtSet {
	for _, v := range values {
		s.Add(v)
	}
	return s
}

//First value in PtSet
func (s *SubPtSet) Get(idx int) interface{} {
	if idx < 0 {
		idx += len(s.view)
	}
	return s.view[idx]
}

//Checks if PtSet empty
func (s *SubPtSet) IsEmpty() bool {
	return s.len() == 0
}

func (s *SubPtSet) String() string {
	var buffer bytes.Buffer
	var n = s.len()
	var token string
	buffer.WriteString("[")
	for i, o := range s.view {
		token = fmt.Sprintf("%v", o)
		if i < n-1 {
			token += ", "
		}
		buffer.WriteString(token)
	}
	buffer.WriteString("]")
	return buffer.String()
}

//Loop through items in the queue with a callback
// if callback returns bool. Break looping with callback
// return as false
func (s *SubPtSet) ForEach(fn func(interface{}, int) bool) {
	for i, v := range s.view {
		if !fn(v, i) {
			break
		}
	}
}

//Push item to s - worst case at O(N^2)
//the cost here at O(n^2) is to allow dynamic indexing
//to add an item, Push uses O(lgN) to find where to insert
//and linear time O(1) or O(N-1) to keep s in sorted order
func (s *SubPtSet) Add(item ...interface{}) *SubPtSet {
	for _, v := range item {
		s.add(v)
	}
	return s
}

func (s *SubPtSet) add(v interface{}) *SubPtSet {
	if s.IsEmpty() {
		s.appendLeft(v)
		return s
	}
	//reserve enough room to the left and right
	s.reserve(true, true)
	array := s.base
	n := s.j - 1

	rawIdx := algor.BS(array, v, s.cmp, s.i, n)

	idx := rawIdx
	if idx < 0 {
		idx = -idx - 1
	}

	//o := v.(int)
	if idx == s.i {
		// equal to i
		if rawIdx < 0 {
			s.appendLeft(nil) //make room
		}
		array[s.i] = v
	} else if s.i < idx && idx < n {
		//between i & j exclusive
		if rawIdx < 0 {
			s.append(nil) //make room
			copy(array[idx+1:n+2], array[idx:n+1])
		}
		array[idx] = v
	} else if idx == n {
		//equal to j-1
		if rawIdx < 0 {
			// >> array[n] >> array[n+1]
			s.append(array[n])
		}
		array[n] = v
	} else if idx > n {
		//greater than nth loc
		s.append(v)
	}
	return s
}

func initQue(initSize int) ([]interface{}, []interface{}, int, int) {
	i := initSize / 2
	j := i
	base := make([]interface{}, initSize, initSize)
	view := base[i:j]
	return base, view, i, j
}

//Length of number of items in PtSet
func (s *SubPtSet) len() int {
	return len(s.view)
}

//First value in PtSet
func (s *SubPtSet) first() interface{} {
	return s.Get(0)
}

//Last value in PtSet
func (s *SubPtSet) last() interface{} {
	return s.Get(-1)
}

//Append to right side of PtSet
func (s *SubPtSet) append(o interface{}) *SubPtSet {
	s.reserve(false, true)
	s.base[s.j] = o
	s.j += 1
	s.view = s.base[s.i:s.j]
	return s
}

//AppendLeft: appends to left of PtSet
func (s *SubPtSet) appendLeft(o interface{}) *SubPtSet {
	s.reserve(true, false)

	if s.atPivot() {
		s.j += 1
	} else {
		s.i -= 1
	}
	s.base[s.i] = o

	s.view = s.base[s.i:s.j]
	return s
}

//reserve enough space left or right
// sufficient to contain elements on insert
func (s *SubPtSet) reserve(left, right bool) {
	if left && s.i == 0 {
		s.expandBase()
	}

	if right && s.j == len(s.base) {
		s.expandBase()
	}
}

func (s *SubPtSet) expandBase() {
	bn := len(s.base)
	vn := len(s.view)

	nn := 2 * bn
	if vn+(nn/2-vn/2) >= nn {
		nn = 2 * nn //not big enough
	}

	k := nn / 2
	mid := vn / 2

	ii := k - mid
	jj := ii + vn

	newBase := make([]interface{}, nn)
	copy(newBase[k:], s.view[mid:])
	copy(newBase[k-mid:k], s.view[0:mid])
	s.base = newBase

	s.i, s.j = ii, jj
	s.view = s.base[s.i:s.j]
}

func (s *SubPtSet) atPivot() bool {
	n := len(s.base)
	return s.i == s.j && (s.i >= 0 && s.i < n)
}

//Clear everything in PtSet
func (s *SubPtSet) clear() *SubPtSet {
	s.base, s.view, s.i, s.j = initQue(s.initSize)
	return s
}

//Union - s union
func (s *SubPtSet) Union(other *SubPtSet) *SubPtSet {
	u := s.Clone()
	for _, v := range other.view {
		u.Add(v)
	}
	return u
}

//Intersection - s intersection
func (s *SubPtSet) Intersection(other *SubPtSet) *SubPtSet {
	inter := NewSubPtSet(s.cmp)
	for _, v := range other.view {
		if s.Contains(v) {
			inter.Add(v)
		}
	}
	return inter
}

//Difference- s difference
//items in s not contained in other
func (s *SubPtSet) Difference(other *SubPtSet) *SubPtSet {
	diff := NewSubPtSet(s.cmp)
	for _, v := range s.view {
		if !other.Contains(v) {
			diff.Add(v)
		}
	}
	return diff
}

//SymDifference - symmetric difference with between s and other
//new s with elements in either s or other but not both
func (s *SubPtSet) SymDifference(other *SubPtSet) *SubPtSet {
	return s.Difference(other).Union(
		other.Difference(s),
	)
}

func (s *SubPtSet) pop() interface{} {
	n := len(s.view) - 1
	val := s.view[n]

	s.view[n] = nil
	s.view = s.view[:n]
	s.j -= 1
	return val
}

func (s *SubPtSet) popLeft() interface{} {
	val := s.view[0]
	s.view[0] = nil

	s.view = s.view[1:]
	s.i += 1
	return val
}

func (s *SubPtSet) Remove(item ...interface{}) *SubPtSet {
	for _, v := range item {
		s.rm(v)
	}
	return s
}

//Remove an item by value from the Array
func (s *SubPtSet) rm(v interface{}) *SubPtSet {
	if s.IsEmpty() {
		return s
	}
	array := s.base
	n := s.j - 1

	rawIdx := algor.BS(array, v, s.cmp, s.i, n)

	idx := rawIdx
	if idx < 0 {
		idx = -idx - 1
	}

	if rawIdx >= 0 {
		if idx == s.i {
			// equal to i
			s.PopLeft()
		} else if s.i < idx && idx < n {
			//between i & j exclusive
			copy(array[idx:n+1], array[idx+1:n+1])
			s.Pop()
		} else if idx == n {
			//equal to nth item
			s.Pop()
		}
	}
	return s
}
