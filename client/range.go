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
	return a.l >= a.r
}

func (a *timeRange) equals(b *timeRange) bool {
	return a.l == b.l && a.r == b.r
}

func (a *timeRange) subsetOf(b *timeRange) bool {
	return a.l >= b.l && a.r <= b.r
}
