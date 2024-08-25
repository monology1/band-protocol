package main

import (
	"fmt"
	"sort"
)

func main() {
	testCases := []struct {
		n, k     int
		position []int
		expected int
	}{
		{4, 3, []int{1, 2, 3, 10}, 3},
		{5, 100, []int{10, 20, 30, 40, 50}, 5},
		{3, 1, []int{10, 20, 30}, 1},
		{6, 4, []int{1, 2, 3, 6, 7, 8}, 3},
		{5, 2, []int{1, 2, 999999998, 999999999, 1000000000}, 2},
		{1000000, 1000000, func() []int {
			arr := make([]int, 1000000)
			for i := 0; i < 1000000; i++ {
				arr[i] = i + 1
			}
			return arr
		}(), 1000000},
	}
	for i, tc := range testCases {
		result := OptimizeSupermanChickenRescue(tc.n, tc.k, tc.position)
		if result != tc.expected {
			fmt.Printf("Test Case %d: FAIL - got %d, expected %d\n", i+1, result, tc.expected)
		} else {
			fmt.Printf("Test Case %d: PASS - got %d, expected %d\n", i+1, result, tc.expected)
		}
	}
}

func BruteForceSupermanChickenRescue(n, k int, position []int) int {
	//brute force
	//time complexity: O(n^2) space complexity: O(1)
	//max rescue chicken
	maxChickens := 0
	//Iteration position of each chicken
	//[1, 2, 3]
	for i := 0; i <= position[n-1]; i++ {
		chickens := 0
		for j := 0; j < n; j++ {
			if position[j] >= i && position[j] < i+k {
				chickens++
			}
			maxChickens = Max(maxChickens, chickens)
		}
	}
	return maxChickens
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func OptimizeSupermanChickenRescue(n, k int, position []int) int {
	//time complexity: O(n log n) because using sort, space complexity: O(1)
	sort.Ints(position)
	maxChickens := 0
	left := 0
	for right := 0; right < n; right++ {
		for position[right]-position[left] >= k {
			left++
		}
		if right-left+1 > maxChickens {
			maxChickens = right - left + 1
		}
	}
	return maxChickens
}
