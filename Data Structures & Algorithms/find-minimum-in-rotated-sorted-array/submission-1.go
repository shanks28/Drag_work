func findMin(nums []int) int {
	left:=0
	right:=len(nums)-1
	for left < right {
		if nums[left] < nums[right]{
			break
		}
		mid:=(left+right)/2
		if nums[mid] < nums[left] {
			right=mid
		} else {
			left=mid+1
		}
	}
	return nums[left]
}
