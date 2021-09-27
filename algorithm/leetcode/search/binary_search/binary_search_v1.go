package binary_search

//只用中点校验
//如果重复的话，先看做再看右还是先看右在看左，决定找到哪个，然后 再向右看有没有

func search_v1(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	//为什么用<= 因为有只有两个数的情况
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] == target {
			return mid
		} else if nums[right] == target {
			return right
		} else if nums[mid] > target {
			//为什么要-1  等于就不动了。
			right = mid - 1
		} else {
			//为什么要+1
			left = mid + 1
		}
	}
	return -1
}
