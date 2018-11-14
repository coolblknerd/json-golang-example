package main

import (
	"encoding/json"
	"fmt"
)

type Language struct {
	ID   int
	Name string
}

func main() {
	text := `[{"ID": 100, "Name": "Go"}, {"ID": 200, "Name": "Java"}]`
	bytes := []byte(text)

	var languages []Language
	json.Unmarshal(bytes, &languages)

	for l := range languages {
		fmt.Printf("Id = %v, Name = %v\n", languages[l].ID, languages[l].Name)
	}
}
