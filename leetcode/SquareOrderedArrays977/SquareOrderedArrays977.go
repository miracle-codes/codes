package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := make([]int, 0)
	nums = append(nums, -4, -1, 0, 3, 10)
	fmt.Println("before:", nums)
	nums = sortedSquares2(nums)
	fmt.Println("resturn:", nums)
}

//题目
//977.有序数组的平方
//https://leetcode.cn/problems/squares-of-a-sorted-array/
//给你一个按 非递减顺序 排序的整数数组 nums，返回 每个数字的平方 组成的新数组，要求也按 非递减顺序 排序。
//
//示例 1：
//
//输入：nums = [-4,-1,0,3,10]
//输出：[0,1,9,16,100]
//解释：平方后，数组变为 [16,1,0,9,100]，排序后，数组变为 [0,1,9,16,100]
//示例 2：
//
//输入：nums = [-7,-3,2,3,11]
//输出：[4,9,9,49,121]

// 暴力排序
// 最直观的想法，莫过于：每个数平方之后，排个序，代码如下：
// 这个时间复杂度是 O(n + nlogn)， 可以说是O(nlogn)的时间复杂度，
// 但为了和下面双指针法算法时间复杂度有鲜明对比，我记为 O(n + nlog n)。
func sortedSquares(nums []int) []int {
	res := make([]int, 0)
	for _, num := range nums {
		num = num * num
		res = append(res, num)
	}
	//res = sortArray(res)
	sort.Ints(res)
	return res
}

// 双指针
func sortedSquares2(nums []int) []int {
	j, s := len(nums)-1, len(nums)-1
	res := make([]int, len(nums))
	for i := 0; i <= j; {
		if nums[i]*nums[i] >= nums[j]*nums[j] {
			res[s] = nums[i] * nums[i]
			i++
		} else {
			res[s] = nums[j] * nums[j]
			j--
		}
		s--
	}
	return res

}
