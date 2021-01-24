package schema

import (
	"github.com/graphql-go/graphql"
)

type Tweet struct {
	ID         int    `json:"id"`
	Time       int    `json:"time"`
	CreatedAt  int    `json:"createdAt"`
	Sentiment  string `json:"sentiment"`
	Confidence int    `json:"confidence"`
	Text       string `json:"text"`
	Username   string `json:"username"`
	Link       string `json:"link"`
}

func BuildTweetType() *graphql.Object {
	var tweetType = graphql.NewObject(
		graphql.ObjectConfig{
			Name: "Tweet",
			Fields: graphql.Fields{
				"id":         &graphql.Field{Type: graphql.Int},
				"time":       &graphql.Field{Type: graphql.DateTime},
				"createdAt":  &graphql.Field{Type: graphql.DateTime},
				"sentiment":  &graphql.Field{Type: graphql.String},
				"confidence": &graphql.Field{Type: graphql.Int},
				"text":       &graphql.Field{Type: graphql.String},
				"username":   &graphql.Field{Type: graphql.String},
				"link":       &graphql.Field{Type: graphql.String},
			},
		},
	)

	return tweetType
}
