package main

import "github.com/mfigurski80/SentimentAPI/schema"

func QueryPointResolver(at int) (schema.Point, error) {
	return schema.Point{}, nil
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
