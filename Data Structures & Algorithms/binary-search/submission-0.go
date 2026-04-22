func search(nums []int, target int) int {
    if len(nums) == 0 {
        return -1
    }
    left:=0
    right:=len(nums)-1
    for left <= right {
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
