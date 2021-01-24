package main

import (
	"fmt"

	"github.com/mfigurski80/SentimentAPI/client"
	"github.com/mfigurski80/SentimentAPI/schema"
)

func QueryPointResolver(at int64) (schema.Point, error) {
	query := fmt.Sprintf("SELECT * FROM TimeSeries WHERE Time = \"%s\" LIMIT 1", client.ParseUnixTime(at))
	rows, err := client.Execute(query)
	if err != nil {
		return schema.Point{}, err
	}
	points := client.ReadOutPoints(rows)
	if len(*points) <= 0 {
		return schema.Point{}, nil
	}
	return (schema.Point)(*points)[0], nil
}

func QueryPointsResolver(from int, to int) ([]schema.Point, error) {
	return nil, nil
}

func QueryTweetsResolver(at int) ([]schema.Tweet, error) {
	return nil, nil
}

func PointTweetsResolver(p schema.Point) ([]schema.Tweet, error) {
	return nil, nil
}
