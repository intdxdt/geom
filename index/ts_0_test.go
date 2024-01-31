package index

import (
	"github.com/franela/goblin"
	"github.com/intdxdt/geom/mono"
	"github.com/intdxdt/math"
	"github.com/intdxdt/mbr"
	"sort"
	"testing"
	"time"
)

type Boxes []mbr.MBR[float64]

// Len for sort interface
func (o Boxes) Len() int {
	return len(o)
}

// Swap for sort interface
func (o Boxes) Swap(i, j int) {
	o[i], o[j] = o[j], o[i]
}

// Less sorts boxes lexicographically
func (o Boxes) Less(i, j int) bool {
	var d = o[i].MinX - o[j].MinX
	if d == 0 || math.Abs(d) < math.EPSILON {
		d = o[i].MinY - o[j].MinY
	}
	return d < 0
}

func someData(n int) []mbr.MBR[float64] {
	var data = make([]mbr.MBR[float64], n)
	for i := 0; i < n; i++ {
		data[i] = mbr.CreateMBR(float64(i), float64(i), float64(i), float64(i))
	}
	return data
}

func testResults(g *goblin.G, objects []*mono.MBR, boxes Boxes) {
	var results = make([]mbr.MBR[float64], 0, len(objects))
	for i := range objects {
		results = append(results, objects[i].MBR)
	}

	sort.Sort(Boxes(results))
	sort.Sort(boxes)
	g.Assert(len(results)).Equal(len(boxes))
	for i, n := range results {
		g.Assert(n.Equals(&boxes[i])).IsTrue()
	}
}

func getObjs(nodes []node) []*mono.MBR {
	var objs = make([]*mono.MBR, 0, len(nodes))
	for _, o := range nodes {
		objs = append(objs, o.item)
	}
	return objs
}

// data from rbush 1.4.2
var data = []mbr.MBR[float64]{{0, 0, 0, 0}, {10, 10, 10, 10}, {20, 20, 20, 20}, {25, 0, 25, 0}, {35, 10, 35, 10}, {45, 20, 45, 20}, {0, 25, 0, 25}, {10, 35, 10, 35},
	{20, 45, 20, 45}, {25, 25, 25, 25}, {35, 35, 35, 35}, {45, 45, 45, 45}, {50, 0, 50, 0}, {60, 10, 60, 10}, {70, 20, 70, 20}, {75, 0, 75, 0},
	{85, 10, 85, 10}, {95, 20, 95, 20}, {50, 25, 50, 25}, {60, 35, 60, 35}, {70, 45, 70, 45}, {75, 25, 75, 25}, {85, 35, 85, 35}, {95, 45, 95, 45},
	{0, 50, 0, 50}, {10, 60, 10, 60}, {20, 70, 20, 70}, {25, 50, 25, 50}, {35, 60, 35, 60}, {45, 70, 45, 70}, {0, 75, 0, 75}, {10, 85, 10, 85},
	{20, 95, 20, 95}, {25, 75, 25, 75}, {35, 85, 35, 85}, {45, 95, 45, 95}, {50, 50, 50, 50}, {60, 60, 60, 60}, {70, 70, 70, 70}, {75, 50, 75, 50},
	{85, 60, 85, 60}, {95, 70, 95, 70}, {50, 75, 50, 75}, {60, 85, 60, 85}, {70, 95, 70, 95}, {75, 75, 75, 75}, {85, 85, 85, 85}, {95, 95, 95, 95}}

