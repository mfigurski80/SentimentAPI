package main

import (
	"fmt"

	"github.com/mfigurski80/SentimentAPI/schema"
)

func main() {
	fmt.Println("vim-go")
	schema := schema.BuildSchema()
	fmt.Println(schema)
}
