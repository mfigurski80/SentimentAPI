package client

import (
	"database/sql"
	"fmt"
	"os"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/mfigurski80/SentimentAPI/types"
)

var database *sql.DB

func parseConnectionString() string {
	username := os.Getenv("MYSQL_USERNAME")
	password := os.Getenv("MYSQL_PASSWORD")

	service := strings.ToUpper(os.Getenv("MYSQL_SERVICE"))
	service = strings.ReplaceAll(service, "-", "_")
	host := os.Getenv(service + "_SERVICE_HOST")
	port := os.Getenv(service + "_SERVICE_PORT")

	return username + ":" + password + "@tcp(" + host + ":" + port + ")/sentiment"
}

func Open() error {
	connString := parseConnectionString()
	fmt.Println("connecting to mysql database at: " + connString)
	db, err := sql.Open("mysql", connString)
	database = db
	return err
}

func Close() {
	database.Close()
}

func Execute(query string) (*sql.Rows, error) {
	return database.Query(query)
}

func ReadOutPoints(rows *sql.Rows) *[]types.Point {
	points := make([]types.Point, 0)
	for rows.Next() {
		points = append(points, scanPoint(rows))
	}
	return &points
}

func ReadOutTweets(rows *sql.Rows) *[]types.Tweet {
	tweets := make([]types.Tweet, 0)
	for rows.Next() {
		tweets = append(tweets, scanTweet(rows))
	}
	return &tweets
}

func start() {
	if err := Open(); err != nil {
		panic(err.Error())
	}
	defer Close()

	fmt.Println("connected to mysql database")

	results, err := database.Query("SELECT * FROM TimeSeries")
	if err != nil {
		panic(err.Error())
	}

	for results.Next() {
		var point types.Point

		err = results.Scan(
			&point.Time,
			&point.Positive,
			&point.Negative,
			&point.Retweets,
			&point.Total,
		)
		if err != nil {
			panic(err.Error())
		}
		fmt.Println(point)
	}
}
