package main

import (
	"fmt"

	"github.com/mfigurski80/SentimentAPI/client"
)

func main() {
	fmt.Println("starting up api...")

	if err := client.Open(); err != nil {
		panic(err.Error())
	}

	rows, err := client.Execute("SELECT * FROM TimeSeries WHERE time > \"2021-02-24 00:00:00\"")
	if err != nil {
		panic(err.Error())
	}

	points := client.ReadOutPoints(rows)
	for _, v := range *points {
		fmt.Println(v)
	}

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
