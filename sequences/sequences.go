package main

import (
	"fmt"
)

func addOne(a int) int {
	return a + 1
}

func square(a int) int {
	return a * a
}

func double(slice []int) {
	slice = append(slice, slice...)
}

func mapSlice(f func(a int) int, slice []int) {
	for i, n := range slice {
		slice[i] = f(n)
	}
}

func mapArray(f func(a int) int, array [3]int) [3]int {
	a2 := [3]int{}
	for i, n := range array {
		a2[i] = f(n)
	}
	return a2
}

func main() {
	intsSlice := []int{1, 2, 3, 4, 5}

	mapSlice(addOne, intsSlice)

	fmt.Println(intsSlice)

	intsArray := [3]int{1, 2, 3}

	intsArray = mapArray(addOne, intsArray)

	fmt.Println(intsArray)

	newSlice := intsSlice[1:3]

	mapSlice(square, newSlice)

	double(intsSlice)

	fmt.Println(newSlice)

	fmt.Println(intsSlice)

}
