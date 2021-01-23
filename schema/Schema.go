package schema

import (
	"fmt"
	"log"

	"github.com/graphql-go/graphql"
)

func BuildSchema() graphql.Schema {

	rootQuery := BuildQueryFields(func(p graphql.ResolveParams) (interface{}, error) {
		from, ok := p.Args["from"].(int)
		if !ok {
			return nil, fmt.Errorf("Failed reading args")
		}
		to, ok := p.Args["to"].(int)
		if !ok {
			return nil, fmt.Errorf("Failed reading args")
		}

		fmt.Println(from)
		fmt.Println(to)

		return nil, nil
	})

	schemaConfig := graphql.SchemaConfig{Query: rootQuery}
	schema, err := graphql.NewSchema(schemaConfig)
	if err != nil {
		log.Fatalf("Failed to create graphql schema: %v", err)
	}

	return schema
}
