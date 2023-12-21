package main

import (
	"fmt"
	"gogo-exercise/lib"
	"os"
)

const (
	ERROR_TEXT = "Unknown action provided. Usage: Available actions: binary_search, thread_safe_counter"
)

func main() {
	if len(os.Args) < 2 {
		panic(ERROR_TEXT)
	}

	action := os.Args[1]

	switch action {
	case "binary_search":
		fmt.Println(lib.Find([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 5))
	case "thread_safe_counter":
		lib.Counter()
	case "fizzbuzz":
		lib.FizzBuzz(30)
	default:
		panic(ERROR_TEXT)
	}

}
