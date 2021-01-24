package client

import (
	"database/sql"
	"fmt"
	"time"
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

type Tweet struct {
	ID         int
	Time       int64
	CreatedAt  int64
	Sentiment  string
	Confidence int
	Text       string
	Username   string
	Link       string
}

type dbPoint struct {
	Time     []byte
	Positive int
	Negative int
	Retweets int
	Total    int
}

type Point struct {
	Time     int64
	Positive int
	Negative int
	Retweets int
	Total    int
}

func ParseTimeBytes(bytes []byte) int64 {
	t, err := time.Parse("2006-01-02 15:04:05", (string)(bytes))
	if err != nil {
		panic(err.Error())
	}
	return t.Unix()
}

func ParseUnixTime(unixTime int64) string {
	t := time.Unix(unixTime, 0)
	return fmt.Sprintf("%d-%02d-%02d %02d:%02d:%02d",
		t.Year(), t.Month(), t.Day(),
		t.Hour(), t.Minute(), t.Second(),
	)
}

func makeTweet(t *dbTweet) Tweet {
	return Tweet{
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

func makePoint(p *dbPoint) Point {
	return Point{
		Time:     ParseTimeBytes(p.Time),
		Positive: p.Positive,
		Negative: p.Negative,
		Retweets: p.Retweets,
		Total:    p.Total,
	}
}

func scanTweet(scanner *sql.Rows) Tweet {
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

func scanPoint(scanner *sql.Rows) Point {
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
