# Chapter VI: Big O (Cracking the Coding Interview, Pages 38–47)

Big O time is the primary language and metric used to describe the efficiency and scalability of algorithms. Understanding Big O thoroughly is essential for designing high-performance code, evaluating algorithm optimizations, and succeeding in technical interviews.

---

## 1. An Analogy: Electronic vs. Airplane Transfer (Page 38)

To understand asymptotic runtime, consider sending a file to a friend across the country as fast as possible:

*   **Electronic Transfer:** You send it via email, FTP, or cloud sharing. 
    *   **Time Complexity:** $\mathcal{O}(s)$, where $s$ is the size of the file. As the file size increases, the transfer time scales linearly.
*   **Airplane Transfer:** You copy the file to a physical hard drive and fly it across the country.
    *   **Time Complexity:** $\mathcal{O}(1)$ with respect to file size. Whether the file is 1 MB or 10 TB, the travel time remains constant.

> [!NOTE]
> Even if flying across the country takes 10 hours (a massive constant delay), the constant time algorithm $\mathcal{O}(1)$ will always outperform the linear time algorithm $\mathcal{O}(s)$ at a certain file size threshold (e.g., a multi-terabyte file).

```
   Time
     ^
     |       / [Linear Transfer: O(s)]
     |      /
     |     /
     |    /
     |   /  . - - - - - - - - - - - [Airplane Transfer: O(1)]
     |  /  .
     | /   .
     +---------------------------------> File Size (s)
```

---

## 2. Time Complexity & Big O, Big Theta, and Big Omega (Page 38–39)

In academic computer science, three distinct notations describe the bounds of a runtime:

