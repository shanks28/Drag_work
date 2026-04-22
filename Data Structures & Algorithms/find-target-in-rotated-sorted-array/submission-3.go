func search(nums []int, target int) int {
	left:=0
	right:=len(nums)-1
	if len(nums) ==1 {
		if nums[0]==target{
			return 0
		} else {
			return -1
		}
	}
	if nums[left] == target {
		return left
	} else if nums[right] == target {
		return right
	}

	for i:=0 ; i< len(nums) ; i++ {
		next:=i+1
		//point of inflexion found
		// this is the sorted left part of the array
		if next < len(nums) && nums[i] > nums[next]   {
			if idx:=bin_search(next,right,nums,target); idx!=-1{
				return idx
			}
			right=i
			break
		}
		
		// search the greater sorted half
	}
	if idx:=bin_search(left,right,nums,target); idx!=-1{
		return idx
	}
	return -1
}
//returns index if found
func bin_search(left int,right int,nums []int ,target int) int {
	for left<=right {
		mid:=(right+left)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			right=mid-1
		} else {
			left=mid+1
		}
	}
	return -1
}