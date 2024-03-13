package main

import (
	// "fmt"
	"log"
	"math"
	"sort"
	// "github.com/aws/aws-sdk-go/service/bedrockagent"
)

const (
	j = 7
	l
	m
)

func main() {
	// var i int = 2
	// log.Println(j, l, m, i)
	// a := []string{"A", "B", "C", "D", "E"}
	// a = a[:0]

	// recipes := []string{
	// 	"xevvq", "izcad", "p", "we", "bxgnm", "vpio", "i", "hjvu", "igi", "anp", "tokfq", "z", "kwdmb", "g", "qb", "q", "b", "hthy",
	// }
	// ingredients := [][]string{
	// 	{"wbjr"},
	// 	{"otr", "fzr", "g"},
	// 	{"fzr", "wi", "otr", "xgp", "wbjr", "igi", "b"},
	// 	{"fzr", "xgp", "wi", "otr", "tokfq", "izcad", "igi", "xevvq", "i", "anp"},
	// 	{"wi", "xgp", "wbjr"},
	// 	{"wbjr", "bxgnm", "i", "b", "hjvu", "izcad", "igi", "z", "g"},
	// 	{"xgp", "otr", "wbjr"},
	// 	{"wbjr", "otr"},
	// 	{"wbjr", "otr", "fzr", "wi", "xgp", "hjvu", "tokfq", "z", "kwdmb"},
	// 	{"xgp", "wi", "wbjr", "bxgnm", "izcad", "p", "xevvq"},
	// 	{"bxgnm"},
	// 	{"wi", "fzr", "otr", "wbjr"},
	// 	{"wbjr", "wi", "fzr", "xgp", "otr", "g", "b", "p"},
	// 	//g
	// 	{"otr", "fzr", "xgp", "wbjr"},
	// 	{"xgp", "wbjr", "q", "vpio", "tokfq", "we"},
	// 	{"wbjr", "wi", "xgp", "we"},
	// 	{"wbjr"},
	// 	{"wi"},
	// }
	// supplies := []string{
	// 	"wi", "otr", "wbjr", "fzr", "xgp",
	// }
	// recipes := []string{"sandwich","bread"}
	// ingredients := [][]string{{"bread","meat"}, {"yeast", "flour"}}
	// supplies := []string{"meat","yeast","flour"}
	// res := findAllRecipes(recipes, ingredients, supplies)
	candidates := []int{10, 1, 2, 7, 6, 1, 5}
	target := 8
	res := combinationSum2(candidates, target)
	log.Println("RESULT", res)
}

func combinationSum2(nums []int, target int) [][]int {
	var res [][]int
	sort.Ints(nums)
	// 1 1 2 5 6 7 10
	search(nums, target, []int{}, &res)
	return res
}

func search(candidates []int, target int, cur []int, ans *[][]int) {
	if target == 0 {
		temp := []int{}
		temp = append(temp, cur...)
		*ans = append(*ans, temp)
	}
	for i := 0; i < len(candidates); i++ {
		if candidates[i] > target {
			break
		}

		if i > 0 && candidates[i] == candidates[i-1] {
			continue
		}

		cur = append(cur, candidates[i])
		log.Println(cur)
		search(candidates[i+1:], target-candidates[i], cur, ans)
		cur = cur[:len(cur)-1]
	}
}

