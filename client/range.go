package client

type timeRange struct {
	l int64 // left
	r int64 // right
}

func makeNullTimeRange() *timeRange {
	return &timeRange{-1, -1}
}

func makeTimeRange(l int64, r int64) *timeRange {
	return &timeRange{l, r}
}

func (a *timeRange) contains(t int64) bool {
	return t >= a.l && t <= a.r
}

func (a *timeRange) isNull() bool {
	return a.l == -1 && a.r == -1
}

func (a *timeRange) isInvalid() bool {
	return a.l > a.r
}

func (a *timeRange) equals(b *timeRange) bool {
	return a.l == b.l && a.r == b.r
}

func (a *timeRange) subsetOf(b *timeRange) bool {
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
	if a.subsetOf(b) {
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
	// a.intersect(b).subsetOf(a and b)
	r := &timeRange{
		l: max(a.l, b.l),
		r: min(a.r, b.r),
	}
	if r.l > r.r {
		return &timeRange{}
	}
	return r
}
