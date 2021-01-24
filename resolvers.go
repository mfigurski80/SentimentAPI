package main

import (
	"fmt"

	"github.com/mfigurski80/SentimentAPI/client"
	"github.com/mfigurski80/SentimentAPI/types"
)

func QueryPointResolver(at int64) (types.Point, error) {
	query := fmt.Sprintf("SELECT * FROM TimeSeries WHERE Time = \"%s\" LIMIT 1", client.ParseUnixTime(at))
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

func QueryPointsResolver(from int, to int) ([]types.Point, error) {
	return nil, nil
}

func QueryTweetsResolver(at int) ([]types.Tweet, error) {
	return nil, nil
}

func PointTweetsResolver(p types.Point) ([]types.Tweet, error) {
	return nil, nil
}
