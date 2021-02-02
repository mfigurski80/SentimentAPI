package client

import (
	"sync"

	"github.com/mfigurski80/SentimentAPI/types"
)

type timeMap = map[int64]int

type timeRange struct {
	l int64 // left
	r int64 // right
}

func (a *timeRange) contains(t int64) bool {
	return t >= a.l && t <= a.r
}

func (a *timeRange) isSubset(b *timeRange) bool {
	return a.l >= b.l && a.r <= b.r
}

func (a *timeRange) containedBy(b *timeRange) bool {
	return a.l > b.l && a.r < b.r
}

func (a *timeRange) greaterThan(b *timeRange) bool {
	return a.l > b.r
}

func (a *timeRange) lessThan(b *timeRange) bool {
	return a.r < b.l
}

func (a *timeRange) disjointFrom(b *timeRange) bool {
	// return a.greaterThan(b) || a.lessThan(b)
	return a.l > b.r || a.r < b.l
}

func (a *timeRange) union(b *timeRange) *timeRange {
	// Note, if a and b are disjoint, gap will be filled
	return &timeRange{
		l: min(a.l, b.l),
		r: max(a.r, b.r),
	}
}

func (a *timeRange) diff(b *timeRange) *timeRange {
	// Note, if b is subset of a, gap will be filled
	// a.diff(b).union(b) == a
	if a.isSubset(b) {
		return &timeRange{}
	}
	if b.containedBy(a) || b.disjointFrom(a) {
		return a
	}
	return &timeRange{
		l: max(a.l, b.r),
		r: min(a.r, b.l),
	}
}

func (a *timeRange) intersect(b *timeRange) *timeRange {
	// a.intersect(b).subset(a and b)
	r := &timeRange{
		l: max(a.l, b.l),
		r: min(a.r, b.r),
	}
	if r.l > r.r {
		return &timeRange{}
	}
	return r
}

var pointCache = struct {
	d     []types.Point
	index timeMap
	sync.Mutex
}{
	d:     make([]types.Point, 0),
	index: make(timeMap, 0),
}

func getCachedAndUpdateRanges(r *timeRange) (*timeRange, *timeRange, pointCacheUpdateFunc) {
	// Cache: ------[----------]-----
	// Range: ---------[--]---------- : Dont query
	// Range: ------[----------]----- : Dont query
	// Range: ----[---------]-------- : Query reduced range on left
	// Range: [--]------------------- : Query increased range on left
	// Range: [---------------------] : Query range and replace cache
	if len(pointCache.d) == 0 {
		return &timeRange{}, r, pointCacheUpdateAll
	}
	c := &timeRange{
		l: pointCache.d[0].Time,
		r: pointCache.d[len(pointCache.d)-1].Time,
	}
	if c.isSubset(r) {
		return &timeRange{}, r, pointCacheUpdateAll
	}

	// build new cached, uncached time ranges
	cached := r.intersect(c)
	uncached := &timeRange{
		l: max(r.l, cached.r),
		r: min(r.r, cached.l),
	}

	// figure out update function
	if uncached.greaterThan(cached) {
		return cached, uncached, pointCacheUpdateRight
	}
	if uncached.lessThan(cached) {
		return cached, uncached, pointCacheUpdateLeft
	}
	return cached, uncached, pointCacheUpdateAll
}

func min(a int64, b int64) int64 {
	if a > b {
		return b
	}
	return a
}
func max(a int64, b int64) int64 {
	if a > b {
		return a
	}
	return b
}

// -- Cache Update Functions --
type pointCacheUpdateFunc func(*[]types.Point)

func pointCacheUpdateLeft(p *[]types.Point) {
	pointCache.Lock()
	pointCache.d = append(*p, pointCache.d...)
	for i, v := range pointCache.d {
		pointCache.index[v.Time] = i
	}
	pointCache.Unlock()
}

func pointCacheUpdateRight(p *[]types.Point) {
	pointCache.Lock()
	pointCache.d = append(pointCache.d, *p...)
	for i, v := range *p {
		pointCache.index[v.Time] = i
	}
	pointCache.Unlock()
}

func pointCacheUpdateAll(p *[]types.Point) {
	pointCache.Lock()
	pointCache.d = *p
	for i, v := range pointCache.d {
		pointCache.index[v.Time] = i
	}
	pointCache.Unlock()
}
