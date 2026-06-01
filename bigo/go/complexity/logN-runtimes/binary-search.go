// Algorithm to compute the logN runtimes from the given sorted array.

package main

// Input array arr = {1, 2, 3, 5, 6, 8, 9, 12, 15, 18} and target is 8
// Best case if the target is starting element, then it's time complexity is O(1).
// Average case if the target is in middle element, then it should be iterate half of the elements, then it's time complexity is O(n/2).
// Worst case if the target is in at the end element, then it should be iterate over based on the element's size.
// So, if the element size n, then it should be iterate n eements, so time complexity is O(n).
// Space Complexity is O(1).
func binarySearch(arr []int, target int) int {
	for i := 0; i < len(arr); i++ {
		if arr[i] == target {
			return i
		}
	}

	return -1
}

// Input array arr = {1, 5, 8, 9, 11, 13, 15, 19, 21} and target is 8
// We can optimize the time complexity from O(n) to O(logn) with mid value from low and high divided by half for every element iteration.
// That means if the element is in left side of the mid value then we will search it in left side only from next element iteration. so half of the elements we can ignore.
// else we will search it in right side only from the next element iteration, so half of the element we can ignore.
// example with explanation:
// [Searching for 9 in a 9-element array]
// Step 1:  {1, 5, 8, 9, 11, 13, 15, 19, 21}  --> mid = 11. Target 9 < 11. Discard right half.
// Step 2:  {1, 5, 8, 9}                      --> mid = 5 or 8. Target 9 > 8. Discard left half.
// Step 3:  {9}                               --> Found!

// The mathematical explanation for the above theory:
// n + n/2 + n/4 + ..... + 1
// let's say n=16, the it's 16 + 8 + 4 + 2 + 1, the same in reverse way 1 + 2 + 4 + 8 + 16 + ... + n
// What is k in the expression 2 the  power of k = N ? This is exactly what log expresses.
// 2^4 = 16 ==> log216 = 4  ====> k = log2N

// That is the mathematical calculation which is time complexity is O(logN).

func binarySearchOptimize(arr []int, target int) int {
	low := 0
	high := len(arr) - 1

	// Checking the whether the low is less than or equal to high in each iteration.
	for low <= high {
		mid := low + (high-low)/2 // Safe midpoint calculation
		// If the target is equal to arr[mid] value, then return the mid value.
		if target == arr[mid] {
			return mid
		} else if target > arr[mid] { // If the target value is greater than arr[mid] value, then the low value as increment +1 the mid value (mid + 1)
			low = mid + 1
		} else {
			high = mid - 1 // If the target value is less than arr[mid] value, then the high value as decrement -1 the low value (mid - 1)
		}
	}

	return -1
}
