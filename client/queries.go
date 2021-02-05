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

func InsertSubscription(email string, identity string) error {
	tx, err := analytics.Begin()
	if err != nil {
		return err
	}
	stmt, err := tx.Prepare("INSERT INTO Subscriptions (email, identity) VALUES (?, ?)")
	if err != nil {
		tx.Rollback()
		return err
	}
	_, err = stmt.Exec(email, identity)
	if err != nil {
		tx.Rollback()
		return err
	}
	err = tx.Commit()
	if err != nil {
		stmt.Close()
		tx.Rollback()
		return err
	}
	stmt.Close()
	return nil
}

func SelectPointsRange(from int64, to int64) (*[]types.Point, error) {
	// figure out how to interact with cache
	pointRange := &timeRange{RoundUnixTime(from), RoundUnixTime(to)}
	_, uncachedRange, updateCache := getCachedAndUpdateRanges(pointRange)
	if !uncachedRange.isNull() {
		// get from db
		rows, err := Execute(
			"SELECT * FROM TimeSeries WHERE time > ? AND time <= ? ORDER BY time",
			ParseUnixTime(uncachedRange.l),
			ParseUnixTime(uncachedRange.r),
		)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		updateCache(ReadOutPoints(rows))
	}
	points := getPointRangeFromCache(pointRange)
	return &points, nil
}

func SelectPoint(at int64) (*types.Point, error) {
	rows, err := Execute(
		"SELECT * FROM TimeSeries WHERE time = ? LIMIT 1",
		ParseUnixTime(at),
	)
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
	tweets, exists := tweetCache.d[at]
	if !exists {
		rows, err := Execute(
			"SELECT * FROM Tweets WHERE time = ?",
			ParseUnixTime(at),
		)
		if err != nil {
			return nil, err
		}
		tweets = ReadOutTweets(rows)
		tweetCache.d[at] = tweets
	}
	return tweets, nil
}
