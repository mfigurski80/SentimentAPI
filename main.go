package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/graphql-go/graphql"

	"github.com/mfigurski80/SentimentAPI/client"
	"github.com/mfigurski80/SentimentAPI/schema"
)

func main() {
	setupServer()
}

func setupServer() {
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

	http.HandleFunc("/graphql", func(w http.ResponseWriter, r *http.Request) {
		urlParams := r.URL.Query()
		identity := urlParams.Get("identity")
		if identity == "" {
			w.Write([]byte("{\"data\":null, \"errors\": [{\"message\", \"missing identity parameter\"}]}"))
			return
		}
		// TODO: log identity to mysql
		query := r.URL.Query().Get("query")
		result := graphql.Do(graphql.Params{
			Schema:        graphSchema,
			RequestString: query,
		})
		json.NewEncoder(w).Encode(result)
	})

	fmt.Println("server is running on 0.0.0.0:8080")
	http.ListenAndServe(":8080", nil)
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
