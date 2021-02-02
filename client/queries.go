package client

import (
	"fmt"

	"github.com/mfigurski80/SentimentAPI/types"
)

func InsertAnalyticsPing(identity string, ip string, request string) error {
	query := fmt.Sprintf("INSERT INTO Ping (identity, ip, request) VALUES (\"%s\",\"%s\",\"%s\")", identity, ip, request)
	_, err := analytics.Exec(query)
	return err
}

func SelectPointsRange(from int64, to int64) (*[]types.Point, error) {
	// figure out how to interact with cache
	pointRange := &timeRange{from, to}
	_, uncachedRange, updateCache := getCachedAndUpdateRanges(pointRange)
	if !uncachedRange.isNull() {
		// get from db
		fmt.Printf("updating cache for: %v\n", uncachedRange)
		query := fmt.Sprintf(
			"SELECT * FROM TimeSeries WHERE time > \"%s\" AND time <= \"%s\" ORDER BY time",
			ParseUnixTime(uncachedRange.l),
			ParseUnixTime(uncachedRange.r),
		)
		fmt.Println(query)
		rows, err := Execute(query)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		updateCache(ReadOutPoints(rows))
		fmt.Printf("New Cache has size: %d\n", len(pointCache.d))
	}
	points := getPointRangeFromCache(pointRange)
	return &points, nil
}

func SelectPoint(at int64) (*types.Point, error) {
	query := fmt.Sprintf("SELECT * FROM TimeSeries WHERE time = \"%s\" LIMIT 1", ParseUnixTime(at))
	rows, err := Execute(query)
	if err != nil {
		return nil, err
	}
	points := ReadOutPoints(rows)
	if len(*points) <= 0 {
		return &types.Point{}, nil
	}
	return &(*points)[0], nil
}

func SelectTweets(at int64) (*[]types.Tweet, error) {
	query := fmt.Sprintf("SELECT * FROM Tweets WHERE time = \"%s\"", ParseUnixTime(at))
	rows, err := Execute(query)
	if err != nil {
		return nil, err
	}
	tweets := ReadOutTweets(rows)
	return tweets, nil
}
