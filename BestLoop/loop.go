package main

import (
	"fmt"
	"strings"
)

func FindIntersection(strArr []string) string {
	// Split the argument array into seperate arrays and
	// initialize an empty array to append intersections into
	firstArr := strings.Split(strArr[0], ",")
	secondArr := strings.Split(strArr[1], ",")
	intersectionsArr := []string{}

	for i := 0; i < len(firstArr); i++ {
		for j := 0; j < len(secondArr); j++ {
			if firstArr[i] == secondArr[j] {
				intersectionsArr = append(intersectionsArr, firstArr[i])
			}
		}
	}

	// Join intersetions array into a string
	// and remove all whitespace as per output requirements
	output := strings.Join(intersectionsArr, ",")
	output = strings.ReplaceAll(output, " ", "")

	// If there are no intersections return false
	if len(output) == 0 {
		return "False"
	}

	return output
}

func main() {
	// do not modify below here, readline is our function
	// that properly reads in the input for you

	// ! Uncomment 41-42 to test in a local environtment
	// arr := []string{"1, 3, 9, 10, 17, 18", "1, 4, 9, 10"}
	// fmt.Println(FindIntersection(arr))

	fmt.Println(FindIntersection(readline()))
}
