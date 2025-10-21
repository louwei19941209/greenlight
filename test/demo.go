package main

import "fmt"

func main() {

	S("")

}

func A(a *int, b int) {
	fmt.Println(*a + b)
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
