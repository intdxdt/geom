package index

import (
	"time"
	"math"
	"testing"
	"math/rand"
	"github.com/intdxdt/mbr"
	"github.com/intdxdt/geom/mono"
)

func RandBox(size float64, rnd *rand.Rand) mbr.MBR {
	var x = rnd.Float64() * (100.0 - size)
	var y = rnd.Float64() * (100.0 - size)
	return mbr.MBR{
		x, y,
		x + size*rnd.Float64(),
		y + size*rnd.Float64(),
	}
}

func GenDataItems(N int, size float64) []mbr.MBR {
	var data = make([]mbr.MBR, N, N)
	var seed = rand.NewSource(time.Now().UnixNano())
	var rnd = rand.New(seed)
	for i := 0; i < N; i++ {
		data[i] = RandBox(size, rnd)
	}
	return data
}

var N = int(1e6)
var maxFill = 64
var BenchData = GenDataItems(N, 1)
var bboxes100 = GenDataItems(1000, 100*math.Sqrt(0.1))
var bboxes10 = GenDataItems(1000, 10)
var bboxes1 = GenDataItems(1000, 1)
var tree = NewIndex(maxFill).loadBoxes(BenchData)
var box *mbr.MBR
var foundTotal int

func Benchmark_Insert_OneByOne_SmallBigData(b *testing.B) {
	var tree = NewIndex(maxFill)
	for i := 0; i < len(BenchData); i++ {
		tree.insert(&mono.MBR{MBR: BenchData[i]})
	}
	box = tree.data.BBox()
}

func Benchmark_Load_Data(b *testing.B) {
	var tree = NewIndex(maxFill)
	tree.loadBoxes(BenchData)
	box = tree.data.BBox()
}

func Benchmark_Insert_Load_SmallBigData(b *testing.B) {
	var tree = NewIndex(maxFill)
	tree.loadBoxes(BenchData)
	box = tree.data.BBox()
}

func BenchmarkRTree_Search_1000_10pct(b *testing.B) {
	var found = 0
	var items []*mono.MBR
	for i := 0; i < 1000; i++ {
		items = tree.Search(bboxes100[i])
		found += len(items)
	}
	foundTotal = found
}
func BenchmarkRTree_Search_1000_1pct(b *testing.B) {
	var found = 0
	var items []*mono.MBR
	for i := 0; i < 1000; i++ {
		items = tree.Search(bboxes10[i])
		found += len(items)
	}
	foundTotal = found
}

func BenchmarkRTree_Search_1000_01pct(b *testing.B) {
	var found = 0
	var items []*mono.MBR
	for i := 0; i < 1000; i++ {
		items = tree.Search(bboxes1[i])
		found += len(items)
	}
	foundTotal = found
}

func BenchmarkRTree_Build_And_Remove1000(b *testing.B) {
	var tree = NewIndex(maxFill).loadBoxes(BenchData)
	for i := 0; i < 1000; i++ {
		tree = tree.removeMBR(&BenchData[i])
	}
	box = tree.data.BBox()
}
