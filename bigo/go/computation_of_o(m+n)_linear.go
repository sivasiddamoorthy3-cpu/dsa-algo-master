// Algorithm to compute of O(m+n) Linear time complexity

package bigo

// Linear Search Operation
// Input arr1 = []int{1, 3, 5, 7, 9, 11, 13, 15, 17, 19, 20, 21, 24, 26}
// Usually time complexity will be three types worst, average, best cases. Here we are going to discuss.
// 1. worst case: If the element(26) is at last postion, then it's worst, and the time complexity is O(n).
// 2. average case: if the element(13) is at middle position, then it's average, and the time complexity is O(n/2).
// 3. best case: if the element(1) is at starting position, then it's best, and the time complexity is O(1).

// So always we will consider to worst case and try to improve/optimize that only most of the times.
// In this case with n elements the linear search operation time complexity is O(n).
// Space complexity here we are not store and iterate anything, so i.e O(1).
func linearComputationOfOmn(arr1 []int, arr2 []int, element int) bool {

	// Iterate the arr1 to find out the given element
	for i := 0; i < len(arr1); i++ {
		if arr1[i] == element {
			return true
		}
	}

	// Iterate the arr2 to find out the given element
	for j := 0; j < len(arr2); j++ {
		if arr2[j] == element {
			return true
		}
	}

	// If no element found, return false
	return false
}
