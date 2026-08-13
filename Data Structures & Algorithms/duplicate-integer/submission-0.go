func hasDuplicate(nums []int) bool {
    var set map[int]int = map[int]int{} // nil map
	for i,v:=range nums {
		if _,ok:=set[v]; ok {
			return true
		}
		set[v]=i
	}
	return false
}
