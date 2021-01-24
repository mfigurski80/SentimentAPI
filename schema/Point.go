package schema

import "github.com/graphql-go/graphql"

type Point struct {
	Time     int     `json:"time"`
	Positive int     `json:"positive"`
	Negative int     `json:"negative"`
	Retweets int     `json:"retweets"`
	Total    int     `json:"total"`
	Tweets   []Tweet `json:"tweet"`
}

func BuildPointType(tweetType *graphql.Object) *graphql.Object {
	var pointType = graphql.NewObject(
		graphql.ObjectConfig{
			Name: "point",
			Fields: graphql.Fields{
				"time":     &graphql.Field{Type: graphql.Int},
				"positive": &graphql.Field{Type: graphql.Int},
				"negative": &graphql.Field{Type: graphql.Int},
				"retweets": &graphql.Field{Type: graphql.Int},
				"total":    &graphql.Field{Type: graphql.Int},
				"tweets":   &graphql.Field{Type: graphql.NewList(tweetType)},
			},
		},
	)

	return pointType
}
