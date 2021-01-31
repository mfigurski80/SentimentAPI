package schema

import (
	"fmt"
	"log"

	"github.com/graphql-go/graphql"
	"github.com/mfigurski80/SentimentAPI/types"
)

// QueryResolverStruct describes all resolvers required to complete schema
type QueryResolverStruct struct {
	QueryPoint  func(at int64) (types.Point, error)
	QueryPoints func(from int64, to int64) ([]types.Point, error)
	QueryTweets func(at int64) ([]types.Tweet, error)
	PointTweets func(types.Point) ([]types.Tweet, error)
}

func BuildSchema(res QueryResolverStruct) graphql.Schema {

	// query

	tweetType := graphql.NewObject(graphql.ObjectConfig{
		Name: "tweet",
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
	})

	pointType := graphql.NewObject(graphql.ObjectConfig{
		Name: "point",
		Fields: graphql.Fields{
			"time":     &graphql.Field{Type: graphql.Int},
			"positive": &graphql.Field{Type: graphql.Int},
			"negative": &graphql.Field{Type: graphql.Int},
			"retweets": &graphql.Field{Type: graphql.Int},
			"total":    &graphql.Field{Type: graphql.Int},
			"tweets": &graphql.Field{
				Type: graphql.NewList(tweetType),
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if point, ok := p.Source.(types.Point); ok {
						return res.PointTweets(point)
					}
					return nil, nil
				},
			},
		},
	})

	queryType := graphql.NewObject(graphql.ObjectConfig{
		Name: "RootQuery",
		Fields: graphql.Fields{
			"point": &graphql.Field{
				Type:        pointType,
				Description: "Get a specific TimeSeries point",
				Args: graphql.FieldConfigArgument{
					"at": &graphql.ArgumentConfig{Type: graphql.Int},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if at, ok := p.Args["at"].(int); ok {
						return res.QueryPoint(int64(at))
					}
					return nil, fmt.Errorf("couldn't parse args")
				},
			},
			"points": &graphql.Field{
				Type:        graphql.NewList(pointType),
				Description: "Get TimeSeries points between times",
				Args: graphql.FieldConfigArgument{
					"from": &graphql.ArgumentConfig{Type: graphql.Int},
					"to":   &graphql.ArgumentConfig{Type: graphql.Int},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if from, ok := p.Args["from"].(int); ok {
						if to, ok := p.Args["to"].(int); ok {
							return res.QueryPoints(int64(from), int64(to))
						}
					}
					return nil, fmt.Errorf("couldn't parse args")
				},
			},
			"tweets": &graphql.Field{
				Type:        graphql.NewList(tweetType),
				Description: "Get Tweets for specific point",
				Args: graphql.FieldConfigArgument{
					"at": &graphql.ArgumentConfig{Type: graphql.Int},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if at, ok := p.Args["at"].(int); ok {
						return res.QueryTweets(int64(at))
					}
					return nil, fmt.Errorf("couldn't parse args")
				},
			},
		},
	})

	// mutation

	jwtType := graphql.NewObject(graphql.ObjectConfig{
		Name: "jwt",
		Fields: graphql.Fields{
			"token": &graphql.Field{Type: graphql.String},
		},
	})

	mutationType := graphql.NewObject(graphql.ObjectConfig{
		Name: "RootMutation",
		Fields: graphql.Fields{
			"getToken": &graphql.Field{
				Type: jwtType,
				Args: graphql.FieldConfigArgument{
					"identity": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if identity, ok := p.Args["identity"].(string); ok {
						return CreateJWT(identity)
					}
					return nil, fmt.Errorf("couldn't parse args")
				},
			},
		},
	})

	schemaConfig := graphql.SchemaConfig{Query: queryType, Mutation: mutationType}
	schema, err := graphql.NewSchema(schemaConfig)
	if err != nil {
		log.Fatalf("Failed to create graphql schema: %v", err)
	}

	return schema
}