func findAllRecipes(recipes []string, ingredients [][]string, supplies []string) []string {
	// Use a dependents map for easier lookup
	dependents := make(map[string][]string)
	outdegrees := make(map[string]int)
	for i := 0; i < len(recipes); i++ {
		for _, ingredient := range ingredients[i] {
			dependents[ingredient] = append(dependents[ingredient], recipes[i])
			outdegrees[recipes[i]]++
		}
	}
	log.Println("DEPENDENTS", dependents)
	log.Println("OUTDEGREE", outdegrees)
	recipesMap := make(map[string]struct{})
	for _, recipe := range recipes {
		recipesMap[recipe] = struct{}{}
	}
	log.Println("recipesMap", recipesMap)
	var queue []string
	// Add all the supplies; they shouldn't have any outdegree
	queue = append(queue, supplies...)
	var res []string
	var curr string
	log.Println("QEUES", queue)
	for len(queue) != 0 {
		curr, queue = queue[0], queue[1:]
		log.Println("CURR", curr, "QUEUE", queue)
		if _, ok := recipesMap[curr]; ok {
			log.Println("ALREDY EXIST", curr)
			res = append(res, curr)
		}
		for _, dependent := range dependents[curr] {
			outdegrees[dependent]--
			if outdegrees[dependent] == 0 {
				queue = append(queue, dependent)
				log.Println("== 0", queue)
			}
		}
		log.Println(outdegrees)
	}
	return res
}

func combinationSum(candidates []int, target int) [][]int {
	var ans [][]int

	if len(candidates) == 0 {
		return ans
	}

	bt(&ans, make([]int, 0), candidates, 0, target)

	return ans
}

func bt(ans *[][]int, tmp, nums []int, idx, tgt int) {
	var force bool
	defer func() {
		log.Println("EXIT WITH", tgt, tmp, force, "TARGET", tgt)
	}()
	log.Println("INIT", idx, tmp, tgt)
	if tgt < 0 || idx > len(nums) {
		log.Println("RETURNING")
		force = true
		return
	}

	if tgt == 0 {
		cpyTmp := make([]int, len(tmp))
		copy(cpyTmp, tmp)
		log.Println("COPY", cpyTmp)
		*ans = append(*ans, cpyTmp)
	}

	for i := idx; i < len(nums); i++ {
		tmp = append(tmp, nums[i])
		bt(ans, tmp, nums, i, tgt-nums[i])
		tmp = tmp[:len(tmp)-1]
		log.Println(tmp, "PREV", ":", i)
	}
}

func Search(n int, f func(int) bool) int {
	// Define f(-1) == false and f(n) == true.
	// Invariant: f(i-1) == false, f(j) == true.
	i, j := 0, n
	for i < j {
		h := int(uint(i+j) / 2) // avoid overflow when computing h
		// log.Println("HALF",h)
		// i ≤ h < j
		if !f(h) {
			i = h + 1 // preserves f(i-1) == false
		} else {
			log.Println("HALF CUT", h)
			j = h // preserves f(j) == true
		}
	}
	// i == j, f(i-1) == false, and f(j) (= f(i)) == true  =>  answer is i.
	return i
}

func platesBetweenCandles(s string, queries [][]int) []int {
	candles := make([]int, 0, len(s))

	for i, ch := range s {
		if ch == '|' {
			candles = append(candles, i)
		}
	}
	log.Println("CANDLES", candles)

	ans := make([]int, 0, len(queries))

	for _, query := range queries {
		l := query[0]
		r := query[1]

		left := Search(len(candles), func(i int) bool {
			// log.Println("LEFT I",i,candles[i],l)
			return candles[i] >= l
		})
		log.Println("LEFT", left, candles[left])

		right := Search(len(candles), func(i int) bool {
			log.Println("RIGHT I", i, candles[i], r)
			return candles[i] > r
		}) - 1
		log.Println("RIGHT", right, candles[right])

		if right <= left {
			ans = append(ans, 0)
			continue
		}
		c := (candles[right] - candles[left]) - (right - left)
		log.Println("C", c)
		ans = append(ans, c)
	}

	return ans
}

func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}
	left, right := 0, len(nums)-1
	for left < right && (nums[right] != target || nums[left] != target) {
		if nums[right] != target {
			right--
		}
		if nums[left] != target {
			left++
		}
	}
	if nums[right] == target && nums[left] == target {
		return []int{left, right}
	}
	if nums[left] != target && nums[right] != target {
		return []int{-1, -1}
	}
	if nums[left] == target && nums[right] != target {
		return []int{left, left}
	}
	if nums[right] == target && nums[left] != target {
		return []int{right, right}
	}
	return []int{-1, -1}
}

