## Tabulation AKA Bottom Up Approach

Dynamic Programming (DP) is a method used to solve complex problems by breaking them down into simpler subproblems, solving each subproblem once, and storing their results to avoid redundant computations. The **Tabulation** technique, also known as the **Bottom-Up Approach**, is one way to implement DP. This approach involves solving all related subproblems first and using their solutions to build up the solution to the original problem. It typically employs an iterative process and often utilizes a table (array) to store intermediate results.

**Key Characteristics of Tabulation:**

- **Iterative Computation:** The approach avoids recursion by solving subproblems in an iterative manner.
- **Precomputed Subproblem Solutions:** It starts by solving the smallest subproblems and uses their solutions to tackle larger subproblems.
- **Table Storage:** Intermediate results are stored in a table (usually an array) to facilitate efficient computation of the overall solution.

**Example: Calculating Fibonacci Numbers Using Tabulation**

The Fibonacci sequence is a classic example to demonstrate the tabulation approach. The sequence is defined as:

- F(0) = 0
- F(1) = 1
- F(n) = F(n-1) + F(n-2) for n ≥ 2

Here's how you can implement this in Go using the tabulation method:

```go
package main

import "fmt"

// fibonacci calculates the nth Fibonacci number using the tabulation approach.
func fibonacci(n int) int {
    if n <= 1 {
        return n
    }
    // Initialize a slice to store Fibonacci numbers up to n
    fib := make([]int, n+1)
    fib[0], fib[1] = 0, 1
    // Iteratively compute each Fibonacci number from 2 to n
    for i := 2; i <= n; i++ {
        fib[i] = fib[i-1] + fib[i-2]
    }
    return fib[n]
}

func main() {
    n := 10
    fmt.Printf("The %dth Fibonacci number is: %d\n", n, fibonacci(n))
}
```

**Explanation:**

1. **Initialization:** A slice `fib` of length `n+1` is created to store Fibonacci numbers from `F(0)` to `F(n)`. The base cases are initialized as `fib[0] = 0` and `fib[1] = 1`.

2. **Iterative Computation:** A `for` loop iterates from `2` to `n`, calculating each Fibonacci number by summing the two preceding numbers: `fib[i] = fib[i-1] + fib[i-2]`.

3. **Result:** After the loop completes, `fib[n]` contains the nth Fibonacci number, which is then returned.

This approach efficiently computes the Fibonacci sequence in O(n) time and O(n) space. By iteratively building up the solution from the base cases, it avoids the overhead associated with recursive calls and redundant computations.

**Example: Longest Common Subsequence (LCS) Using Tabulation**

The LCS problem involves finding the longest subsequence present in both of two given sequences. Here's how you can implement it in Go using the tabulation method:

```go
package main

import "fmt"

// longestCommonSubsequence calculates the length of the longest common subsequence between two strings.
func longestCommonSubsequence(text1, text2 string) int {
    m, n := len(text1), len(text2)
    // Initialize a 2D slice to store lengths of LCS for substrings
    dp := make([][]int, m+1)
    for i := range dp {
        dp[i] = make([]int, n+1)
    }
    // Fill the dp table in a bottom-up manner
    for i := 1; i <= m; i++ {
        for j := 1; j <= n; j++ {
            if text1[i-1] == text2[j-1] {
                dp[i][j] = dp[i-1][j-1] + 1
            } else {
                dp[i][j] = max(dp[i-1][j], dp[i][j-1])
            }
        }
    }
    return dp[m][n]
}

// max returns the maximum of two integers.
func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func main() {
    text1 := "abcde"
    text2 := "ace"
    fmt.Printf("The length of the longest common subsequence is: %d\n", longestCommonSubsequence(text1, text2))
}
```

**Explanation:**

1. **Initialization:** A 2D slice `dp` of size `(m+1) x (n+1)` is created, where `m` and `n` are the lengths of `text1` and `text2`, respectively. This table stores the lengths of LCS for substrings of `text1` and `text2`.

2. **Iterative Computation:** Nested loops iterate over each character of `text1` and `text2`. If characters match (`text1[i-1] == text2[j-1]`), `dp[i][j]` is set to `dp[i-1][j-1] + 1`. Otherwise, it's set to the maximum of `dp[i-1][j]` and `dp[i][j-1]`.

3. **Result:** After filling the table, `dp[m][n]` contains the length of the LCS, which is then returned.

This implementation efficiently computes the LCS length in O(m\*n) time and space, providing an optimal solution to the problem.

By employing the tabulation approach in Go, you can solve various dynamic programming problems efficiently, leveraging iterative computation and table storage to build up solutions from simpler subproblems.
