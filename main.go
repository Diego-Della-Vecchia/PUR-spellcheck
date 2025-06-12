package main

import (
	"fmt"
	"os"
	"spellcheck/lev"
)

func main() {
	strings := os.Args[1:]

	var string1 string = strings[0]

	var string2 string = strings[1]

	fmt.Println(lev.Iterative(string1, string2))

	fmt.Println(lev.Recursive(string1, string2))

}
