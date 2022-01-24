package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	fmt.Println("vim-go")

	var nums []int
	var target int
	var output [][]int

	nums = []int{1, 0, -1, 0, -2, 2}
	target = 0

	output = fourSum(nums, target)
	fmt.Println(output)
}

func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	return nSum(0, 4, nums, target)
}

func nSum(left int, n int, nums []int, target int) [][]int {
	var res [][]int
	if n == 2 {
		return twoSum(left, nums, target)
	} else {
		for pivot := left; pivot < len(nums)-(n-1); pivot++ {
			if pivot == left || nums[pivot] != nums[pivot-1] {
				subres := nSum(pivot+1, n-1, nums, target-nums[pivot])
				for _, v := range subres {
					if len(v) != 0 {
						v = append([]int{nums[pivot]}, v...)
						res = append(res, v)
					}
				}
			}
		}
	}

	return res
}

func twoSum(left int, nums []int, target int) [][]int {
	checkDuplicate := make(map[string]struct{})
	var res [][]int
	right := len(nums) - 1
	for left < right {
		sum := nums[left] + nums[right]
		if sum == target {
			key := strconv.Itoa(nums[left]) + strconv.Itoa(nums[right])
			if _, ok := checkDuplicate[key]; !ok { //check if there is duplicated slices
				res = append(res, []int{nums[left], nums[right]})
				checkDuplicate[key] = struct{}{}
			}
			left++
			right--
		} else if sum > target {
			right--
		} else {
			left++
		}
	}

	return res
}
