package cache

import (
	"fmt"

	"github.com/mfigurski80/SentimentAPI/client"
	"github.com/mfigurski80/SentimentAPI/types"
)

type PointCacheType []types.Point

var PointCache = make(PointCacheType, 0)

func ReadPoints(from int64, to int64) []types.Point {
	queryFrom, queryTo, cacheUpdateFunc := getUncachedRange(from, to)
	countUpdated := updateCacheForRange(queryFrom, queryTo, cacheUpdateFunc)
	if countUpdated != 0 {
		fmt.Printf("cache updated with #%v\n", countUpdated)
	}

	lo, _ := binSearchPointCache(from)
	hi, _ := binSearchPointCache(to)

	return PointCache[lo:hi:hi]
}

func updateCacheForRange(from int64, to int64, cacheUpdateFunc UpdateFuncType) int {
	if cacheUpdateFunc == nil {
		return 0
	}
	// perform a cache update
	query := fmt.Sprintf(
		"SELECT * FROM TimeSeries WHERE time > \"%s\" AND time < \"%s\" ORDER BY time",
		client.ParseUnixTime(from),
		client.ParseUnixTime(to),
	)
	fmt.Println(query)
	rows, err := client.Execute(query)
	if err != nil {
		panic(err.Error())
	}
	points := client.ReadOutPoints(rows)
	cacheUpdateFunc(points)
	return len(*points)
}

// GetUncachedRange subtracts given range from cached range
// in order to provide the full range you need to pull
func getUncachedRange(from int64, to int64) (int64, int64, UpdateFuncType) {
	// Cache: ------[----------]-----
	// Range: ---------[--]---------- : Dont query
	// Range: ------[----------]----- : Dont query
	// Range: ----[---------]-------- : Query reduced range on left
	// Range: [--]------------------- : Query increased range on left
	// Range: [---------------------] : Query everything and replace cache
	if len(PointCache) == 0 {
		return from, to, updateAll
	}
	// build new range
	minCached := PointCache[0].Time
	maxCached := PointCache[len(PointCache)-1].Time
	fmt.Printf("Getting uncached range\n%v - %v  vs:\n%v - %v\n", from, to, minCached, maxCached)
	if minCached <= from {
		fmt.Println("left range cached")
		from = maxCached
	}
	if maxCached >= to {
		fmt.Println("right range cached")
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
