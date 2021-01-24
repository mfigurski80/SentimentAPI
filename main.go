package main

import (
	"fmt"

	"github.com/mfigurski80/SentimentAPI/client"
)

func main() {
	client.Start()
	fmt.Println("vim-go")
	// So we want something close to python, where we can define
	// @resolver('field.subfield')
	// def func():
	// in order to resolve a subfield
	// since we don't really need a whatsitcalled, higher order function
	// we can do this differently by having a resolver object passed down to schema builder?
}

// func testSchema() {
// schema := schema.BuildSchema()
//
// query := `
// {
// points(from: 45, to: 60) {
// time
// }
// }
// `
//
// params := graphql.Params{Schema: schema, RequestString: query}
// r := graphql.Do(params)
//
// rJSON, _ := json.Marshal(r)
// fmt.Printf("%s \n", rJSON)
// }
