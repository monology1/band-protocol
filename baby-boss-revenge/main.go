package main

import (
	"fmt"
	"strings"
)

func main() {
	testCases := []struct {
		input    string
		expected string
	}{
		{"SRSSRRR", "Good boy"},
		{"RSSRR", "Bad boy"},
		{"SSSRRRRS", "Bad boy"},
		{"SRRSSR", "Bad boy"},
		{"SSRSRRR", "Good boy"},
	}
	for _, tc := range testCases {
		result := isBossBabyGood(tc.input)
		if result != tc.expected {
			fmt.Printf("FAIL: expected %s, got %s\n", tc.expected, result)
		} else {
			fmt.Printf("PASS: expected %s, got %s\n", tc.expected, result)
		}
	}
}
func isBossBabyGood(s string) string {
	//time complexity: O(n) space complexity: O(1)
	s = strings.ToLower(s)
	shots := 0
	firstShot := false
	for _, ch := range s {
		if ch == 's' {
			shots++
			firstShot = true
		} else if ch == 'r' {
			//It Revenge when kids shots first
			if !firstShot {
				return "Bad boy"
			}
			if shots > 0 {
				shots--
			}
		}
	}

	// If all shots have been revenged, return "Good boy"
	if shots == 0 {
		return "Good boy"
	}
	return "Bad boy"
}
