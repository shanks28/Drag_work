func productExceptSelf(nums []int) []int {
	n:=len(nums)
	prefix_array:=make([]int,n)
	suffix_array:=make([]int,n)
	for i:=0 ; i < n ;i++ {
		j:=i+1
		product:=1
		for j < n {
			product=product*nums[j]
			j++
		}
		suffix_array[i]=product
	}
	for i:=n-1 ; i > -1; i--{
		j:=i-1
		product:=1
		for j > -1 {
			product=product*nums[j]
			j--
		}
		prefix_array[i]=product

	}
	result:=make([]int,n)
	for i:=0 ;i<n ;i++ {
		result[i]=prefix_array[i] * suffix_array[i]
	}
	return result

}