func removeElement(nums []int, val int) int {
	var j int
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			log.Println(i, j)
			nums[j], nums[i] = nums[i], nums[j]
			j++
		}
	}
	log.Println(nums)
	return j
}

func threeSum(nums []int) [][]int {
	var res [][]int
	if len(nums) < 3 {
		return res
	}
	sort.Ints(nums)
	// -4 -1 -1 0 1 2
	// m := make(map[[3]int]struct{})
	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		left, right := i+1, len(nums)-1
		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			if sum == 0 {
				res = append(res, []int{nums[i], nums[left], nums[right]})
				right--
				for left < right && nums[right] == nums[right+1] {
					right--
				}
			} else if sum > 0 {
				right--
			} else {
				left++
			}
		}
	}
	return res
}

func threeSum2(nums []int) [][]int {
	n := len(nums)
	sort.Ints(nums)
	var result [][]int
	for num1Idx := 0; num1Idx < n-2; num1Idx++ {
		if num1Idx > 0 && nums[num1Idx] == nums[num1Idx-1] {
			continue
		}
		num2Idx := num1Idx + 1
		num3Idx := n - 1
		for num2Idx < num3Idx {
			sum := nums[num2Idx] + nums[num3Idx] + nums[num1Idx]
			if sum == 0 {
				result = append(result, []int{nums[num1Idx], nums[num2Idx], nums[num3Idx]})
				num3Idx--
				// Skip all duplicates from right
				for num2Idx < num3Idx && nums[num3Idx] == nums[num3Idx+1] {
					num3Idx--
				}
			} else if sum > 0 {
				num3Idx--
			} else {
				num2Idx++
			}
		}
	}
	return result
}

func fourSum(nums []int, target int) [][]int {
	var res [][]int
	if len(nums) < 4 {
		return res
	}
	sort.Ints(nums)
	m := make(map[[4]int]struct{})

	for i := 0; i < len(nums)-3; i++ {
		for j := i + 1; j < len(nums); j++ {
			left, right := j+1, len(nums)-1
			for left < right {
				sum := nums[i] + nums[j] + nums[left] + nums[right]
				// log.Println(sum)
				if sum == target {
					val := [4]int{nums[i], nums[j], nums[left], nums[right]}
					if _, ok := m[val]; !ok {
						res = append(res, []int{nums[i], nums[j], nums[left], nums[right]})
						m[val] = struct{}{}
					}
					left++
					right--
					for left < right && nums[left] == nums[left-1] {
						left++
					}
					for left < right && nums[right] == nums[right+1] {
						right--
					}

				} else if sum > target {
					right--
				} else {
					log.Println("MOVE TO LEFT ")
					left++
				}
			}
		}
	}
	return res
}

func fourSum2(nums []int, target int) [][]int {
	var res [][]int
	if len(nums) < 4 {
		return res
	}
	sort.Ints(nums)
	m := make(map[[4]int]struct{})
	for i := 0; i < len(nums)-3; i++ {
		for j := i + 1; j < len(nums); j++ {
			left, right := j+1, len(nums)-1
			for left < right {
				sum := nums[i] + nums[j] + nums[left] + nums[right]
				if sum == target {
					val := [4]int{nums[i], nums[j], nums[left], nums[right]}
					if _, ok := m[val]; !ok {
						res = append(res, []int{nums[i], nums[j], nums[left], nums[right]})
						m[val] = struct{}{}
					}
					left++
					right--
					for left < right && nums[left] == nums[left-1] {
						left++
					}

					for left < right && nums[right] == nums[right+1] {
						right--
					}
				} else if sum > target {
					right--
				} else {
					left++
				}
			}
		}
	}

	return res
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
