package schema

import (
	"log"

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

type Point struct {
	Time     int     `json:"time"`
	Positive int     `json:"positive"`
	Negative int     `json:"negative"`
	Retweets int     `json:"retweets"`
	Total    int     `json:"total"`
	Tweets   []Tweet `json:"tweet"`
}

type QueryResolverStruct struct {
	QueryPoints func(from int, to int) ([]Point, error)
	QueryTweets func(at int) ([]Tweet, error)
	PointTweets func(Point) ([]Tweet, error)
}

func BuildSchema(res QueryResolverStruct) graphql.Schema {

	tweetType := graphql.NewObject(graphql.ObjectConfig{
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
					if point, ok := p.Source.(Point); ok {
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

			"points": &graphql.Field{
				Type:        graphql.NewList(pointType),
				Description: "Get TimeSeries points",
				Args: graphql.FieldConfigArgument{
					"from": &graphql.ArgumentConfig{Type: graphql.Int},
					"to":   &graphql.ArgumentConfig{Type: graphql.Int},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if from, ok := p.Args["from"].(int); ok {
						if to, ok := p.Args["to"].(int); ok {
							return res.QueryPoints(from, to)
						}
					}
					return nil, nil
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
						return res.QueryTweets(at)
					}
					return nil, nil
				},
			},
		},
	})

	schemaConfig := graphql.SchemaConfig{Query: queryType}
	schema, err := graphql.NewSchema(schemaConfig)
	if err != nil {
		log.Fatalf("Failed to create graphql schema: %v", err)
	}

	return schema
}
