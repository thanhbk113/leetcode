package main

import (
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

func containsDuplicate(nums []int) bool {
	size := len(nums)
	check := make(map[int]bool, size)

	for i := 0; i < size; i++ {
		if check[nums[i]] {
			return true
		} else {
			check[nums[i]] = true
		}
	}

	return false
}

func BenchmarkContainsDuplicate(b *testing.B) {
	nums := []int{1, 2, 3, 1}
	containsDuplicate(nums)
}