func TestRtreeRbush(t *testing.T) {
	g := goblin.Goblin(t)

	g.Describe("Index Tests - From Rbush", func() {
		g.It("should test load 9 & 10", func() {
			var tree0 = NewIndex().loadBoxes(someData(0))
			g.Assert(tree0.data.height).Equal(1)

			var tree1 = NewIndex(8).loadBoxes(someData(8))
			g.Assert(tree1.data.height).Equal(1)

			var tree2 = NewIndex(8).loadBoxes(someData(10))
			g.Assert(tree2.data.height).Equal(2)
		})

		g.It("tests search with some other", func() {
			var data = []mbr.MBR[float64]{
				{-115, 45, -105, 55}, {105, 45, 115, 55}, {105, -55, 115, -45}, {-115, -55, -105, -45},
			}
			var tree = NewIndex(4)
			tree.loadBoxes(data)
			testResults(g, tree.Search(mbr.CreateMBR[float64](-180, -90, 180, 90)), []mbr.MBR[float64]{
				{-115, 45, -105, 55}, {105, 45, 115, 55}, {105, -55, 115, -45}, {-115, -55, -105, -45},
			})

			testResults(g, tree.Search(mbr.CreateMBR[float64](-180, -90, 0, 90)), []mbr.MBR[float64]{
				{-115, 45, -105, 55}, {-115, -55, -105, -45},
			})
			testResults(g, tree.Search(mbr.CreateMBR[float64](0, -90, 180, 90)), []mbr.MBR[float64]{
				{105, 45, 115, 55}, {105, -55, 115, -45},
			})
			testResults(g, tree.Search(mbr.CreateMBR[float64](-180, 0, 180, 90)), []mbr.MBR[float64]{
				{-115, 45, -105, 55}, {105, 45, 115, 55},
			})
			testResults(g, tree.Search(mbr.CreateMBR[float64](-180, -90, 180, 0)), []mbr.MBR[float64]{
				{105, -55, 115, -45}, {-115, -55, -105, -45},
			})
		})

		g.It("#load uses standard insertion when given a low number of items", func() {
			var tree = NewIndex(8).loadBoxes(data)
			tree.loadBoxes(data[0:3])
			var tree2 = NewIndex(8).loadBoxes(data).insert(
				&mono.MBR{MBR: data[0]},
			).insert(&mono.MBR{MBR: data[1]}).insert(&mono.MBR{MBR: data[2]})
			g.Assert(tree.data).Eql(tree2.data)
		})

		g.It("#load does nothing if loading empty data", func() {
			var tree = NewIndex(0).Load(make([]mono.MBR, 0))
			g.Assert(tree.IsEmpty()).IsTrue()
		})

		g.It("#load properly splits tree root when merging trees of the same height", func() {
			var cloneData = make([]mbr.MBR[float64], len(data))
			for i := 0; i < len(data); i++ {
				cloneData[i] = data[i].Clone()
			}
			for i := 0; i < len(data); i++ {
				cloneData = append(cloneData, data[i].Clone())
			}
			var tree = NewIndex(4).loadBoxes(data).loadBoxes(data)
			testResults(g, tree.All(), cloneData)
		})

		g.It("#load properly merges data of smaller or bigger tree heights", func() {
			var smaller = someData(10)
			var cloneData = make([]mbr.MBR[float64], len(data))
			for i := 0; i < len(data); i++ {
				cloneData[i] = data[i].Clone()
			}
			for i := 0; i < len(smaller); i++ {
				cloneData = append(cloneData, smaller[i].Clone())
			}

			var tree1 = NewIndex(4).loadBoxes(data).loadBoxes(smaller)
			var tree2 = NewIndex(4).loadBoxes(smaller).loadBoxes(data)
			g.Assert(tree1.data.height).Equal(tree2.data.height)
			testResults(g, tree1.All(), cloneData)
			testResults(g, tree2.All(), cloneData)
		})

		g.It("#load properly merges data of smaller or bigger tree heights 2", func() {
			N = 8020
			var smaller = GenDataItems(N, 1)
			var larger = GenDataItems(2*N, 1)
			var cloneData = make([]mbr.MBR[float64], len(larger))

			for i := 0; i < len(larger); i++ {
				cloneData[i] = larger[i].Clone()
			}
			for i := 0; i < len(smaller); i++ {
				cloneData = append(cloneData, smaller[i].Clone())
			}

			var tree1 = NewIndex(64).loadBoxes(larger).loadBoxes(smaller)
			var tree2 = NewIndex(64).loadBoxes(smaller).loadBoxes(larger)
			g.Assert(tree1.data.height).Equal(tree2.data.height)
			testResults(g, tree1.All(), cloneData)
			testResults(g, tree2.All(), cloneData)
		})

		g.It("#search finds matching points in the tree given a bbox", func() {
			var tree = NewIndex(4).loadBoxes(data)
			var result = tree.Search(mbr.CreateMBR[float64](40, 20, 80, 70))
			testResults(g, result, []mbr.MBR[float64]{
				{70, 20, 70, 20}, {75, 25, 75, 25}, {45, 45, 45, 45}, {50, 50, 50, 50}, {60, 60, 60, 60}, {70, 70, 70, 70},
				{45, 20, 45, 20}, {45, 70, 45, 70}, {75, 50, 75, 50}, {50, 25, 50, 25}, {60, 35, 60, 35}, {70, 45, 70, 45},
			})
		})

		g.It("#collides returns true when search finds matching points", func() {
			var tree = NewIndex(4).loadBoxes(data)
			g.Assert(tree.Collides(mbr.CreateMBR[float64](40, 20, 80, 70))).IsTrue()
			g.Assert(tree.Collides(mbr.CreateMBR[float64](200, 200, 210, 210))).IsFalse()
		})

		g.It("#search returns an empty array if nothing found", func() {
			var result = NewIndex(4).loadBoxes(data).Search(
				mbr.CreateMBR[float64](200, 200, 210, 210),
			)
			g.Assert(len(result)).Equal(0)
		})

		g.It("#all <==>.data returns all points in the tree", func() {
			var cloneData = make([]mbr.MBR[float64], len(data))
			for i := 0; i < len(data); i++ {
				cloneData[i] = data[i]
			}

			var tree = NewIndex(4).loadBoxes(data)
			var result = tree.Search(mbr.CreateMBR[float64](0, 0, 100, 100))
			testResults(g, result, cloneData)
		})

		g.It("#insert adds an item to an existing tree correctly", func() {
			var data = []mbr.MBR[float64]{{0, 0, 0, 0}, {2, 2, 2, 2}, {1, 1, 1, 1}}
			var tree = NewIndex(4)
			tree.loadBoxes(data)
			tree.insert(&mono.MBR{MBR: mbr.CreateMBR[float64](3, 3, 3, 3)})
			g.Assert(tree.data.leaf).IsTrue()
			g.Assert(tree.data.height).Equal(1)
			var box = mbr.CreateMBR[float64](0, 0, 3, 3)
			g.Assert(tree.data.bbox.Equals(&box)).IsTrue()
			testResults(g, getObjs(tree.data.children), []mbr.MBR[float64]{
				{0, 0, 0, 0}, {1, 1, 1, 1}, {2, 2, 2, 2}, {3, 3, 3, 3},
			})
		})

		g.It("#insert does nothing if given nil", func() {
			var o *mono.MBR
			var tree = NewIndex(4).loadBoxes(data)
			g.Assert(tree.data).Eql(NewIndex(4).loadBoxes(data).insert(o).data)
		})

		g.It("#insert forms a valid tree if items are inserted one by one", func() {
			var tree = NewIndex(4)
			for i := 0; i < len(data); i++ {
				tree.insert(&mono.MBR{MBR: data[i]})
			}

			var tree2 = NewIndex(4).loadBoxes(data)
			g.Assert(tree.data.height-tree2.data.height <= 1).IsTrue()

			var boxes2 = make([]mbr.MBR[float64], 0)
			var all2 = tree2.All()
			for i := 0; i < len(all2); i++ {
				boxes2 = append(boxes2, all2[i].MBR)
			}
			testResults(g, tree.All(), boxes2)
		})

		g.It("#remove removes items correctly", func() {
			var tree = NewIndex(4).loadBoxes(data)
			var N = len(data)
			tree.removeMBR(&data[0])
			tree.removeMBR(&data[1])
			tree.removeMBR(&data[2])

			tree.removeMBR(&data[N-1])
			tree.removeMBR(&data[N-2])
			tree.removeMBR(&data[N-3])
			var cloneData []mbr.MBR[float64]
			for i := 3; i < len(data)-3; i++ {
				cloneData = append(cloneData, data[i].Clone())
			}

			testResults(g, tree.All(), cloneData)

		})

		g.It("#remove does nothing if nothing found", func() {
			var item *mono.MBR
			var tree = NewIndex(0).loadBoxes(data)
			var tree2 = NewIndex(0).loadBoxes(data)
			var query = mbr.CreateMBR[float64](13, 13, 13, 13)
			var querybox = &mono.MBR{MBR: mbr.CreateMBR[float64](13, 13, 13, 13)}
			g.Assert(tree.data).Eql(tree2.removeMBR(&query).data)
			g.Assert(tree.data).Eql(tree2.Remove(querybox).data)
			g.Assert(tree.data).Eql(tree2.Remove(item).data)
		})

		g.It("#remove brings the tree to a clear state when removing everything one by one", func() {
			var tree = NewIndex(4).loadBoxes(data)
			var result = tree.Search(mbr.CreateMBR[float64](0, 0, 100, 100))
			for i := 0; i < len(result); i++ {
				tree.Remove(result[i])
			}
			g.Assert(tree.Remove(&mono.MBR{}).IsEmpty()).IsTrue()
		})

		g.It("#clear should clear all the data in the tree", func() {
			var tree = NewIndex(4).loadBoxes(data).Clear()
			g.Assert(tree.IsEmpty()).IsTrue()
		})

		g.It("should have chainable API", func() {
			g.Assert(NewIndex(4).loadBoxes(data).insert(
				&mono.MBR{MBR: data[0]},
			).removeMBR(&data[0]).Clear().IsEmpty()).IsTrue()
		})
	})

}

