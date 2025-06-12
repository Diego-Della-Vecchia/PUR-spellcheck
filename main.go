package main

import (
	"fmt"
	"spellcheck/trees"
)

func main() {

	t := trees.Init(&trees.Node{Value: "book", Children: make(map[int]*trees.Node)})
	t.Insert(&trees.Node{Value: "book", Children: make(map[int]*trees.Node)})
	t.Insert(&trees.Node{Value: "books", Children: make(map[int]*trees.Node)})
	t.Insert(&trees.Node{Value: "back", Children: make(map[int]*trees.Node)})
	t.Insert(&trees.Node{Value: "boon", Children: make(map[int]*trees.Node)})
	t.Insert(&trees.Node{Value: "Hallo", Children: make(map[int]*trees.Node)})

	results := t.Lookup("book", 1)

	fmt.Println(results)

	fmt.Println()
}
