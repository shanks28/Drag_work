func twoSum(nums []int, target int) []int {
    var diff_map map[int]int=map[int]int{}
    for i,value:=range nums {
        var diff int = target-value
        if j,ok:=diff_map[diff]; ok {
            return []int{j,i}
        }
        diff_map[value]=i
    }
    return []int{}
}
