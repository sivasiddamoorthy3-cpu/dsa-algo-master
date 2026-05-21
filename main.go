package main

import (
	"fmt"
)

var (
	rc      bool
	element int
	arr1    []int
	arr2    []int
)

func main() {
	fmt.Println("Hello Algos! Welcome to DSA Algo Master Program...")
	fmt.Println("=======================================")
	fmt.Println("Hello Algos! Welcome to BigO Complexity...")

	arr1 = []int{1, 3, 5, 7, 9, 11, 13, 15, 17, 19, 20, 21, 24, 26}

	arr2 = []int{22, 32, 42, 52, 62, 68, 72, 82, 92, 100}

	fmt.Println("Enter the value of the element to be searched: ")
	fmt.Scan(&element)

	rc = linearComplexity(arr1, element)
	if rc {
		fmt.Printf("Given element %v is found", element)
	} else {
		fmt.Printf("Given element %v is not found", element)
	}
}
