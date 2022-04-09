package main

func fourSum(nums []int, target int) [][]int {
	var bnm [][]int
	sort(nums)
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums)-1; j++ {
			for k := j + 1; k < len(nums)-1; k++ {
				for m := k + 1; m < len(nums); m++ {
					if check_arr(nums[i], nums[j], nums[k], nums[m], target) {
						rev := []int{nums[i], nums[j], nums[k], nums[m]}
						flags := true
						for w := 0; w < len(bnm); w++ {
							if !check(bnm[w], rev) {
								flags = false
							}
						}
						if flags {
							bnm = append(bnm, rev)
						}
					}
				}
			}
		}
	}
	return bnm
}

func check_arr(a, b, c, d, target int) bool {
	if a+b+c+d == target {
		return true
	}
	return false
}

func sort(arr []int) {
	for {
		flags := true
		for j := 1; j < len(arr); j++ {
			if arr[j-1] > arr[j] {
				arr[j-1], arr[j] = arr[j], arr[j-1]
				flags = false
			}
		}
		if flags {
			return
		}
	}
}

func check(a []int, b []int) bool {
	count := 0
	for i := 0; i < len(a); i++ {
		if a[i] == b[i] {
			count++
		}
	}
	if count == len(a) {
		return false
	} else {
		return true
	}
}
