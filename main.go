package main

import (
	"encoding/json"
	"fmt"

	"github.com/graphql-go/graphql"

	"github.com/mfigurski80/SentimentAPI/client"
	"github.com/mfigurski80/SentimentAPI/schema"
)

func main() {
	fmt.Println("starting up api...")

	if err := client.Open(); err != nil {
		panic(err.Error())
	}
	defer client.Close()

	resolvers := schema.QueryResolverStruct{
		QueryPoint:  QueryPointResolver,
		QueryPoints: QueryPointsResolver,
		QueryTweets: QueryTweetsResolver,
		PointTweets: PointTweetsResolver,
	}

	graphSchema := schema.BuildSchema(resolvers)
	testSchema(&graphSchema)
}

func testSchema(graphSchema *graphql.Schema) {

	query := `query TimeSeriesQuery {
			points(from: 1611410000, to: 1611490000) {
				time
				total
			}
		}
	`

	params := graphql.Params{Schema: *graphSchema, RequestString: query}
	r := graphql.Do(params)

	rJSON, _ := json.Marshal(r)
	fmt.Printf("%s \n", rJSON)
}
