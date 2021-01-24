package schema

import (
	"fmt"

	"github.com/graphql-go/graphql"
)

func BuildQueryFields(pointsResolver func(from int, to int) ([]Point, error)) *graphql.Object {
	tweetType := BuildTweetType()
	pointType := BuildPointType(tweetType)

	queryFields := graphql.Fields{
		"points": &graphql.Field{
			Type:        graphql.NewList(pointType),
			Description: "Get TimeSeries points",
			Args: graphql.FieldConfigArgument{
				"from": &graphql.ArgumentConfig{Type: graphql.Int},
				"to":   &graphql.ArgumentConfig{Type: graphql.Int},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				from, ok1 := p.Args["from"].(int)
				to, ok2 := p.Args["to"].(int)
				if !ok1 || !ok2 {
					return nil, fmt.Errorf("Failed parsing args")
				}
				return pointsResolver(from, to)
			},
		},
	}

	return graphql.NewObject(graphql.ObjectConfig{Name: "RootQuery", Fields: queryFields})
}
