package main

import (
	"fmt"

	"github.com/mfigurski80/SentimentAPI/client"
	"github.com/mfigurski80/SentimentAPI/types"
)

func QueryPointResolver(at int64) (types.Point, error) {
	query := fmt.Sprintf("SELECT * FROM TimeSeries WHERE time = \"%s\" LIMIT 1", client.ParseUnixTime(at))
	rows, err := client.Execute(query)
	if err != nil {
		return types.Point{}, err
	}
	points := client.ReadOutPoints(rows)
	if len(*points) <= 0 {
		return types.Point{}, nil
	}
	return (*points)[0], nil
}

func QueryPointsResolver(from int64, to int64) ([]types.Point, error) {
	query := fmt.Sprintf("SELECT * FROM TimeSeries WHERE time > \"%s\" AND time <= \"%s\" ORDER BY time", client.ParseUnixTime(from), client.ParseUnixTime(to))
	fmt.Println(query)
	rows, err := client.Execute(query)
	if err != nil {
		return nil, err
	}
	points := client.ReadOutPoints(rows)
	fmt.Println(points)
	return *points, nil
}

func QueryTweetsResolver(at int64) ([]types.Tweet, error) {
	query := fmt.Sprintf("SELECT * FROM Tweets WHERE time = \"%s\"", client.ParseUnixTime(at))
	rows, err := client.Execute(query)
	if err != nil {
		return nil, err
	}
	tweets := client.ReadOutTweets(rows)
	return *tweets, nil
}

func PointTweetsResolver(p types.Point) ([]types.Tweet, error) {
	return QueryTweetsResolver(p.Time)
}
