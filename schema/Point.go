package schema

import "github.com/graphql-go/graphql"

type Point struct {
	Time     int `json:"time"`
	Positive int `json:"positive"`
	Negative int `json:"negative"`
	Retweets int `json:"retweets"`
	Total    int `json:"total"`
}

func BuildPointType() *graphql.Object {
	var pointType = graphql.NewObject(
		graphql.ObjectConfig{
			Name: "Point",
			Fields: graphql.Fields{
				"Time":     &graphql.Field{Type: graphql.Int},
				"Positive": &graphql.Field{Type: graphql.Int},
				"Negative": &graphql.Field{Type: graphql.Int},
				"Retweets": &graphql.Field{Type: graphql.Int},
				"Total":    &graphql.Field{Type: graphql.Int},
			},
		},
	)

	return pointType
}
