package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"spellcheck/trees"
	"time"
)

var ascii = `
 _                             _   _          _         ____  _     _                       
| |    _____   _____ _ __  ___| |_| |__   ___(_)_ __   |  _ \(_)___| |_ __ _ _ __   ___ ___ 
| |   / _ \ \ / / _ \ '_ \/ __| __| '_ \ / _ \ | '_ \  | | | | / __| __/ _' | '_ \ / __/ _ \
| |__|  __/\ V /  __/ | | \__ \ |_| | | |  __/ | | | | | |_| | \__ \ || (_| | | | | (_|  __/
|_____\___| \_/ \___|_| |_|___/\__|_| |_|\___|_|_| |_| |____/|_|___/\__\__,_|_| |_|\___\___|
`

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

	for i := 0; i >= 0; i++ {

		fmt.Println(ascii)

		fmt.Println("Wörter in Baum eingefügt:", len(words))
		if i == 0 {
			fmt.Println("Geben Sie ein Wort ein:")
		} else {
			fmt.Println("Geben Sie ein anderes Wort ein:")
		}
		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		fmt.Println("Eingegebenes Wort:", input)
		results := tree.Lookup(input, 2)
		fmt.Println("Ähnliche Wörter:")

		for _, word := range results {
			fmt.Println(word)
			time.Sleep(200 * time.Millisecond)
		}
		time.Sleep(5 * time.Second)

		clearTerminal()

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
func clearTerminal() {
	var cmd *exec.Cmd

	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("cmd", "/c", "cls")
	default: // Unix-like (Linux, macOS)
		cmd = exec.Command("clear")
	}

	cmd.Stdout = os.Stdout
	cmd.Run()
}
