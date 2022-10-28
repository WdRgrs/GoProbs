package main

import (
	"fmt"
	// ! Un-comment/Swap to account for inputs less than 0 & Error handling
	// "errors"
)

// ! Un-comment/Swap to account for inputs less than 0 & Error handling
// func BracketCombinations(num int) (int, error) {
func BracketCombinations(num int) int {
	// ! Un-comment/Swap to account for inputs less than 0 & Error handling
	// if num < 0 {
	// 	return 0, errors.New("Input must be a positive integer")
	// }

	if num <= 1 {
		// ! Un-comment/Swap to account for inputs less than 0 & Error handling
		// return num, nil
		return num
	}

	combinations := make([]string, 0)

	var backtrack func(str string, n, openNum, closeNum int)

	backtrack = func(str string, n, openNum, closeNum int) {
		// 'Open == n' is redundant and handled later but helps for code readability and DevX
		// This wouldn't be a normal comment, but worth calling out in the test - I prefer to have code more readable than clever
		if (openNum == n) && (closeNum == n) {
			combinations = append(combinations, str)
			return
		}

		if closeNum < openNum {
			str += ")"
			backtrack(str, n, openNum, closeNum+1)
		}

		if openNum < n {
			str += "("
			backtrack(str, n, openNum+1, closeNum)
		}
	}

	backtrack("", num, 0, 0)

	// ! Un-comment/Swap to account for inputs less than 0 & Error handling
	// return len(combinations), nil
	return len(combinations)
}

func main() {
	// do not modify below here, readline is our function
	// that properly reads in the input for you

	// ! Uncomment & change bracketNumber variable to desired test
	// bracketNumber := 32
	// fmt.Println(BracketCombinations(bracketNumber))

	fmt.Println(BracketCombinations(readline()))
}
