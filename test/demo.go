package main

import (
	"fmt"
	"maps"
)

func main() {
	m1 := map[string]int{
		"one": 1,
		"two": 2,
	}

	m1s := maps.All(m1)

	for k, v := range m1s {
		fmt.Println(k, v)
	}
}
