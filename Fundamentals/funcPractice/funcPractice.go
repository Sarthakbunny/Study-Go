package funcPractice

import "fmt"

type transformFn func(int) int

// type anotherFn func(int, []string, map[string][]int) ([]int, string)

func FuncPractice() {
	array := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	doubledArray := transformArray(&array, double)
	tripleddArray := transformArray(&array, triple)

	fmt.Println(doubledArray)
	fmt.Println(tripleddArray)

	anotherArray := []int{3, 2, 1, 3, 2}

	tranformFunc1 := getTranformFunction(&array)
	tranformFunc2 := getTranformFunction(&anotherArray)

	tranformedArray := transformArray(&array, tranformFunc1)
	tranformedArray1 := transformArray(&anotherArray, tranformFunc2)

	fmt.Println(tranformedArray)
	fmt.Println(tranformedArray1)

	//anonymous function
	tranformArrayUsingAnonymousFunction := transformArray(&array, func(num int) int {
		return num * 2
	})
	fmt.Println(tranformArrayUsingAnonymousFunction)

	//closure call
	tFn1 := getClosureFunction(2)
	tFn2 := getClosureFunction(3)

	tArray1 := transformArray(&array, tFn1)
	tArray2 := transformArray(&array, tFn2)

	fmt.Println(tArray1)
	fmt.Println(tArray2)

	//Variadic func
	arr := []int{1, 2, 4, 3}
	sum := sumup(1, 2, 3, 2, 3, 2, 2, 4, 2)
	anotherSum := sumup(1, arr...)
	fmt.Println(sum)
	fmt.Println(anotherSum)
}

func getTranformFunction(numbers *[]int) transformFn {
	if (*numbers)[0] == 1 {
		return double
	}
	return triple
}

func transformArray(array *[]int, tranform transformFn) []int {
	tranformedArray := make([]int, len(*array))
	for index, value := range *array {
		tranformedArray[index] = tranform(value)
	}

	return tranformedArray
}

func double(val int) int {
	return val * 2
}

func triple(val int) int {
	return val * 3
}

// adding closure function
func getClosureFunction(factor int) transformFn {
	return func(value int) int {
		return value * factor
	}
}

// Variadic function
// ... means "collect all"
func sumup(startingValue int, rest ...int) int {
	sum := 0
	for _, value := range rest {
		sum += value
	}

	return sum - startingValue
}
