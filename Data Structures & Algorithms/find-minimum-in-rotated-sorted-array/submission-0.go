func findMin(nums []int) int {
	if len(nums) ==1 {
		return nums[0]
	}
	left:=0
	right:=len(nums)-1
	for left <=right {
		next:=left+1
		if next < len(nums) && nums[left] > nums[next] {
			return nums[next]
		}
		left++
	}
	return nums[0]
	

}
