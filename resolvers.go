package main

import (
	"github.com/mfigurski80/SentimentAPI/client"
	"github.com/mfigurski80/SentimentAPI/types"
)

func QueryPointResolver(at int64) (types.Point, error) {
	point, err := client.SelectPoint(at)
	return *point, err
}

func QueryPointsResolver(from int64, to int64) ([]types.Point, error) {
	points, err := client.SelectPointsRange(from, to)
	return *points, err
}

func QueryTweetsResolver(at int64) ([]types.Tweet, error) {
	tweets, err := client.SelectTweets(at)
	return *tweets, err
}

func PointTweetsResolver(p types.Point) ([]types.Tweet, error) {
	return QueryTweetsResolver(p.Time)
}

func MutateSubscription(email string, identity string) error {
	return nil
}
