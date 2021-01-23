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
				"ID":         &graphql.Field{Type: graphql.Int},
				"Time":       &graphql.Field{Type: graphql.DateTime},
				"CreatedAt":  &graphql.Field{Type: graphql.DateTime},
				"Sentiment":  &graphql.Field{Type: graphql.String},
				"Confidence": &graphql.Field{Type: graphql.Int},
				"Text":       &graphql.Field{Type: graphql.String},
				"Username":   &graphql.Field{Type: graphql.String},
				"Link":       &graphql.Field{Type: graphql.String},
			},
		},
	)

	return tweetType
}
