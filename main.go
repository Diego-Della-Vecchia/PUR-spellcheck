package main

import (
	"fmt"
	"spellcheck/trees"
)

func main() {

	t := trees.Init(&trees.Node{Value: "book", Children: make(map[int]*trees.Node)})

	t.InsertList([]string{
		"book",
		"books",
		"back",
		"boon",
		"Hallo",
	})

	//Should return "book" and "books" and "boon"
	results := t.Lookup("book", 1)

	fmt.Println(results)

}
