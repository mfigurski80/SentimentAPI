package client

import (
	"fmt"
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

func getPointRangeFromCache(r *timeRange) []types.Point {
	left, exists := pointCache.index[r.l]
	if !exists {
		left = 0
		fmt.Printf("Cache is missing on left (%v vs existing left %v)", r.l, pointCache.d[0].Time)
	}
	left++ // don't include left
	right, exists := pointCache.index[r.r]
	right++ // include right
	if !exists || right > len(pointCache.d) {
		right = len(pointCache.d)
		fmt.Printf("Cache is missing on right (%v vs existing right %v)", r.r, pointCache.d[right-1].Time)
	}
	return pointCache.d[left:right]
}

func getPointFromCache(t int64) types.Point {
	i, _ := pointCache.index[t]
	return pointCache.d[i]
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

	// build new cached, uncached time ranges
	cached := makeNullTimeRange()
	uncached := makeNullTimeRange()
	if r.subsetOf(c) {
		// base case, request is cached
		return r, uncached, nil
	}
	if c.subsetOf(r) {
		// cache is too small from both ends
		return cached, r, pointCacheUpdateAll
	}
	if r.contains(c.l) {
		// left is uncached
		uncached = &timeRange{r.l, c.l - 1}
		cached = &timeRange{c.l, r.r}
		return cached, uncached, pointCacheUpdateLeft
	}
	if r.contains(c.r) {
		// right is uncached
		uncached = &timeRange{c.r + 1, r.r}
		cached = &timeRange{r.l, c.r}
		return cached, uncached, pointCacheUpdateRight
	}
	// cache and req are disjoint
	if r.l > c.r {
		// req on right
		return cached, &timeRange{c.r + 1, r.r}, pointCacheUpdateRight
	}
	// req on left
	return cached, &timeRange{r.l, c.l - 1}, pointCacheUpdateLeft
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
	startIndex := len(pointCache.d)
	pointCache.d = append(pointCache.d, *p...)
	for i, v := range *p {
		pointCache.index[v.Time] = startIndex + i
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
