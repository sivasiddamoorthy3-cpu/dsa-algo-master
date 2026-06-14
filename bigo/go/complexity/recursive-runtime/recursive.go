package main

import "fmt"

// When a recursive function triggers multiple recursive calls (branches), its runtime often scales exponentially.

// recursive executes a recursive branching operation.
// Time Complexity:  O(2^N)
// Space Complexity: O(N)

// Visualizing the Recursive Call Tree for F(4):
//                   F(4)
//                 /      \
//              F(3)      F(3)
//             /   \      /   \
//          F(2)  F(2)  F(2)  F(2)
//          /  \  /  \  /  \  /  \
//         F1  F1 F1 F1 F1 F1 F1 F1

// Counting the Total Nodes (Function Calls):
// *   **Level 0:** 2^0 = 1 node
// *   **Level 1:** 2^1 = 2 nodes
// *   **Level 2:** 2^2 = 4 nodes
// *   **Level 3:** 2^3 = 8 nodes
// *   **Level 4:** 2^4 = 16 nodes

// > [!IMPORTANT]
// > **Branching Formula:** When a recursive function branches B times per call and reaches a maximum depth of D, the runtime will scale at:
// O(B^D)

// #### Exponent Bases Matter:
// Unlike logs, the base of an exponent **does matter** and cannot be dropped. For example, O(8^N) is vastly different from O(2^N) because 8^N = (2^3)^N = 2^{3N} = 2^{2N} 2^N, which differs by a factor of 2^{2N} (which is not constant!).

// #### Recursive Space Complexity:
// Although there are O(2^N) total calls in the tree, only O(N) nodes exist on the call stack at any single moment. Therefore, the memory/space complexity is only O(N).

func recursive(n int) int {
	if n <= 1 {
		return 1
	}

	return recursive(n-1) + recursive(n-1)
}

func main() {
	r := recursive(4)
	fmt.Printf("Recursive response %v", r)
}
