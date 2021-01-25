package cache

import "github.com/mfigurski80/SentimentAPI/types"

type PointCacheType []types.Point

var PointCache = make(PointCacheType, 0)

// GetUncachedRange subtracts given range from cached range
// in order to provide the full range you need to pull
func GetUncachedRange(from int64, to int64) (int64, int64) {
	// Cache: ------[----------]-----
	// Range: ---------[--]---------- : Dont query
	// Range: ----[---------]-------- : Query reduced range on left
	// Range: [---------------------] : Query everything
	if len(PointCache) == 0 {
		return from, to
	}

	minCached := PointCache[0].Time
	maxCached := PointCache[len(PointCache)-1].Time
	if minCached < from {
		from = maxCached
	}
	if maxCached > to {
		to = minCached
	}
	if from >= to {
		return -1, -1
	}
	return from, to
}

// --- Update Funcs ---
type UpdateFuncType func(*[]types.Point)

func updateLeft(p *[]types.Point) {
	PointCache = append(*p, PointCache...)
}

func updateRight(p *[]types.Point) {
	PointCache = append(PointCache, *p...)
}

func updateAll(p *[]types.Point) {
	PointCache = *p
}

// -- Util --
func binSearchPointCache(target int64) (int, bool) {
	hi, lo := len(PointCache)-1, 0
	for lo <= hi {
		mid := (lo + hi) / 2
		if PointCache[mid].Time == target {
			return mid, true
		} else if PointCache[mid].Time < target {
			lo = mid + 1
		} else {
			hi = mid - 1
		}
	}
	return lo, false
}
