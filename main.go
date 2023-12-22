package main

import (
	"fmt"
	"gogo-exercise/lib"
	"os"
	"strconv"
)

const (
	ERROR_TEXT = "Unknown action provided. Usage: Available actions: binary_search, thread_safe_counter, fib, fizzbuzz, linked_list, blockchain"
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
	case "linked_list":
		myList := &lib.List{}

		err := myList.Insert("Hello")
		if err != nil {
			fmt.Println(err)
			return
		}
		myList.Insert("can")
		myList.Insert("you")
		myList.Insert("hear")
		myList.Insert("me")

		fmt.Println(myList)
		fmt.Println(myList.Len())
	case "fib":
		input, _ := strconv.Atoi(os.Args[2])
		fmt.Println(lib.Fib(input))
	case "blockchain":
		bc := &lib.BlockChain{}

		bc.AddBlock(234.5)
		bc.AddBlock(233.5)
		bc.AddBlock(322.5)
		bc.AddBlock(1.5)

		bc.DangerouslySetBlockAmount(2, 9999)

	default:
		panic(ERROR_TEXT)
	}

}
