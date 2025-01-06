package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	// arr := [][]int{
	//     {2, 1, 1},
	//     {1, 1, 0},
	//     {0, 1, 1},

	// }
	// arr := []int{1,2,3,4}
	// arr = arr[:len(arr)-1]
	// fmt.Println(isValid("(("))

	// obj := problems.Constructor()
	// obj.Push(10)
	// obj.Push(10)
	// fmt.Println(obj.Top())

	fmt.Println(topKFrequent([]int{-1, -1}, 1))

}

func topKFrequent(nums []int, k int) []int {

	buckets := make([][]int, len(nums)+1)
	topk := make([]int, 0, k)
	freqs := make(map[int]int)

	for _, num := range nums {
		freqs[num]++
	}

	for key, value := range freqs {
		if buckets[value] == nil {
			buckets[value] = make([]int, 0)
		}
		buckets[value] = append(buckets[value], key)
	}

	for i := len(buckets) - 1; i >= 0 && len(topk) != k; i-- {
		if len(buckets[i]) > 0 {
			topk = append(topk, buckets[i]...)
		}
	}

	return topk
}

func longestConsecutive(nums []int) int {
	seqs := make(map[int]struct{})
	for _, v := range nums {
		seqs[v] = struct{}{}
	}

	distance := 0

	for v := range seqs {
		if _, exists := seqs[v-1]; !exists {
			ldist := 1
			for {
				if _, ok := seqs[v+ldist]; ok {
					ldist++
				} else {
					break
				}
			}

			if ldist > distance {
				distance = ldist
			}
		}
	}
	return distance
}

func miniMaxSum(arr []int32) {
	intArr := make([]int, len(arr))
	for i, v := range arr {
		intArr[i] = int(v)
	}
	sort.Ints(intArr)
	var min, max int
	for i := 0; i < 4; i++ {
		min += intArr[i]
		max += intArr[len(intArr)-i-1]
	}
	fmt.Println(min, max)
}

func staircase(n int32) {
	// Write your code here

	for i := 1; i <= int(n); i++ {
		str := fmt.Sprintf("%*s", n, strings.Repeat("#", i))
		fmt.Println(str)
	}
}

func maxScore(s string) int {
	pRight, count := make([]int, len(s)), 0

	for i := range s {
		if s[len(s)-i-1] == '1' {
			count++
		}
		pRight[len(s)-i-1] = count
	}

	zeros, max := 0, 0
	for i := 1; i < len(s)-1; i++ {
		if s[i] == '0' {
			zeros++
		}
		sum := pRight[i] + zeros
		if sum > max {
			max = sum
		}
	}
	return max
}

func isValid(s string) bool {
	if len(s) <= 1 {
		return false
	}
	stack := make([]rune, 0)
	for _, v := range s {
		switch v {
		case '(':
			stack = append(stack, ')')
		case '[':
			stack = append(stack, ']')
		case '{':
			stack = append(stack, '}')
		default:
			if len(stack) == 0 {
				return false
			}
			pop := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if len(stack) != 0 || v != pop {
				return false
			}
		}
	}
	return true
}

func trap(height []int) int {

	left, right, trapped := 0, 0, 0

	for i := 0; i <= len(height); i++ {

		if height[i] <= 0 {
			continue
		}
		tmptrap := 0
		right = i + 1

		for height[right] < height[i] {
			right++
		}

		for j := i + 1; j < len(height); j++ {
			right = height[j]
			if right < left {
				tmptrap += left - height[j]
				continue
			}
		}

		trapped += tmptrap

	}

	return trapped
}
