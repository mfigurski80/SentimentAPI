package client

import (
	"sync"

	"github.com/mfigurski80/SentimentAPI/types"
)

type timeMap = map[int64]int

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
	if c.subsetOf(r) {
		return &timeRange{}, r, pointCacheUpdateAll
	}

	// build new cached, uncached time ranges
	cached := r.intersect(c)
	uncached := &timeRange{
		l: max(r.l, cached.r),
		r: min(r.r, cached.l),
	}
	if uncached.isInvalid() {
		uncached = makeNullTimeRange()
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
