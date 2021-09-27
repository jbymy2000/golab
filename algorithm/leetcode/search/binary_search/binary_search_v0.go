package binary_search

func search_v0(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	return search_by_pos(nums, 0, len(nums)-1, target)
}

func search_by_pos(nums []int, start, end, target int) int {
	if start == end && nums[start] == target {
		return start
	}
	mid := (start + end) / 2
	if target >= nums[start] && target <= nums[mid] {
		return search_by_pos(nums, start, mid, target)
	} else if target > nums[mid] && target <= nums[end] {
		return search_by_pos(nums, mid+1, end, target)
	} else {
		return -1
	}
}
