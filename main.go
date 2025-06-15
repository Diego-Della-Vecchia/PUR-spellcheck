package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"spellcheck/trees"
	"time"
)

func main() {

	words, err := readJSONFile("./dictionary.json")

	if err != nil {
		fmt.Printf("Error reading JSON file: %v\n", err)
		return
	}

	fmt.Println("Wörter geleson von dictionary.json:", len(words))

	tree := trees.Init(&trees.Node{
		Value: words[0], Children: make(map[int]*trees.Node),
	})

	tree.InsertList(words[1:])
	fmt.Println("Wörter in Baum eingefügt:", len(words))

	fmt.Println("Geben Sie ein Wort ein:")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	fmt.Println("Eingegebenes Wort:", input)
	results := tree.Lookup(input, 2)
	fmt.Println("Ähnliche Wörter:")

	for _, word := range results {
		fmt.Println(word)
		time.Sleep(200 * time.Millisecond) // Pause for 1 second
	}

}

func readJSONFile(filePath string) ([]string, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %w", err)
	}

	var result []string
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON array: %w", err)
	}

	return result, nil
}
