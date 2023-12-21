package lib

import (
	"fmt"
	"strconv"
)

func FizzBuzz(target int) {
	for i := 1; i <= target; i++ {
		result := ""

		if i%2 == 0 {
			result += "fizz"
		}

		if i%5 == 0 {
			result += "buzz"
		}

		if result == "" {
			result = strconv.Itoa(i)
		}

		fmt.Println(result)
	}
}
