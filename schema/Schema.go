package schema

import (
	"log"

	"github.com/graphql-go/graphql"
)

func BuildSchema() graphql.Schema {

	rootQuery := BuildQueryFields()

	schemaConfig := graphql.SchemaConfig{Query: rootQuery}
	schema, err := graphql.NewSchema(schemaConfig)
	if err != nil {
		log.Fatalf("Failed to create graphql schema: %v", err)
	}

	return schema
}
