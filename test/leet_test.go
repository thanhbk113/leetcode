package main

import (
	"fmt"
	"strconv"
	"testing"
)

// using sort
// func containsDuplicate(nums []int) bool {

// 	sort.Ints(nums)

// 	for i := 0; i < len(nums); i++ {
// 		if nums[i] == nums[i+1] {
// 			return true
// 		}
// 	}

// 	return false
// }

func countFreqChar(s string) string {
	count := make([]int, 26)

	for _, val := range s {
		count[val-'a']++
	}

	key := ""

	for _, c := range count {
		key += strconv.Itoa(c) + "/"
	}
	return key
}

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	return countFreqChar(s) == countFreqChar(t)
}

func BenchmarkContainsDuplicate(b *testing.B) {
	s := "aacc"
	t := "ccac"

	fmt.Println(isAnagram(s, t))
}
