package cache

import "github.com/mfigurski80/SentimentAPI/types"

type PointCacheType []types.Point

var PointCache = make(PointCacheType, 0)

// GetUncachedRange subtracts given range from cached range
// in order to provide the full range you need to pull
func GetUncachedRange(from int64, to int64) (int64, int64, UpdateFuncType) {
	// Cache: ------[----------]-----
	// Range: ---------[--]---------- : Dont query
	// Range: ----[---------]-------- : Query reduced range on left
	// Range: [--]------------------- : Query increased range on left
	// Range: [---------------------] : Query everything
	if len(PointCache) == 0 {
		return from, to, updateAll
	}
	// build new range
	minCached := PointCache[0].Time
	maxCached := PointCache[len(PointCache)-1].Time
	if minCached < from {
		from = maxCached
	}
	if maxCached > to {
		to = minCached
	}
	// figure out update func
	if from >= to {
		return -1, -1, nil
	}
	if to == minCached {
		return from, to, updateLeft
	}
	if from == maxCached {
		return from, to, updateRight
	}
	return from, to, updateAll
}
