package main

import (
	"fmt"
	"time"
)

func main() {

}

func A(a *int, b int) {
	start := time.Now()
	time.Sleep(3 * time.Second)
	fmt.Println(time.Since(start))
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
