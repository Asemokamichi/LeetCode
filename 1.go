package main

func twoSum(nums []int, target int) []int {
	var bnm []int
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				bnm = append(bnm, i)
				bnm = append(bnm, j)
				return bnm
			}
		}
	}
	return bnm
}
