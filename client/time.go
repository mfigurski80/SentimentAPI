package client

import (
	"fmt"
	"time"
)

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

const hour = int64(60 * 60)

func RoundUnixTime(unixTime int64) int64 {
	// round to nearest hour
	return unixTime - (unixTime % hour)
}
