package cache

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
