package problems

import (
	"slices"
	"strings"
	"unicode"
)

// Two Pointer
func maxArea(height []int) int {

	area, left, right := 0, 0, len(height)-1

	for left < right {
		var tmpArea int
		if height[left] > height[right] {
			tmpArea = (right - left) * height[right]
			right--
		} else {
			tmpArea = (right - left) * height[left]
			left++
		}
		if tmpArea > area {
			area = tmpArea
		}
	}

	return area
}

func threeSum(nums []int) [][]int {
	var solutions [][]int

	if len(nums) < 3 {
		return solutions
	}
	slices.Sort(nums)

	for i := range nums {

		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		target := 0 - nums[i]

		left, right := i+1, len(nums)-1

		for left < right {

			num := nums[left] + nums[right]

			if num == target {
				solutions = append(solutions, []int{nums[i], nums[left], nums[right]})
				left++
				for left < right && nums[left] == nums[left-1] {
					left++
				}
				for left < right && right < len(nums) && nums[right] == nums[right+1] {
					right--
				}
			} else if num > target {
				right--
			} else if num < target {
				left++
			}
		}
	}
	return solutions
}

func twoSum(numbers []int, target int) []int {

	left, right := 0, len(numbers)-1

	for left < right {
		cur := numbers[left] + numbers[right]

		if cur == target {
			break
		} else if cur > target {
			right--
		} else if cur < target {
			left++
		}

	}

	return []int{left + 1, right + 1}
}

func isPalindrome(s string) bool {

	format := func(r rune) rune {

		if !unicode.IsLetter(r) && !unicode.IsNumber(r) {
			return -1
		}

		return unicode.ToLower(r)

	}

	s = strings.Map(format, s)

	r, l := 0, len(s)

	for r < l {

		if s[r] != s[l] {
			return false
		}

		r++
		l--

	}
	return true
}
