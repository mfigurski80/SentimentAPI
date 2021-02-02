package client

import (
	"testing"

	"github.com/mfigurski80/SentimentAPI/types"
)

func r(l int64, r int64) *timeRange {
	return &timeRange{l, r}
}

var cacheTests = []struct {
	cache              *timeRange
	requestRange       *timeRange
	expectedCached     *timeRange
	expectedUncached   *timeRange
	expectedUpdateFunc pointCacheUpdateFunc
}{
	{r(3, 6), r(4, 5), r(4, 5), r(-1, -1), nil},
	{r(3, 6), r(4, 8), r(4, 6), r(7, 8), nil},
}

func TestDivideRequestRange(t *testing.T) {
	f := getCachedAndUpdateRanges
	for i, test := range cacheTests {
		// set up cache
		pointCache.d = make([]types.Point, test.cache.r-test.cache.l)
		for i := range pointCache.d {
			pointCache.d = append(pointCache.d, types.Point{Time: test.cache.r + int64(i)})
		}

		actualCached, actualUncached, _ := f(test.requestRange)
		if !actualCached.equals(test.expectedCached) {
			t.Fatalf(
				"[Test #%d] expected cached range doesn't match computed (%v != %v)",
				i, actualCached, test.expectedCached,
			)
		}
		if !actualUncached.equals(test.expectedUncached) {
			t.Fatalf("Test #%d] expected uncached range doesn't match computed (%v != %v)",
				i, actualUncached, test.expectedUncached)
		}
	}
}
