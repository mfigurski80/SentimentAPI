package client

import (
	"sync"

	"github.com/mfigurski80/SentimentAPI/types"
)

type tweetListMap map[int64]*[]types.Tweet

var tweetCache = struct {
	d tweetListMap
	sync.Mutex
}{
	d: make(tweetListMap, 0),
}
