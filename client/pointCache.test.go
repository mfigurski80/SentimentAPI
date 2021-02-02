package client

import (
	"testing"

	"github.com/mfigurski80/SentimentAPI/types"
)

func r(l int64, r int64) *timeRange {
	return &timeRange{l, r}
}

// TEST DIVIDE REQUEST RANGE

var divideRequestRangeTests = []struct {
	testName           string
	cache              *timeRange
	requestRange       *timeRange
	expectedCached     *timeRange
	expectedUncached   *timeRange
	expectedUpdateFunc pointCacheUpdateFunc
}{
	{"query inside", r(3, 6), r(4, 5), r(4, 5), r(-1, -1), nil},           // query inside
	{"query full", r(3, 6), r(3, 6), r(3, 6), r(-1, -1), nil},             // query full
	{"query extra right", r(3, 6), r(4, 8), r(4, 6), r(7, 8), nil},        // query extra right
	{"query extra left", r(3, 6), r(1, 4), r(3, 4), r(1, 2), nil},         // query extra left
	{"query disjoint right", r(3, 6), r(8, 10), r(-1, -1), r(7, 10), nil}, // query disjoint right
	{"query disjoint left", r(3, 6), r(0, 1), r(-1, -1), r(0, 2), nil},    // query disjoint left
	{"query all", r(3, 6), r(0, 10), r(-1, -1), r(0, 10), nil},
}

func TestDivideRequestRange(t *testing.T) {
	f := getCachedAndUpdateRanges
	for i, test := range divideRequestRangeTests {
		// set up cache
		pointCache.d = make([]types.Point, test.cache.r-test.cache.l+1)
		for j := range pointCache.d {
			pointCache.d[j] = types.Point{Time: test.cache.l + int64(j)}
			pointCache.index[test.cache.l+int64(j)] = j
		}

		actualCached, actualUncached, _ := f(test.requestRange)
		if !actualCached.equals(test.expectedCached) {
			t.Errorf(
				"[Test #%d - %s]\nexpected cached range doesn't match computed (%v != %v)",
				i, test.testName, test.expectedCached, actualCached,
			)
		}
		if !actualUncached.equals(test.expectedUncached) {
			t.Errorf("[Test #%d - %s]\nexpected uncached range doesn't match computed (%v != %v)",
				i, test.testName, test.expectedUncached, actualUncached)
		}

	}
}

// TEST GRAB RANGE FROM CACHE
var getPointRangeTests = []struct {
	name           string
	cache          *timeRange
	request        *timeRange
	expectedResult *timeRange
}{
	{"query from cache", r(0, 10), r(4, 8), r(4, 8)},
	// {"query 0 range from cache", r(0, 10), r(4, 4), r(0, 0)},
}

func TestGetPointRange(t *testing.T) {
	f := getPointRangeFromCache
	for i, test := range getPointRangeTests {
		// set up cache
		pointCache.d = make([]types.Point, test.cache.r-test.cache.l)
		for j := range pointCache.d {
			pointCache.d[j] = types.Point{Time: test.cache.l + int64(j)}
			pointCache.index[test.cache.l+int64(j)] = j
		}

		// run test
		actualResult := f(test.request)
		if len(actualResult) != int(test.expectedResult.r-test.expectedResult.l) {
			t.Errorf(
				"[Test #%d - %s]\nexpected result range doesn't match computed (%v != %v)",
				i, test.name, test.expectedResult, actualResult,
			)
		}
		if actualResult[0].Time != test.expectedResult.l {
			t.Errorf(
				"[Test #%d - %s]\nexpected result start doesn't match computer (%v != %v)",
				i, test.name, test.expectedResult.l, actualResult[0],
			)
		}
	}
}
