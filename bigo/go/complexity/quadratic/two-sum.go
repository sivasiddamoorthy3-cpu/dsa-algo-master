// Two sum problem for demonstation of quadratic time complexity

package main

import "fmt"

// Two Sum Operation
// Input arr1 := []int{0, 3, 6, 9, 12, 15}
// Here if the length is n elements, we need to iterate each element by with other elements, so there are below posibilities.
// Worst case: n(n-1)/2 = n2/2 -n/2, so here if we drop the constant and non-dominant ones then the complexity is n2, i.e O(n2)
// because, the first time through j runs n-1 steps, second time n-2 and third time n-3 and so on.
// Therefore, the number of steps total is: (n-1) + (n-2) + (n-3) + ... + 3 + 2 + 1 ==> 1 + 2 + 3 + ... + n-1
// The sum of 1 through n-1 is n(n-1)/2, so the runtime (Time Complexity) will be O(n2).
// Space Complexity is O(1).
func TwoSum(arr1 []int, expectedSum int) (bool, int, int) {

	for i := 0; i < len(arr1); i++ {
		for j := i + 1; j < len(arr1); j++ {
			if expectedSum == arr1[i]+arr1[j] {
				return true, i, j
			}
		}
	}

	return false, -1, -1
}

// Two Sum Otimization
// Input arr1 := []int{0, 3, 6, 9, 12, 15}
// Here if the length is n elements, we need to iterate each element and therefore we are checking with map of Key and Value pair, so there are below posibilities.
// Worst case:
// Time Complexity: O(n)
// Space Complexity: O(n), because everytime with the element we are checking with the map,
// Whether the key is present in map(we are minusing the element with expected sum), for n elements it should be iterate n times.
func TwoSumOptimize(arr1 []int, expextedSum int) (bool, int, int) {
	// Created the map key, value pair of int
	sumMap := make(map[int]int)
	for i := 0; i < len(arr1); i++ {
		// If the expectedSum pair is there in given array, then instead of iterating additional for loop, we can get the remaining element value based on the input element value.
		// Like finding the second element = output - first element.
		sumMapKey := expextedSum - arr1[i]
		// So with that second element we are checking with map as key, if it finds then return, else adding it into the map key as element and value as index.
		if val, ok := sumMap[sumMapKey]; ok {
			return true, i, val
		} else {
			sumMap[arr1[i]] = i
		}
	}
	return false, -1, -1
}

func main() {
	fmt.Println("Hello Algos! Welcome to BigO Complexity...")

	arr1 := []int{0, 3, 6, 9, 12, 15}

	expectedSum := 24
	fmt.Println("========================================")
	fmt.Println("Welcome to BigO Quadratic Complexity...")
	isFound, ind1, ind2 := TwoSum(arr1, expectedSum)
	if isFound {
		fmt.Printf("ExpectedSum:%d of %d and %d is found from indexes of arr1: %d and arr1: %d.\n", expectedSum, arr1[ind1], arr1[ind2], ind1, ind2)
	} else {
		fmt.Printf("ExpectedSum %d is not found.\n", expectedSum)
	}

	fmt.Println("Time Complexity for TwoSum: O(n2)")
	fmt.Println("Space Complexity for TwoSum: O(1)")
	fmt.Println("=========================================")

	fmt.Println("========================================")
	fmt.Println("Welcome to BigO Quadratic Optimize Complexity...")
	isFound, ind1, ind2 = TwoSumOptimize(arr1, expectedSum)
	if isFound {
		fmt.Printf("ExpectedSum:%d of %d and %d is found from indexes of arr1: %d and arr1: %d.\n", expectedSum, arr1[ind1], arr1[ind2], ind1, ind2)
	} else {
		fmt.Printf("ExpectedSum %d is not found.\n", expectedSum)
	}

	fmt.Println("Time Complexity for TwoSum Optimize: O(n)")
	fmt.Println("Space Complexity for TwoSum Optimize: O(n)")
	fmt.Println("=========================================")
}
