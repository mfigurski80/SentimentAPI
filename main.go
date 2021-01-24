package main

import (
	"encoding/json"
	"fmt"

	"github.com/graphql-go/graphql"

	"github.com/mfigurski80/SentimentAPI/schema"
)

func main() {
	fmt.Println("vim-go")

}

func testSchema() {
	schema := schema.BuildSchema()

	query := `
		{
			points(from: 45, to: 60) {
				time
			}
		}
	`

	params := graphql.Params{Schema: schema, RequestString: query}
	r := graphql.Do(params)

	rJSON, _ := json.Marshal(r)
	fmt.Printf("%s \n", rJSON)
}
