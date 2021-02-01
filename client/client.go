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
var analytics *sql.DB

func parseConnectionString() string {
	username := os.Getenv("MYSQL_USERNAME")
	password := os.Getenv("MYSQL_PASSWORD")

	service := strings.ToUpper(os.Getenv("MYSQL_SERVICE"))
	service = strings.ReplaceAll(service, "-", "_")
	host := os.Getenv(service + "_SERVICE_HOST")
	port := os.Getenv(service + "_SERVICE_PORT")

	return fmt.Sprintf("%s:%s@tcp(%s:%s)", username, password, host, port)
}

func Open() error {
	connString := parseConnectionString()
	fmt.Println("connecting to mysql database at: " + connString)
	db, err := sql.Open("mysql", connString+"/sentiment")
	database = db
	if err != nil {
		return err
	}
	db, err = sql.Open("mysql", connString+"/analytics")
	analytics = db
	return err
}

func Close() {
	database.Close()
	analytics.Close()
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
