package client

import (
	"database/sql"

	"github.com/mfigurski80/SentimentAPI/types"
)

type dbTweet struct {
	ID         int
	Time       []byte
	CreatedAt  []byte
	Sentiment  string
	Confidence int
	Text       string
	Username   string
	Link       string
}

func makeTweet(t *dbTweet) types.Tweet {
	return types.Tweet{
		ID:         t.ID,
		Time:       ParseTimeBytes(t.Time),
		CreatedAt:  ParseTimeBytes(t.CreatedAt),
		Sentiment:  t.Sentiment,
		Confidence: t.Confidence,
		Text:       t.Text,
		Username:   t.Username,
		Link:       t.Link,
	}
}

func scanTweet(scanner *sql.Rows) types.Tweet {
	var t dbTweet
	if err := (*scanner).Scan(
		&t.ID,
		&t.Time,
		&t.CreatedAt,
		&t.Sentiment,
		&t.Confidence,
		&t.Text,
		&t.Username,
		&t.Link,
	); err != nil {
		panic(err.Error())
	}
	return makeTweet(&t)
}

type dbPoint struct {
	Time     []byte
	Positive int
	Negative int
	Retweets int
	Total    int
}

func makePoint(p *dbPoint) types.Point {
	return types.Point{
		Time:     ParseTimeBytes(p.Time),
		Positive: p.Positive,
		Negative: p.Negative,
		Retweets: p.Retweets,
		Total:    p.Total,
	}
}

func scanPoint(scanner *sql.Rows) types.Point {
	var p dbPoint
	if err := (*scanner).Scan(
		&p.Time,
		&p.Positive,
		&p.Negative,
		&p.Retweets,
		&p.Total,
	); err != nil {
		panic(err.Error())
	}
	return makePoint(&p)
}

type dbAnalyticsPing struct {
	ip       string
	identity string
	request  string
}
