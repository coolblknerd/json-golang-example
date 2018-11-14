package main

import (
	"encoding/json"
	"fmt"
)

type Human struct {
	FirstName string
	LastName  string
	Age       int
	Gender    string
}

func main() {
	human := Human{
		FirstName: "Jordan",
		LastName:  "Franklin",
		Age:       28,
		Gender:    "Female",
	}

	b, _ := json.Marshal(human)
	s := string(b)
	fmt.Println(s)
}
