package leetcode

func search(nums []int, target int) int {
	start := 0
	end := len(nums) - 1

	for {
		if start > end {
			break
		}
		mid := start + (end-start)/2
		if nums[mid] > target {
			end = mid - 1
		} else if target < nums[mid] {
			start = mid + 1
		} else {
			return mid
		}
	}
	return -1
}

func isBadVersion(version int) bool{
	return true
}
func firstBadVersion(n int) int {
	start := 0
	end := n -1
	for{
		if start > end {
			break
		}
		mid := start + (end - start) /2
		if isBadVersion(mid) {
			end = mid - 1
		}else{
			start = end + 1
		}
	}
	return start
}

func searchInsert(nums []int, target int) int {
	start := 0
	end := len(nums) -1
	for {
		if start > end {
			break
		}
		mid := start + (end - start) /2
		if target < nums[mid]  {
			end = mid - 1
		} else if target >= nums[mid]{
			start = mid + 1
		}
	}
	return end
}