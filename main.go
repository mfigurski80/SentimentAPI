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
		// get url parameters
		urlParams := r.URL.Query()
		query := r.URL.Query().Get("query")
		identity := urlParams.Get("identity")

		// make sure identity is given
		if identity == "" {
			w.Write([]byte("{\"data\":null, \"errors\": [{\"message\", \"missing identity parameter\"}]}"))
			return
		}

		// return graphql result
		result := graphql.Do(graphql.Params{
			Schema:        graphSchema,
			RequestString: query,
		})
		json.NewEncoder(w).Encode(result)

		// write to analytics
		writeAnalytics(identity, query, r)
	})

	fmt.Println("server is running on 0.0.0.0:8080")
	http.ListenAndServe(":8080", nil)
}

func writeAnalytics(identity string, query string, r *http.Request) {
	ip := r.Header.Get("X-Real-Ip")
	if ip == "" {
		ip = r.Header.Get("X-Forwarded-For")
	}
	if ip == "" {
		ip = r.RemoteAddr
	}
	if err := client.InsertAnalyticsPing(identity, ip, ""); err != nil {
		panic(err)
	}
}
