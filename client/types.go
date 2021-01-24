package client

import "database/sql"

type Tweet struct {
	ID         int
	Time       int
	CreatedAt  int
	Sentiment  string
	Confidence int
	Text       string
	Username   string
	Link       string
}

type Point struct {
	Time     []uint8
	Positive int
	Negative int
	Retweets int
	Total    int
}

func scanTweet(scanner *sql.Rows) *Tweet {
	var t Tweet
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
	return &t
}

func scanPoint(scanner *sql.Rows) *Point {
	var p Point
	if err := (*scanner).Scan(
		&p.Time,
		&p.Positive,
		&p.Negative,
		&p.Retweets,
		&p.Total,
	); err != nil {
		panic(err.Error())
	}
	return &p
}
