package main

import (
	"encoding/json"
	"fmt"
	"greenlight/internal/data"
)

func main() {
	var input struct {
		Title   string       `json:"title"`
		Year    int32        `json:"year"`
		Runtime data.Runtime `json:"runtime"`
		Genres  []string     `json:"genres"`
	}

	input.Title = "a"
	input.Year = 2002
	input.Year = 32
	input.Genres = []string{"a", "b"}

	marshal, _ := json.Marshal(input)
	fmt.Println(string(marshal))

}