| Notation | Academic Definition | Analogy | Industry Usage (Interviews) |
| :--- | :--- | :--- | :--- |
| **$\mathcal{O}$ (Big O)** | **Upper Bound** (tells us that an algorithm is at least as fast as this. Similar to $\le$). | If age is $X$, $X \le 130$ (also $X \le 1000$). | Merged with $\Theta$. Refers to the **tight upper bound** (i.e. we say printing an array is $\mathcal{O}(N)$ rather than $\mathcal{O}(N^2)$). |
| **$\Omega$ (Big Omega)** | **Lower Bound** (tells us that an algorithm won't be faster than this. Similar to $\ge$). | If age is $X$, $X \ge 0$ (also $X \ge \log X$). | Rarely used explicitly unless discussing theoretical lower limits. |
| **$\Theta$ (Big Theta)** | **Tight Bound** (both $\mathcal{O}$ and $\Omega$, meaning a precise match). | If age is $X$, then $X \approx 30$ precisely. | What industry actually means when they ask for "Big O". |

---

## 3. Best, Worst, and Expected Case (Page 40)

An algorithm's runtime can vary significantly based on the input structure. We describe these variations in three ways:

*   **Best Case:** The absolute most favorable input scenario. (Rarely used in practice because we can easily hardcode specific $\mathcal{O}(1)$ shortcuts for specific inputs).
*   **Worst Case:** The absolute worst possible input configuration.
*   **Expected Case:** The average or typical behavior under randomized inputs.

### Quick Sort Analysis Case Study

Quick Sort works by choosing a pivot element, partitioning elements around it, and recursing:

```
[Equal / Ideal Pivot]                 [Reverse Sorted Array / Poor Pivot]
      [  5  ]                                     [  9  ]
     /       \                                   /
  [2]         [8]                             [8]
 /   \       /   \                           /
[1]  [3]   [6]   [9]                       [7]
```

*   **Best Case $\mathcal{O}(N)$:** Occurs if all elements are equal (with smart partitioning) or pre-sorted with optimal pivots, taking a single pass.
*   **Worst Case $\mathcal{O}(N^2)$:** Occurs if the pivot is consistently the maximum/minimum element (e.g. reverse-sorted arrays with the first element as the pivot). The recursion tree shrinks by only 1 element per level rather than splitting in half.
*   **Expected Case $\mathcal{O}(N \log N)$:** Average case runtime on randomized input arrays.

> [!WARNING]
> Do not muddle **Best/Worst/Expected case** with **Big O/Omega/Theta**. 
> *   *Best/Worst/Expected* describe the **inputs**.
> *   *Big O/Omega/Theta* describe the **mathematical bounds** (upper/lower/tight) for those inputs.

---

## 4. Space Complexity (Page 40–41)

Memory (or space) complexity is the amount of memory required relative to the input size. This includes both heap allocations and **call stack space** in recursive algorithms.

### Recursive Stack Memory: $\mathcal{O}(N)$ Space

In recursion, each active call adds a stack frame to memory.

```go
package complexity

// Sum calculates the sum of integers from 1 to n recursively.
// Time Complexity:  O(n)
// Space Complexity: O(n) - due to n active stack frames
func Sum(n int) int {
	if n <= 0 {
		return 0
	}
	return n + Sum(n-1)
}
```

#### Visualizing the Active Stack Frames:
```
Call Stack:
[ Sum(4) ] -> needs Sum(3)
  [ Sum(3) ] -> needs Sum(2)
    [ Sum(2) ] -> needs Sum(1)
      [ Sum(1) ] -> needs Sum(0)
        [ Sum(0) ] -> Returns 0 (Max Depth: 5 frames active simultaneously)
```

---

### Iterative Call Stack: $\mathcal{O}(1)$ Space

Even if an algorithm makes $N$ total calls, if they are not active at the same time, the stack depth remains $\mathcal{O}(1)$.

```go
package complexity

// PairSumSequence computes the sum of sequential pairs.
// Time Complexity:  O(n)
// Space Complexity: O(1) - calls execute and return immediately, 
//                          never nesting on the stack.
func PairSumSequence(n int) int {
	sum := 0
	for i := 0; i < n; i++ {
		sum += pairSum(i, i+1)
	}
	return sum
}

func pairSum(a, b int) int {
	return a + b
}
```

---

## 5. Drop the Constants (Page 41)

Big O only describes the **rate of increase** as the input grows toward infinity. It does not measure the exact execution time. Therefore, we drop all multiplicative constants.

An algorithm with two independent loops is still $\mathcal{O}(N)$, not $\mathcal{O}(2N)$.

### Case Study: Single vs. Double Loop (Golang)

Here are two Go implementations to find the minimum and maximum values in a slice:

```go
package complexity

import "math"

// MinMaxOne finds min and max in a single loop.
// Conceptually: O(N) time with 2 operations per iteration.
func MinMaxOne(array []int) (int, int) {
	minVal := math.MaxInt
	maxVal := math.MinInt

	for _, x := range array {
		if x < minVal {
			minVal = x
		}
		if x > maxVal {
			maxVal = x
		}
	}
	return minVal, maxVal
}

// MinMaxTwo finds min and max in two separate sequential loops.
// Conceptually: O(2N) time with 1 operation per loop iteration.
// This is STILL O(N) asymptotically!
func MinMaxTwo(array []int) (int, int) {
	minVal := math.MaxInt
	maxVal := math.MinInt

	for _, x := range array {
		if x < minVal {
			minVal = x
		}
	}
	for _, x := range array {
		if x > maxVal {
			maxVal = x
		}
	}
	return minVal, maxVal
}
```

> [!TIP]
> Do not attempt to count assembly-level instructions or cycles to determine if $\mathcal{O}(2N)$ is faster than $\mathcal{O}(N)$. Modern compiler optimizations, CPU branch prediction, and cache hits make absolute instruction counting highly unreliable. Focus purely on asymptotic growth.

---

## 6. Drop the Non-Dominant Terms (Page 42)

When expressing runtime complexities with multiple components, always drop the terms that grow slower as the input size $N$ becomes extremely large.

*   $\mathcal{O}(N^2 + N) \rightarrow \mathcal{O}(N^2)$ (Since $N^2$ dominates $N$)
*   $\mathcal{O}(N + \log N) \rightarrow \mathcal{O}(N)$
*   $\mathcal{O}(5 \cdot 2^N + 1000N^{100}) \rightarrow \mathcal{O}(2^N)$ (Exponential dominates polynomial)

---

## 7. Multi-Part Algorithms: Add vs. Multiply (Page 42–43)

Suppose you have an algorithm with two steps. Determining whether to add or multiply the runtimes depends on the flow of execution.

### Rule 1: Add the Runtimes — $\mathcal{O}(A + B)$
If the algorithm is of the form: *"Do this, then when you are all done, do that."*

```go
// AddRuntimes completes A steps, then completes B steps sequentially.
// Complexity: O(A + B)
func AddRuntimes(arrA []int, arrB []int) {
	for _, a := range arrA {
		fmt.Println(a)
	}
	for _, b := range arrB {
		fmt.Println(b)
	}
}
```

### Rule 2: Multiply the Runtimes — $\mathcal{O}(A \cdot B)$
If the algorithm is of the form: *"Do this for each time you do that."*

```go
// MultiplyRuntimes nestedly prints elements.
// Complexity: O(A * B)
func MultiplyRuntimes(arrA []int, arrB []int) {
	for _, a := range arrA {
		for _, b := range arrB {
			fmt.Printf("%d,%d\n", a, b)
		}
	}
}
```

> [!IMPORTANT]
> **Multi-Variable Runtimes:** If the inputs represent different, independent data variables, they **cannot** be simplified or dropped:
> *   $\mathcal{O}(B^2 + A)$ cannot be reduced because we do not know the relationship between $A$ and $B$.

---

## 8. Amortized Time (Page 43)

Amortized time is a method of analyzing the average performance of an operation over a sequence of actions, specifically when a rare, expensive "worst-case" operation is guaranteed to be compensated by a long sequence of very cheap "best-case" operations.

### Case Study: Dynamic Resizing Arrays (`ArrayList` / Go Slices)
A Go slice or Java `ArrayList` is backed by a fixed-size array. When full, inserting a new element requires:
1. Allocating a new array of double the size ($2N$).
2. Copying all existing $N$ elements into the new array.
3. Appending the new element.

#### Deriving the Amortized Runtime:
If we insert $X$ elements into a dynamic array starting from size 1, capacity doubling events occur at sizes:
$$1, 2, 4, 8, 16, \dots, X$$

The number of copies required at each resize event is:
$$1 + 2 + 4 + 8 + 16 + 32 + \dots + X$$

Reading this sum **from right to left**:
$$X + \frac{X}{2} + \frac{X}{4} + \frac{X}{8} + \dots + 1 \approx 2X$$

*   **Total Work for $X$ insertions:** $\mathcal{O}(2X)$ copies/operations.
*   **Amortized Work per insertion:** $\frac{\mathcal{O}(2X)}{X} = \mathcal{O}(1)$ constant time.

---

## 9. $\mathcal{O}(\log N)$ Runtimes (Page 44)

$\mathcal{O}(\log N)$ complexities appear when the problem space is halved at each step of the algorithm.

### Case Study: Binary Search
In a sorted array, we find a target by continuously checking the midpoint and discarding the half that cannot contain the target.

```go
// BinarySearch searches for a target in a sorted slice.
// Time Complexity:  O(log N)
// Space Complexity: O(1)
func BinarySearch(array []int, target int) int {
	low := 0
	high := len(array) - 1

	for low <= high {
		mid := low + (high-low)/2 // Safe midpoint calculation
		if array[mid] == target {
			return mid
		} else if array[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}
```

```
[Searching for 9 in a 9-element array]
Step 1:  {1, 5, 8, 9, 11, 13, 15, 19, 21}  --> mid = 11. Target 9 < 11. Discard right half.
Step 2:  {1, 5, 8, 9}                      --> mid = 5 or 8. Target 9 > 8. Discard left half.
Step 3:  {9}                               --> Found!
```

> [!TIP]
> **Halving Rule:** When you see a problem where the number of elements in the problem space gets halved at each step, it will likely scale at a logarithmic $\mathcal{O}(\log N)$ runtime.

#### Base of Logs Note:
In Big O notation, the base of the logarithm does not matter (e.g., $\log_2 N$ is asymptotically equivalent to $\log_{10} N$) because change-of-base rules differ only by a constant factor, which we drop.

---

## 10. Recursive Branching Runtimes: $\mathcal{O}(2^N)$ (Page 44–45)

When a recursive function triggers multiple recursive calls (branches), its runtime often scales exponentially.

```go
// F executes a recursive branching operation.
// Time Complexity:  O(2^N)
// Space Complexity: O(N)
func F(n int) int {
	if n <= 1 {
		return 1
	}
	return F(n-1) + F(n-1)
}
```

### Visualizing the Recursive Call Tree for `F(4)`:

```
                  F(4)
                /      \
             F(3)      F(3)
            /   \      /   \
         F(2)  F(2)  F(2)  F(2)
         /  \  /  \  /  \  /  \
        F1  F1 F1 F1 F1 F1 F1 F1
```

### Counting the Total Nodes (Function Calls):
*   **Level 0:** $2^0 = 1$ node
*   **Level 1:** $2^1 = 2$ nodes
*   **Level 2:** $2^2 = 4$ nodes
*   **Level 3:** $2^3 = 8$ nodes
*   **Level 4:** $2^4 = 16$ nodes

Total calls for size $N$ is the sum of powers of 2 up to $N$:
$$\sum_{i=0}^{N} 2^i = 2^{N+1} - 1 \text{ nodes} \rightarrow \mathcal{O}(2^N)$$

> [!IMPORTANT]
> **Branching Formula:** When a recursive function branches $B$ times per call and reaches a maximum depth of $D$, the runtime will scale at:
> $$\mathcal{O}(B^D)$$

#### Exponent Bases Matter:
Unlike logs, the base of an exponent **does matter** and cannot be dropped. For example, $\mathcal{O}(8^N)$ is vastly different from $\mathcal{O}(2^N)$ because $8^N = (2^3)^N = 2^{3N} = 2^{2N} \cdot 2^N$, which differs by a factor of $2^{2N}$ (which is not constant!).

#### Recursive Space Complexity:
Although there are $\mathcal{O}(2^N)$ total calls in the tree, only $\mathcal{O}(N)$ nodes exist on the call stack at any single moment. Therefore, the memory/space complexity is only **$\mathcal{O}(N)$**.

---

## 11. Practical Examples and Exercises (Page 46–47)

### Example 1: Multiple Sequential Loops — $\mathcal{O}(N)$
```go
// ExampleOne runs two sequential loops.
// Complexity: O(N) time, O(1) space.
func ExampleOne(array []int) {
	sum := 0
	product := 1
	for i := 0; i < len(array); i++ {
		sum += array[i]
	}
	for i := 0; i < len(array); i++ {
		product *= array[i]
	}
	fmt.Printf("%d, %d\n", sum, product)
}
```

### Example 2: Nested Loops — $\mathcal{O}(N^2)$
```go
// ExampleTwo prints all pairs in an array.
// Complexity: O(N^2) time, O(1) space.
func ExampleTwo(array []int) {
	for i := 0; i < len(array); i++ {
		for j := 0; j < len(array); j++ {
			fmt.Printf("%d, %d\n", array[i], array[j])
		}
	}
}
```

### Example 3: Nested Loops Starting at `i + 1` — $\mathcal{O}(N^2)$
```go
// ExampleThree prints unique unordered pairs in an array.
// Complexity: O(N^2) time, O(1) space.
func ExampleThree(array []int) {
	for i := 0; i < len(array); i++ {
		for j := i + 1; j < len(array); j++ {
			fmt.Printf("%d, %d\n", array[i], array[j])
		}
	}
}
```

#### Three Derivation Methods for Unordered Pairs:
1. **Mathematical Sequence:** 
   The inner loop runs $N-1$ times, then $N-2$, then $N-3$, down to 1 time.
   $$\text{Total steps} = (N-1) + (N-2) + \dots + 2 + 1 = \frac{N(N-1)}{2} = \frac{N^2}{2} - \frac{N}{2} \rightarrow \mathcal{O}(N^2)$$
2. **Visual Representation (Matrix Half):**
   The $(i, j)$ iterations form a triangle covering exactly half of an $N \times N$ matrix.
   $$\text{Total area} = \frac{N^2}{2} \rightarrow \mathcal{O}(N^2)$$
   
   ```
   (0,1) (0,2) (0,3) (0,4)
         (1,2) (1,3) (1,4)
               (2,3) (2,4)
                     (3,4)
   ```
3. **Average Work Method:**
   The outer loop runs $N$ times. The inner loop length starts at $N$ and decreases to 0. Its average iteration size is $\frac{N}{2}$.
   $$\text{Total work} = N \cdot \frac{N}{2} = \frac{N^2}{2} \rightarrow \mathcal{O}(N^2)$$

### Example 4: Unordered Pairs from Two Different Arrays — $\mathcal{O}(A \cdot B)$
```go
// ExampleFour prints unordered pairs using two different slices.
// Complexity: O(A * B) time, where A = len(arrayA) and B = len(arrayB).
//             Space Complexity is O(1).
func ExampleFour(arrayA []int, arrayB []int) {
	for i := 0; i < len(arrayA); i++ {
		for j := 0; j < len(arrayB); j++ {
			if arrayA[i] < arrayB[j] {
				fmt.Printf("%d,%d\n", arrayA[i], arrayB[j])
			}
		}
	}
}
```

---

## 12. Rate of Growth Graph & Runtimes (Page 42)

The rate of increase determines how well an algorithm will perform when scaled to large datasets.

```
   Complexity (Time)
     ^
     |         / O(x!)        / O(2^x)      / O(x^2)
     |        /              /             /
     |       /              /             /      / O(x log x)
     |      /              /             /      /
     |     /              /             /      /      / O(x)
     |    /              /             /      /      /
     |   /              /             /      /      /
     |  /              /             /      /      /
     | /              /             /      /      /  _ _ _ _ O(log x)
     +-------------------------------------------------------------> Input Size (x)
```

### Typical Growth Orders (Fastest/Best to Slowest/Worst):
1.  **$\mathcal{O}(1)$** — Constant Time (e.g., hash map lookup)
2.  **$\mathcal{O}(\log N)$** — Logarithmic Time (e.g., binary search)
3.  **$\mathcal{O}(N)$** — Linear Time (e.g., finding min/max)
4.  **$\mathcal{O}(N \log N)$** — Linearithmic Time (e.g., Merge Sort, Heap Sort)
5.  **$\mathcal{O}(N^2)$** — Quadratic Time (e.g., Bubble Sort, slow nested loops)
6.  **$\mathcal{O}(2^N)$** — Exponential Time (e.g., recursive Fibonacci)
7.  **$\mathcal{O}(N!)$** — Factorial Time (e.g., traveling salesperson via brute-force)

---

## 13. More Examples and Exercises (Pages 48–52)

### Example 5: Triple Nested Loop with a Constant Inner Loop — $\mathcal{O}(AB)$ (Page 48)

Despite three nested loops, the innermost loop runs a **fixed** 100,000 iterations regardless of input size, making it a constant $\mathcal{O}(1)$ contribution. Only the outer two loops drive growth.

```go
// ExampleFive prints unordered pairs from two arrays with a fixed inner loop.
// Time Complexity:  O(A * B)  — the k-loop adds only a constant factor
// Space Complexity: O(1)
func ExampleFive(arrayA []int, arrayB []int) {
	for i := 0; i < len(arrayA); i++ {
		for j := 0; j < len(arrayB); j++ {
			for k := 0; k < 100000; k++ { // constant — does NOT change complexity
				fmt.Printf("%d,%d\n", arrayA[i], arrayB[j])
			}
		}
	}
}
```

> [!NOTE]
> 100,000 units of work is still **constant**. Therefore the runtime remains $\mathcal{O}(AB)$, not $\mathcal{O}(100000 \cdot AB)$.

---

### Example 6: Array Reversal — $\mathcal{O}(N)$ (Page 48)

The loop runs only to the midpoint of the array, but we still call this $\mathcal{O}(N)$ because Big O drops constants.

```go
// ExampleSix reverses a slice in-place.
// Time Complexity:  O(N) — iterates through N/2 elements, constant drops.
// Space Complexity: O(1)
func ExampleSix(array []int) {
	for i := 0; i < len(array)/2; i++ {
		other := len(array) - i - 1
		array[i], array[other] = array[other], array[i]
	}
}
```

---

### Example 7: Equivalent $\mathcal{O}(N)$ Expressions (Page 48)

Which of the following are equivalent to $\mathcal{O}(N)$ and why?

| Expression | Equivalent to $\mathcal{O}(N)$? | Reason |
| :--- | :---: | :--- |
| $\mathcal{O}(N + P)$ where $P < \frac{N}{2}$ | ✅ Yes | Since $P < \frac{N}{2}$, $N$ is the dominant term; we drop $P$. |
| $\mathcal{O}(2N)$ | ✅ Yes | Constants are always dropped; $\mathcal{O}(2N) = \mathcal{O}(N)$. |
| $\mathcal{O}(N + \log N)$ | ✅ Yes | $N$ dominates $\log N$; we drop the non-dominant term. |
| $\mathcal{O}(N + M)$ | ❌ No | No known relationship between $N$ and $M$; both must be kept. |

---

### Example 8: Sorting an Array of Strings — $\mathcal{O}(a \cdot s(\log a + \log s))$ (Page 48–49)

Suppose you sort each string in an array and then sort the full array of strings.

**Setting up clear variable names:**
*   Let `s` = length of the **longest** string.
*   Let `a` = length of the **array** (number of strings).

**Step-by-step runtime analysis:**

| Step | Operation | Runtime | Reason |
| :--- | :--- | :--- | :--- |
| 1 | Sort each string | $\mathcal{O}(s \log s)$ | Sorting a string of length $s$ |
| 2 | Do this for all `a` strings | $\mathcal{O}(a \cdot s \log s)$ | Repeat for every string in the array |
| 3 | Sort the array of `a` strings | $\mathcal{O}(a \log a)$ comparisons × $\mathcal{O}(s)$ per comparison | Each string comparison is $\mathcal{O}(s)$ |
| — | Total | $\mathcal{O}(a \cdot s(\log a + \log s))$ | Add both parts |

> [!IMPORTANT]
> Using $N$ ambiguously for both string count and string length is a **common and costly mistake** in interviews. Always define separate, meaningful variables to avoid $\mathcal{O}(N^2 \log N)$ style errors. The final answer $\mathcal{O}(a \cdot s(\log a + \log s))$ cannot be reduced further.

```go
import "sort"

// ExampleEight sorts each string then sorts the array of strings.
// Time Complexity:  O(a * s * (log a + log s))
//                  where a = number of strings, s = length of longest string
// Space Complexity: O(a * s) for intermediate storage
func ExampleEight(strs []string) []string {
	// Step 1: Sort characters within each string — O(a * s log s)
	sorted := make([]string, len(strs))
	for i, s := range strs {
		runes := []rune(s)
		sort.Slice(runes, func(a, b int) bool { return runes[a] < runes[b] })
		sorted[i] = string(runes)
	}

	// Step 2: Sort the array of strings — O(a * s * log a)
	// (each comparison of two strings costs O(s))
	sort.Strings(sorted)
	return sorted
}
```

---

### Example 9: Sum of All Nodes in a Balanced Binary Search Tree — $\mathcal{O}(N)$ (Page 49)

```go
// TreeNode represents a node in a binary search tree.
type TreeNode struct {
	Value int
	Left  *TreeNode
	Right *TreeNode
}

// ExampleNine returns the sum of all node values in a BST.
// Time Complexity:  O(N) — every node is visited exactly once
// Space Complexity: O(log N) for a balanced tree (max call stack depth = tree height)
func ExampleNine(node *TreeNode) int {
	if node == nil {
		return 0
	}
	return ExampleNine(node.Left) + node.Value + ExampleNine(node.Right)
}
```

**Two ways to derive the $\mathcal{O}(N)$ runtime:**

1. **What It Means:** The code touches each node in the tree exactly once and does a constant-time operation (addition) per node. If there are $N$ nodes, the runtime is $\mathcal{O}(N)$.

2. **Recursive Pattern:** Using the branching formula $\mathcal{O}(\text{branches}^{\text{depth}})$, we have 2 branches per call and depth $\log N$ (for a balanced BST). This gives $\mathcal{O}(2^{\log N})$.

   Simplifying: let $P = 2^{\log_2 N}$.
   $$\log_2 P = \log_2 N \implies P = N$$
   $$\therefore 2^{\log N} = N \implies \mathcal{O}(N)$$

> [!CAUTION]
> Just because it's a binary search tree doesn't mean there's a logarithm in the runtime! When **all** nodes must be visited, the complexity is $\mathcal{O}(N)$, not $\mathcal{O}(\log N)$.

---

### Example 10: Primality Check — $\mathcal{O}(\sqrt{N})$ (Pages 50–51)

A number $n$ is prime only if it has no divisors between 2 and $\sqrt{n}$. If $n$ is divisible by a number greater than $\sqrt{n}$, then its counterpart divisor must be **less** than $\sqrt{n}$ — meaning we'd have already found it.

For example: 33 is divisible by 11 (> $\sqrt{33} \approx 5.7$), but 11's counterpart is 3 (< $\sqrt{33}$), which we check first.

```go
import "math"

// IsPrime checks if n is a prime number.
// Time Complexity:  O(sqrt(N)) — the loop runs from 2 to sqrt(n)
// Space Complexity: O(1)
func IsPrime(n int) bool {
	if n < 2 {
		return false
	}
	sqrtN := int(math.Sqrt(float64(n)))
	for x := 2; x <= sqrtN; x++ {
		if n%x == 0 {
			return false
		}
	}
	return true
}
```

> [!TIP]
> The loop condition `x * x <= n` (or equivalently `x <= sqrt(n)`) is the key insight. We stop early because any composite factor beyond $\sqrt{n}$ must have a paired factor below $\sqrt{n}$ that would already have been found.

---

### Example 11: Factorial — $\mathcal{O}(N)$ (Page 51)

A straightforward linear recursion descending from $N$ down to 1.

```go
// Factorial computes n! recursively.
// Time Complexity:  O(N) — one recursive call per step, N steps total
// Space Complexity: O(N) — N stack frames on the call stack simultaneously
func Factorial(n int) int {
	if n < 0 {
		return -1 // undefined
	}
	if n == 0 {
		return 1
	}
	return n * Factorial(n-1)
}
```

---

### Example 12: All Permutations of a String — $\mathcal{O}(N^2 \cdot N!)$ (Pages 51–52)

This is a tricky one. We can analyze by asking two questions: **How many times is `permutation` called?** and **How long does each call take?**

```go
// Permutation prints all permutations of a string.
// Time Complexity:  O(n^2 * n!)
// Space Complexity: O(n^2) — n! permutations each of length n on the call stack
func Permutation(str string) {
	permutationHelper(str, "")
}

func permutationHelper(str string, prefix string) {
	if len(str) == 0 {
		fmt.Println(prefix) // O(n) work per base case — printing n characters
	} else {
		for i := 0; i < len(str); i++ {
			// Remove char at index i, append to prefix
			rem := str[:i] + str[i+1:]         // O(n) string concatenation
			permutationHelper(rem, prefix+string(str[i])) // O(n) prefix build
		}
	}
}
```

**Deriving the runtime:**

| Question | Answer | Reason |
| :--- | :--- | :--- |
| How many times does it reach the base case? | $N!$ times | There are $N!$ permutations total. |
| How many internal nodes are in the call tree? | $\le N \cdot N!$ nodes | Each leaf has a path of length $N$; at most $N \times N!$ internal calls. |
| How long does each call take? | $\mathcal{O}(N)$ | String concatenation of `rem + prefix + char` takes $\mathcal{O}(N)$ each call. |
| **Total runtime** | $\mathcal{O}(N^2 \cdot N!)$ | $N \cdot N!$ nodes × $N$ work per node |

> [!NOTE]
> Through more complex mathematics, a tighter bound can be derived, but $\mathcal{O}(N^2 \cdot N!)$ is a valid and acceptable upper bound for interview purposes.

---

### Example 13: Nth Fibonacci Number — $\mathcal{O}(2^N)$ (Page 52)

```go
// Fib computes the Nth Fibonacci number recursively.
// Time Complexity:  O(2^N) — 2 branches per call, depth N
// Space Complexity: O(N)   — maximum call stack depth is N
func Fib(n int) int {
	if n <= 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return Fib(n-1) + Fib(n-2)
}
```

Using the branching pattern: there are **2 branches** per recursive call and the tree goes **as deep as N**, giving:
$$\mathcal{O}(\text{branches}^{\text{depth}}) = \mathcal{O}(2^N)$$

> [!TIP]
> Through more complex math, the tighter bound is actually $\mathcal{O}(1.6^N)$ because at the bottom of the call stack, single-call nodes exist (the `Fib(n-2)` branch terminates sooner). However, $\mathcal{O}(2^N)$ is the standard and acceptable interview answer.

---

## Guide: How to Execute These Go Code Examples

To run and verify any of the code examples provided below, you can set up a simple `main.go` file inside this directory:

1. Create a file named `main.go` and copy-paste the package declaration and imports:
   ```go
   package main

   import (
       "fmt"
       "sort"
   )
   ```
2. Paste the function you wish to test (e.g., `PowersOf2` or `Intersection`).
3. Add a `main()` function to call and print results:
   ```go
   func main() {
       // Example: Running Example 16
       fmt.Println("--- Running PowersOf2(50) ---")
       PowersOf2(50)

       // Example: Running Additional Problem VI.12
       fmt.Println("\n--- Running Intersection ---")
       a := []int{5, 2, 9}
       b := []int{9, 1, 5, 3}
       fmt.Printf("Intersection count: %d\n", Intersection(a, b)) // Expected: 2 (5 and 9)
   }
   ```
4. Execute it directly from your terminal:
   ```bash
   go run main.go
   ```

---

## 14. Fibonacci Complexity — Recursive vs. Memoized (Pages 53–54)

An algorithm that makes multiple recursive calls often results in exponential runtime. However, when we apply optimization techniques like memoization (caching), we can drastically reduce the complexity to linear time.

### Example 14: Fibonacci Numbers Without Memoization (Page 53)

This implementation prints all Fibonacci numbers from $0$ to $n$ sequentially, demonstrating the slow, raw recursion.

```go
package main

import "fmt"

// AllFib prints all Fibonacci numbers from 0 to n.
// Time Complexity:  O(2^n)
// Space Complexity: O(n)
func AllFib(n int) {
	for i := 0; i < n; i++ {
		// Call standard exponential Fibonacci for every number
		fmt.Printf("Fib(%d) = %d\n", i, Fib(i))
	}
}

func Fib(n int) int {
	if n <= 0 {
		return 0
	} else if n == 1 {
		return 1
	}
	// Makes two recursive branches per call, creating an O(2^n) branching tree
	return Fib(n-1) + Fib(n-2)
}
```

#### Time Complexity Analysis
A common pitfall is to conclude that since `Fib(n)` takes $\mathcal{O}(2^n)$ time and it is called $n$ times, the total runtime is $\mathcal{O}(n \cdot 2^n)$. 

**This logic is incorrect because the input $n$ changes for each call to `Fib(i)`:**
*   `Fib(1)` $\rightarrow 2^1$ steps
*   `Fib(2)` $\rightarrow 2^2$ steps
*   `Fib(3)` $\rightarrow 2^3$ steps
*   `Fib(4)` $\rightarrow 2^4$ steps
*   $\dots$
*   `Fib(n)` $\rightarrow 2^n$ steps

The total amount of work is the sum of these steps:
$$\text{Total steps} = 2^1 + 2^2 + 2^3 + 2^4 + \dots + 2^n = 2^{n+1} - 2$$

> [!NOTE]
> Since we drop constants and low-order terms, $\mathcal{O}(2^{n+1} - 2) \rightarrow \mathcal{O}(2^n)$. The total runtime to compute the first $n$ Fibonacci numbers recursively is still **$\mathcal{O}(2^n)$**.

---

### Example 15: Fibonacci Numbers With Memoization (Pages 53–54)

By caching (or memoizing) previously computed values in a slice, we can avoid recalculating identical sub-problems.

```go
package main

import "fmt"

// AllFibMemo prints all Fibonacci numbers from 0 to n using memoization.
// Time Complexity:  O(n)
// Space Complexity: O(n)
func AllFibMemo(n int) {
	// Create a slice to cache computed Fibonacci numbers (indexed 0 through n)
	memo := make([]int, n+1)
	for i := 0; i < n; i++ {
		fmt.Printf("Fib(%d) = %d\n", i, FibMemo(i, memo))
	}
}

func FibMemo(n int, memo []int) int {
	if n <= 0 {
		return 0
	} else if n == 1 {
		return 1
	} else if memo[n] > 0 {
		// Return the cached value immediately if it's already computed!
		return memo[n]
	}

	// Calculate and store the result in the cache slice before returning
	memo[n] = FibMemo(n-1, memo) + FibMemo(n-2, memo)
	return memo[n]
}
```

#### Walkthrough of the Memoization Process:
Let's trace how the recursive calls resolve sequentially for `AllFibMemo(5)`:
```
Loop executes for i = 0 to 4:

i = 1: FibMemo(1) -> returns 1 (Base case)
i = 2: FibMemo(2)
       -> FibMemo(1) -> returns 1 (Base case)
       -> FibMemo(0) -> returns 0 (Base case)
       -> memo[2] stores (1 + 0) = 1. Returns 1.
i = 3: FibMemo(3)
       -> FibMemo(2) -> CACHE HIT: memo[2] is 1. Returns 1 immediately without recursing!
       -> FibMemo(1) -> returns 1 (Base case)
       -> memo[3] stores (1 + 1) = 2. Returns 2.
i = 4: FibMemo(4)
       -> FibMemo(3) -> CACHE HIT: memo[3] is 2. Returns 2 immediately!
       -> FibMemo(2) -> CACHE HIT: memo[2] is 1. Returns 1 immediately!
       -> memo[4] stores (2 + 1) = 3. Returns 3.
i = 5: FibMemo(5)
       -> FibMemo(4) -> CACHE HIT: memo[4] is 3. Returns 3 immediately!
       -> FibMemo(3) -> CACHE HIT: memo[3] is 2. Returns 2 immediately!
       -> memo[5] stores (3 + 2) = 5. Returns 5.
```

> [!TIP]
> **Why it is $\mathcal{O}(n)$ time:** At each call to `FibMemo(i)`, we have already computed and stored the values for `FibMemo(i-1)` and `FibMemo(i-2)` in previous loop iterations. Consequently, the call stack does a simple constant-time lookup $\mathcal{O}(1)$, sums them, stores the new result, and returns. Since we do a constant amount of work $n$ times, the runtime is **$\mathcal{O}(n)$**.

---

## 15. Logarithmic Powers of 2 (Pages 54–55)

Logarithmic time complexities are highly efficient because the problem size is dramatically halved at each step.

### Example 16: Printing Powers of 2 from 1 to N (Pages 54–55)

This function prints all powers of 2 from 1 through $n$ (inclusive) via successive halving.

```go
package main

import "fmt"

// PowersOf2 prints powers of 2 from 1 through n (inclusive).
// Time Complexity:  O(log n)
// Space Complexity: O(log n)
func PowersOf2(n int) int {
	if n < 1 {
		return 0
	} else if n == 1 {
		fmt.Println(1)
		return 1
	} else {
		// Halve n on every recursive call
		prev := PowersOf2(n / 2)
		curr := prev * 2
		fmt.Println(curr)
		return curr
	}
}
```

We can derive the $\mathcal{O}(\log n)$ runtime in three distinct ways:

#### 1. What It Does (Execution Trace)
Let's walk through the call stack when invoking `PowersOf2(50)`:
```
PowersOf2(50)
  -> PowersOf2(25)
    -> PowersOf2(12)
      -> PowersOf2(6)
        -> PowersOf2(3)
          -> PowersOf2(1)
            -> print & return 1
          print & return 2
        print & return 4
      print & return 8
    print & return 16
  print & return 32
```
The runtime is the number of times we can divide $n$ by $2$ until we reach the base case of $1$. The number of times we can halve $n$ before reaching $1$ is **$\mathcal{O}(\log n)$**.

#### 2. What It Means (Output Counting)
The code is designed to print all powers of 2 from 1 through $n$. Each call prints exactly one value and returns. Therefore, the number of function calls (which represents our runtime) is directly equal to the number of powers of 2 between 1 and $n$.
Since there are exactly $\log_2 n$ powers of 2 in that range, the runtime is **$\mathcal{O}(\log n)$**.

#### 3. Rate of Increase
Consider how the runtime grows as $n$ increases:
*   If $n$ grows from $P$ to $P+1$, the number of calls might not change at all.
*   The number of calls only increases by exactly $1$ **each time $n$ doubles in size**.

Thus, the number of calls is the number of times we must double 1 to reach $n$. Mathematically, this is the value $x$ in the equation:
$$2^x = n \implies x = \log_2 n$$

Therefore, the runtime is **$\mathcal{O}(\log n)$**.

---

## 16. Additional Problems (Pages 55–58)

Here are detailed Go translations and runtime analyses for the chapter's practice problems.

### VI.1: Product of $a$ and $b$ (Page 55)
Computes the product of two positive integers via repeated addition.
```go
// Product computes the product of a and b.
// Time Complexity:  O(b) — loop runs exactly b times.
// Space Complexity: O(1) — constant space.
func Product(a, b int) int {
	sum := 0
	for i := 0; i < b; i++ {
		sum += a
	}
	return sum
}
```
*   **Concrete Trace `Product(5, 4)`:**
    *   `i = 0`: `sum = 0 + 5` (5)
    *   `i = 1`: `sum = 5 + 5` (10)
    *   `i = 2`: `sum = 10 + 5` (15)
    *   `i = 3`: `sum = 15 + 5` (20)
    *   Loop stops. Returns 20.
*   **Analysis:** The loop iterates exactly `b` times. Each iteration performs a constant-time addition. Therefore, the runtime scales linearly with `b`, i.e., **$\mathcal{O}(b)$**.

---

### VI.2: $a^b$ (Page 55)
Computes $a^b$ recursively by multiplying $a$ by $a^{b-1}$.
```go
// Power computes a^b recursively.
// Time Complexity:  O(b) — recursive calls descend from b to 0.
// Space Complexity: O(b) — requires b call stack frames in memory.
func Power(a, b int) int {
	if b < 0 {
		return 0 // error
	} else if b == 0 {
		return 1
	} else {
		// recursive step: decrements b by 1
		return a * Power(a, b-1)
	}
}
```
*   **Concrete Trace `Power(3, 3)`:**
    *   `Power(3, 3)` $\rightarrow$ returns `3 * Power(3, 2)`
    *   `Power(3, 2)` $\rightarrow$ returns `3 * Power(3, 1)`
    *   `Power(3, 1)` $\rightarrow$ returns `3 * Power(3, 0)`
    *   `Power(3, 0)` $\rightarrow$ returns `1` (Base case)
    *   Total stack frames = 4 active.
*   **Analysis:** The recursion starts at `b` and decrements by 1 at each call down to the base case `b = 0`, making a total of `b` recursive calls. Each call performs constant-time work. This results in **$\mathcal{O}(b)$** time and **$\mathcal{O}(b)$** space due to the active stack frames.

---

### VI.3: $a \% b$ (Page 56)
Computes the remainder of $a$ divided by $b$.
```go
// Mod computes a % b.
// Time Complexity:  O(1) — constant time arithmetic operations.
// Space Complexity: O(1) — constant space.
func Mod(a, b int) int {
	if b <= 0 {
		return -1
	}
	div := a / b
	return a - div*b
}
```
*   **Concrete Trace `Mod(14, 3)`:**
    *   `div = 14 / 3 = 4`
    *   `result = 14 - (4 * 3) = 14 - 12 = 2`
*   **Analysis:** It performs a fixed number of basic arithmetic operations (division, multiplication, and subtraction) which execute in constant time regardless of the values of $a$ or $b$. The runtime is **$\mathcal{O}(1)$**.

---

### VI.4: Integer Division (Page 56)
Performs integer division of $a$ by $b$ using repeated subtraction.
```go
// Div performs integer division of a by b.
// Time Complexity:  O(a/b) — loop runs a/b times.
// Space Complexity: O(1)   — constant space.
func Div(a, b int) int {
	count := 0
	sum := b
	for sum <= a {
		sum += b
		count++
	}
	return count
}
```
*   **Concrete Trace `Div(10, 3)`:**
    *   `sum = 3`, `count = 1`
    *   `sum = 6`, `count = 2`
    *   `sum = 9`, `count = 3`
    *   `sum = 12` (exceeds 10). Loop stops. Returns 3.
*   **Analysis:** The loop adds `b` to `sum` repeatedly until `sum` is greater than `a`. The loop will run exactly $\lfloor a/b \rfloor$ times. Thus, the time complexity is **$\mathcal{O}(a/b)$**.

---

### VI.5: Square Root via Binary Search (Page 56)
Computes the integer square root of a number recursively using binary search (successive guessing).
```go
// SqrtBinarySearch computes the integer square root of n using binary search.
// Time Complexity:  O(log n) — search space is halved at each recursion step.
// Space Complexity: O(log n) — recursive call stack depth is log n.
func SqrtBinarySearch(n int) int {
	return sqrtHelper(n, 1, n)
}

func sqrtHelper(n, min, max int) int {
	if max < min {
		return -1 // no integer square root
	}

	guess := (min + max) / 2
	square := guess * guess

	if square == n {
		return guess // found it!
	} else if square < n {
		// Search the higher half of the remaining interval
		return sqrtHelper(n, guess+1, max) 
	} else {
		// Search the lower half of the remaining interval
		return sqrtHelper(n, min, guess-1) 
	}
}
```
*   **Concrete Trace `SqrtBinarySearch(16)`:**
    *   `sqrtHelper(16, 1, 16)` $\rightarrow$ `guess = 8`, `8 * 8 = 64 > 16` $\rightarrow$ calls `sqrtHelper(16, 1, 7)`
    *   `sqrtHelper(16, 1, 7)` $\rightarrow$ `guess = 4`, `4 * 4 = 16 == 16` $\rightarrow$ returns 4 (Found!)
*   **Analysis:** This search operates over the interval $[1, n]$. At each recursive call, the algorithm computes the midpoint `guess` and checks if `guess * guess` equals `n`. If not, it discards half of the remaining interval. Halving the search space at each step results in a logarithmic time complexity of **$\mathcal{O}(\log n)$** and a space complexity of **$\mathcal{O}(\log n)$** for the recursive stack.

---

### VI.6: Square Root via Linear Search (Page 56)
Computes the integer square root of a number by scanning numbers sequentially.
```go
// SqrtLinearSearch computes the integer square root of n using linear search.
// Time Complexity:  O(sqrt(n)) — loop runs until guess * guess > n.
// Space Complexity: O(1)       — constant space.
func SqrtLinearSearch(n int) int {
	for guess := 1; guess*guess <= n; guess++ {
		if guess*guess == n {
			return guess
		}
	}
	return -1
}
```
*   **Concrete Trace `SqrtLinearSearch(16)`:**
    *   `guess = 1`: `1 <= 16` $\rightarrow$ false
    *   `guess = 2`: `4 <= 16` $\rightarrow$ false
    *   `guess = 3`: `9 <= 16` $\rightarrow$ false
    *   `guess = 4`: `16 <= 16` $\rightarrow$ true (Returns 4)
*   **Analysis:** The loop continues as long as `guess * guess <= n`, which is mathematically equivalent to `guess <= sqrt(n)`. The loop runs at most $\sqrt{n}$ times, making the time complexity **$\mathcal{O}(\sqrt{n})$**.

---

### VI.7: Worst-Case Search in an Unbalanced BST (Page 57)
*If a binary search tree is not balanced, how long might it take (worst case) to find an element in it?*

> [!CAUTION]
> If a Binary Search Tree (BST) is completely unbalanced (e.g., all nodes only have right children), the tree structure collapses into a straight line, which is functionally equivalent to a singly linked list.

```
       [ 1 ]
         \
         [ 2 ]
           \
           [ 3 ]
             \
             [ 4 ] (Depth = N, lookup for 4 requires traversing all nodes)
```

**Analysis:** In the worst-case scenario (finding the leaf or a non-existent element at the end of the line), you must traverse every single node from the root down to the bottom. Thus, the worst-case lookup runtime is **$\mathcal{O}(n)$**, where $n$ is the number of nodes in the tree.

---

### VI.8: Search in an Ordinary Binary Tree (Page 57)
*You are looking for a specific value in a binary tree, but the tree is not a binary search tree. What is the time complexity of this?*

**Analysis:** Because the tree is not sorted (it lacks the BST property), you cannot discard any branches based on comparisons at the current node. You have no way of knowing whether the target lies in the left subtree or the right subtree. You are forced to perform a full traversal (such as DFS or BFS) and inspect every single node. The time complexity is **$\mathcal{O}(n)$**, where $n$ is the number of nodes.

---

### VI.9: Repeated Element Appends (Page 57)
Copies an array by repeatedly appending elements to a new array that expands by 1 slot at a time.
```go
// CopyArray copies an array by repeatedly appending elements to a new array of size + 1.
// Time Complexity:  O(n^2) — quadratic element copying.
// Space Complexity: O(n)   — space for the newly copied slice.
func CopyArray(array []int) []int {
	copyArr := []int{}
	for _, value := range array {
		copyArr = AppendToNew(copyArr, value)
	}
	return copyArr
}

func AppendToNew(array []int, value int) []int {
	// copy all elements over to new slice of length + 1
	bigger := make([]int, len(array)+1)
	for i := 0; i < len(array); i++ {
		bigger[i] = array[i]
	}

	// add new element
	bigger[len(bigger)-1] = value
	return bigger
}
```
*   **Concrete Trace `CopyArray([9, 8, 7])` (Length N = 3):**
    *   `value = 9`: `AppendToNew([], 9)` $\rightarrow$ copies `0` elements $\rightarrow$ copyArr is `[9]`
    *   `value = 8`: `AppendToNew([9], 8)` $\rightarrow$ copies `1` element $\rightarrow$ copyArr is `[9, 8]`
    *   `value = 7`: `AppendToNew([9, 8], 7)` $\rightarrow$ copies `2` elements $\rightarrow$ copyArr is `[9, 8, 7]`
    *   Total copies = $0 + 1 + 2 = 3$ operations.
*   **Analysis:** Let's count the total number of copy operations required:
    1.  The 1st call to `AppendToNew` copies $0$ elements.
    2.  The 2nd call copies $1$ element.
    3.  The 3rd call copies $2$ elements.
    4.  $\dots$
    5.  The $n$-th call copies $n-1$ elements.

    The total number of copy operations is the sum of integers from $0$ to $n-1$:
    $$\text{Total copies} = \sum_{i=0}^{n-1} i = \frac{n(n-1)}{2} = \frac{n^2}{2} - \frac{n}{2}$$

    Dropping non-dominant terms and constant factors, the time complexity is **$\mathcal{O}(n^2)$**.

---

### VI.10: Sum of Digits (Page 57)
Sums all the individual digits of a number.
```go
// SumDigits computes the sum of all digits in a number.
// Time Complexity:  O(log n) — specifically base-10 logarithm.
// Space Complexity: O(1)     — constant space.
func SumDigits(n int) int {
	sum := 0
	for n > 0 {
		sum += n % 10
		n /= 10
	}
	return sum
}
```
*   **Concrete Trace `SumDigits(324)`:**
    *   `n = 324`: `sum = 0 + 4 = 4`, `n = 32`
    *   `n = 32`: `sum = 4 + 2 = 6`, `n = 3`
    *   `n = 3`: `sum = 6 + 3 = 9`, `n = 0` (Loop terminates. Returns 9.)
*   **Analysis:** In each iteration, `n` is divided by 10. The loop runs exactly as many times as there are digits in `n`.
    *   A number with $d$ digits can have a value up to $10^d$.
    *   If $n = 10^d$, then taking the base-10 logarithm of both sides gives $d = \log_{10} n$.
    *   Since the number of loop iterations is directly equal to the number of digits $d$, the runtime scales logarithmically.

    Therefore, the time complexity is **$\mathcal{O}(\log n)$**.

---

### VI.11: Generating Sorted Strings of Length K (Page 57)
Prints all strings of length $k$ where the characters are in sorted order by generating all possible combinations and verifying them.
```go
const numChars = 26

// PrintSortedStrings prints all sorted strings of length k.
// Time Complexity:  O(k * c^k) where c is the character set size (26).
// Space Complexity: O(k)       — due to recursion depth of k.
func PrintSortedStrings(remaining int) {
	printSortedStringsHelper(remaining, "")
}

func printSortedStringsHelper(remaining int, prefix string) {
	if remaining == 0 {
		if isInOrder(prefix) {
			fmt.Println(prefix)
		}
	} else {
		for i := 0; i < numChars; i++ {
			c := ithLetter(i)
			printSortedStringsHelper(remaining-1, prefix+string(c))
		}
	}
}

func isInOrder(s string) bool {
	for i := 1; i < len(s); i++ {
		prev := s[i-1]
		curr := s[i]
		if prev > curr {
			return false
		}
	}
	return true
}

func ithLetter(i int) byte {
	return 'a' + byte(i)
}
```
**Analysis:**
1.  **Count of generated strings:** We have a choice of 26 characters at each of the $k$ positions, branching $26$ times at each level of the recursion tree. This generates a total of $26^k$ strings.
2.  **Work per string:** When the recursion reaches the base case (`remaining == 0`), it calls `isInOrder` on the string of length $k$. `isInOrder` iterates through the string of length $k$, performing $k-1$ comparisons, which takes $\mathcal{O}(k)$ time.
3.  **Total runtime:** $26^k$ total string leaves multiplied by $\mathcal{O}(k)$ work per leaf results in **$\mathcal{O}(k \cdot 26^k)$** (or more generally $\mathcal{O}(k \cdot c^k)$ where $c$ is the character set size).
4.  **Space complexity:** The maximum recursion stack depth is $k$. Thus, the space complexity is **$\mathcal{O}(k)$**.

---

### VI.12: Intersection of Two Arrays (Page 58)
Computes the number of elements in common between two arrays containing no duplicate elements.

```go
import "sort"

// Intersection computes the number of elements in common between two slices.
// It assumes neither slice contains duplicates.
// Time Complexity:  O(b log b + a log b) — sort array b, then binary search each x in a.
// Space Complexity: O(1) or O(b) depending on sorting implementation.
func Intersection(a []int, b []int) int {
	// Step 1: Sort the array b in ascending order
	sort.Ints(b) // O(b log b)
	intersect := 0

	// Step 2: For each element in a, run binary search inside sorted b
	for _, x := range a { // a iterations
		if binarySearch(b, x) >= 0 { // O(log b) per iteration
			intersect++
		}
	}
	return intersect
}

func binarySearch(arr []int, x int) int {
	low := 0
	high := len(arr) - 1
	for low <= high {
		mid := low + (high-low)/2
		if arr[mid] == x {
			return mid
		} else if arr[mid] < x {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}
```

*   **Concrete Trace `Intersection(a, b)`:**
    *   Input: `a = [5, 2, 9]`, `b = [9, 1, 5, 3]`
    *   Step 1: Sort `b` $\rightarrow$ `b` becomes `[1, 3, 5, 9]` (Cost: $\mathcal{O}(b \log b)$)
    *   Step 2: Loop through `a`:
        *   `x = 5`: `binarySearch([1, 3, 5, 9], 5)` $\rightarrow$ found (index 2) $\rightarrow$ `intersect = 1`
        *   `x = 2`: `binarySearch([1, 3, 5, 9], 2)` $\rightarrow$ not found (index -1)
        *   `x = 9`: `binarySearch([1, 3, 5, 9], 9)` $\rightarrow$ found (index 3) $\rightarrow$ `intersect = 2`
    *   Returns 2.

**Analysis:**
1.  **Sorting:** Sorting the second array `b` takes **$\mathcal{O}(b \log b)$** time using standard comparison-based sorting (e.g., merge sort or quicksort).
2.  **Searching:** We iterate through each of the `a` elements in the first array. For each element, we perform a binary search inside the sorted array `b`, which takes $\mathcal{O}(\log b)$ time. Doing this `a` times takes **$\mathcal{O}(a \log b)$** time.
3.  **Total Runtime:** Adding both parts together, the total runtime is **$\mathcal{O}(b \log b + a \log b)$** time.

---

## Summary Table: All Examples (Pages 46–55)

| Example | Description | Time Complexity | Space Complexity |
| :---: | :--- | :---: | :---: |
| 1 | Two sequential loops over same array | $\mathcal{O}(N)$ | $\mathcal{O}(1)$ |
| 2 | All pairs from same array (nested loops) | $\mathcal{O}(N^2)$ | $\mathcal{O}(1)$ |
| 3 | Unordered pairs from same array (`j = i+1`) | $\mathcal{O}(N^2)$ | $\mathcal{O}(1)$ |
| 4 | Unordered pairs from two different arrays | $\mathcal{O}(AB)$ | $\mathcal{O}(1)$ |
| 5 | Two arrays + constant inner loop (100k) | $\mathcal{O}(AB)$ | $\mathcal{O}(1)$ |
| 6 | Reverse array (loop to midpoint) | $\mathcal{O}(N)$ | $\mathcal{O}(1)$ |
| 7 | $\mathcal{O}(N+P)$ where $P < \frac{N}{2}$ | $\mathcal{O}(N)$ | — |
| 8 | Sort each string then sort array | $\mathcal{O}(as(\log a + \log s))$ | $\mathcal{O}(as)$ |
| 9 | Sum all nodes in a balanced BST | $\mathcal{O}(N)$ | $\mathcal{O}(\log N)$ |
| 10 | Primality check | $\mathcal{O}(\sqrt{N})$ | $\mathcal{O}(1)$ |
| 11 | Factorial (recursive) | $\mathcal{O}(N)$ | $\mathcal{O}(N)$ |
| 12 | All permutations of a string | $\mathcal{O}(N^2 \cdot N!)$ | $\mathcal{O}(N^2)$ |
| 13 | Nth Fibonacci (recursive) | $\mathcal{O}(2^N)$ | $\mathcal{O}(N)$ |
| 14 | Print all Fibonacci numbers 0 to N (recursive) | $\mathcal{O}(2^N)$ | $\mathcal{O}(N)$ |
| 15 | Print all Fibonacci numbers 0 to N (memoized) | $\mathcal{O}(N)$ | $\mathcal{O}(N)$ |
| 16 | Print powers of 2 from 1 to N | $\mathcal{O}(\log N)$ | $\mathcal{O}(\log N)$ |

## Summary Table: Additional Problems (Pages 55–58)

| Problem | Description | Time Complexity | Space Complexity |
| :---: | :--- | :---: | :---: |
| **VI.1** | Compute product of $a$ and $b$ | $\mathcal{O}(b)$ | $\mathcal{O}(1)$ |
| **VI.2** | Compute $a^b$ recursively | $\mathcal{O}(b)$ | $\mathcal{O}(b)$ |
| **VI.3** | Compute $a \% b$ | $\mathcal{O}(1)$ | $\mathcal{O}(1)$ |
| **VI.4** | Integer division $a / b$ via sequential subtraction | $\mathcal{O}(a/b)$ | $\mathcal{O}(1)$ |
| **VI.5** | Integer square root via Binary Search | $\mathcal{O}(\log n)$ | $\mathcal{O}(\log n)$ |
| **VI.6** | Integer square root via Linear Search | $\mathcal{O}(\sqrt{n})$ | $\mathcal{O}(1)$ |
| **VI.7** | Search in an unbalanced BST (worst-case) | $\mathcal{O}(n)$ | $\mathcal{O}(n)$ |
| **VI.8** | Search in an unsorted/ordinary Binary Tree | $\mathcal{O}(n)$ | $\mathcal{O}(h)$ (balanced: $\mathcal{O}(\log n)$) |
| **VI.9** | Repeatedly append to a new array of size $len+1$ | $\mathcal{O}(n^2)$ | $\mathcal{O}(n)$ |
| **VI.10** | Sum of digits in a number | $\mathcal{O}(\log n)$ | $\mathcal{O}(1)$ |
| **VI.11** | Generate all sorted strings of length $k$ (character set $c=26$) | $\mathcal{O}(k \cdot 26^k)$ | $\mathcal{O}(k)$ |
| **VI.12** | Compute array intersection via sorting and binary search | $\mathcal{O}(b \log b + a \log b)$ | $\mathcal{O}(1)$ or $\mathcal{O}(b)$ |


