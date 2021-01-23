package schema

import "github.com/graphql-go/graphql"

func BuildQueryFields(fn func(graphql.ResolveParams) (interface{}, error)) *graphql.Object {
	tweetType := BuildTweetType()
	pointType := BuildPointType(tweetType)

	queryFields := graphql.Fields{
		"points": &graphql.Field{
			Type:        graphql.NewList(pointType),
			Description: "Get TimeSeries points",
			Args: graphql.FieldConfigArgument{
				"from": &graphql.ArgumentConfig{Type: graphql.DateTime},
				"to":   &graphql.ArgumentConfig{Type: graphql.DateTime},
			},
			Resolve: fn,
		},
	}

	return graphql.NewObject(graphql.ObjectConfig{Name: "RootQuery", Fields: queryFields})
}
