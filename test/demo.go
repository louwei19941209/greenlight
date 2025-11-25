package main

import (
	"embed"
	"fmt"
	"time"
)

var templateFS embed.FS

func main() {
	b := C()
	for _, k := range *b {
		fmt.Println(k)
	}

}

func A(a *int, b int) {
	start := time.Now()
	time.Sleep(3 * time.Second)
	fmt.Println(time.Since(start))

}

func B(i, k int) *bool {
	var b bool
	b = i > k
	return &b
}

func C() *[]int {
	a := 1
	b := 2
	c := 3
	var d = []int{a, b, c}
	return &d
}

func S(a string) {
	switch a {
	case "a":
		fmt.Println("a")
	case "b":
		fmt.Println("b")
	default:
		fmt.Println("d")
	}
}
