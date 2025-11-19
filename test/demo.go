package main

import (
	"embed"
	"fmt"
	"sync"
	"time"
)

var templateFS embed.FS

func main() {
	var wg sync.WaitGroup

	for i := 1; i < 5; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()
			fmt.Println("Hello from a goroutine")
		}()
	}

	wg.Wait()

	fmt.Println("all goroutines finished!")

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
