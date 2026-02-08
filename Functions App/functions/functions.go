package functions

import (
	"fmt"
)

type transformFn func(int) int

func main() {
	numbers := []int{1, 2, 3, 4}
	newNumbers := []int{5, 6, 7}
	doubled := transformNumbers(&numbers, double)
	tripled := transformNumbers(&numbers, triple)

	fmt.Println(doubled)
	fmt.Println(tripled)

	transformerFn1 := getTransformerFn(&numbers)
	transformerFn2 := getTransformerFn(&newNumbers)

	transformedNumbers := transformNumbers(&numbers, transformerFn1)
	newTransformedNumbers := transformNumbers(&newNumbers, transformerFn2)

	fmt.Println(transformedNumbers)
	fmt.Println(newTransformedNumbers)

	// transformed := transformNumbers(&numbers, func(number int)int {
	// 	return number * 2
	// })
}

func transformNumbers(numbers *[]int, transform func(int) int) []int {
	dNum := []int{}

	for _, val := range *numbers {
		dNum = append(dNum, transform(val))
	}

	return dNum
}

func double(number int) int {
	return number * 2
}

func triple(number int) int {
	return number * 3
}

func getTransformerFn(numbers *[]int) transformFn {
	if (*numbers)[0] == 1 {
		return double
	} else {
		return triple
	}
}

func createTranformer(factor int) func(int) int {
	return func(number int) int {
		return number * factor
	}
}

func factorial(number int) int {
	if number == 0 {
		return 1
	}
	return number * factorial(number-1)
}
