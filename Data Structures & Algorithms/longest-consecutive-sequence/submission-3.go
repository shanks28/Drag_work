import (
	"slices"
)
func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	nums_set:=make(map[int]struct{})
	for _,value:= range nums {
		if _,ok:=nums_set[value]; !ok {
			nums_set[value]=struct{}{}
		}
	}
	keys:=[]int{}
	for key,_:=range nums_set {
		keys=append(keys,key)
	}
	slices.Sort(keys[:])
	res:=1//min 1 largest sequence
	i:=0
	window_len:=1
	prev_element:=keys[i]
	for i < len(keys) {
		if keys[i]-prev_element == 1 {
			window_len++
		} else {
			window_len=1
		}
		prev_element=keys[i]
		res=max(res,window_len)
		i++
	}
	fmt.Println(keys)
	return res
}