/*
g := goblin.Goblin(t)
g.Describe("Index Tests - From Rbush", func() {
*/
func TestRtreeUtil(t *testing.T) {
	g := goblin.Goblin(t)
	g.Describe("Index Util", func() {
		g.It("tests pop nodes", func() {
			g.Timeout(1 * time.Hour)
			var a = createNode(&mono.MBR{MBR: emptyMBR()}, 0, true, nil)
			var b = createNode(&mono.MBR{MBR: emptyMBR()}, 1, true, nil)
			var c = createNode(&mono.MBR{MBR: emptyMBR()}, 1, true, nil)
			var nodes = make([]*node, 0)
			var n *node

			n, nodes = popNode(nodes)
			g.Assert(n == nil).IsTrue()

			nodes = []*node{&a, &b, &c}
			g.Assert(len(nodes)).Equal(3)

			n, nodes = popNode(nodes)
			g.Assert(len(nodes)).Equal(2)
			g.Assert(n == &c).IsTrue()

			n, nodes = popNode(nodes)
			g.Assert(len(nodes)).Equal(1)
			g.Assert(n == &b).IsTrue()

			n, nodes = popNode(nodes)
			g.Assert(len(nodes)).Equal(0)
			g.Assert(n == &a).IsTrue()

			n, nodes = popNode(nodes)
			g.Assert(len(nodes)).Equal(0)
			g.Assert(n == nil).IsTrue()

			var nodes_abc = []node{a, b, c}
			g.Assert(len(nodes_abc)).Equal(3)
			nodes_abc = removeNode(nodes_abc, 1)
			g.Assert(len(nodes_abc)).Equal(2)
			nodes_abc = removeNode(nodes_abc, 4)
			g.Assert(len(nodes_abc)).Equal(2)

		})

		g.It("tests pop index", func() {
			a := 0
			b := 1
			c := 2
			var indexes = make([]int, 0)
			var n int

			n = popIndex(&indexes)
			g.Assert(n == 0).IsTrue()

			indexes = []int{a, b, c}
			g.Assert(len(indexes)).Equal(3)

			n = popIndex(&indexes)
			g.Assert(len(indexes)).Equal(2)
			g.Assert(n).Eql(c)

			n = popIndex(&indexes)
			g.Assert(len(indexes)).Equal(1)
			g.Assert(n).Eql(b)

			n = popIndex(&indexes)
			g.Assert(len(indexes)).Equal(0)
			g.Assert(n).Eql(a)

			n = popIndex(&indexes)
			g.Assert(len(indexes)).Equal(0)
			g.Assert(n == 0).IsTrue()
		})

		g.It("util max min", func() {
			g.Assert(max(2.3, 4.5)).Equal(4.5)
			g.Assert(max(4.5, 2.3)).Equal(4.5)
			g.Assert(min(2.3, 4.5)).Equal(2.3)
			g.Assert(min(4.5, 2.3)).Equal(2.3)
			g.Assert(min(4.50233, 4.50233)).Equal(4.50233)
			g.Assert(max(4.50233, 4.50233)).Equal(4.50233)
		})
	})

}
