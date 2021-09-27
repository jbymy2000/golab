package binary_search
//奇点允许左右相等
//每次端点变化的时候+1
func search_v2(inputSlice []int,target int)int{
	var left,right,mid int
	left = 0
	right  = len(inputSlice)-1
	for(left<=right){
		mid = (left + right)/2
		if inputSlice[mid] == target{
			return mid
		}else if inputSlice[mid]<target{
			left = mid + 1
		}else{
			right = mid - 1
		}
	}
	return -1
}