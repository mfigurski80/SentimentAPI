package cache

import "github.com/mfigurski80/SentimentAPI/types"

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
