package client

import (
	"fmt"
)

func InsertAnalyticsPing(identity string, ip string, request string) error {
	query := fmt.Sprintf("INSERT INTO Ping (identity, ip, request) VALUES (\"%s\",\"%s\",\"%s\")", identity, ip, request)
	_, err := analytics.Exec(query)
	return err
}
