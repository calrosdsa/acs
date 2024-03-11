package main

import (
	"log"
	"math"
	"sort"
	// "github.com/aws/aws-sdk-go/service/bedrockagent"
)

func main() {
	// strs := []string{"reflower","flow","flight"}
	// res := longestCommonPrefix(strs)
	nums := []int{-1, 2, 1, -4, 7}
	res := threeSumClosest(nums, 1)
	log.Println("RES", res)
}

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	abs := func(n int) int {
		if n < 0 {
			return -n
		}
		return n
	}
	closestSum := math.MaxInt
	for i := 0; i < len(nums); i++ {
		left, right := i+1, len(nums)-1
		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			if abs(sum-target) < abs(closestSum-target) {
				closestSum = sum
			}

			if sum < target {
				left++
			} else if sum > target {
				right--
			} else {
				return closestSum
			}

		}
	}
	return closestSum
}

func longestCommonPrefix(strs []string) string {
	p := strs[0]
	for _, s := range strs {
		i := 0
		for ; i < len(s) && i < len(p) && p[i] == s[i]; i++ {
		}
		p = p[:i]
		log.Println(p, i)
	}
	return p
}

// func lengthOfLongestSubstring(s string) int {
// 	store := make(map[uint8]int)
// 	var left, right, result int

// 	for right < len(s) {
// 		var r = s[right]
// 		store[r] += 1
// 		for store[r] > 1 {
// 			l := s[left]
// 			store[l] -= 1
// 			left += 1
// 		}
// 		log.Println(result, (right-left+1))
// 		result = max(result, right-left+1)
// 		right += 1
// 	}
// 	return result
// }

// func max(a, b int) int {
// 	if a > b {
// 		return a
// 	}
// 	return b
// }
