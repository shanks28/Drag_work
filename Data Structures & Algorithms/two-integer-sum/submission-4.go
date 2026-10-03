func twoSum(nums []int, target int) []int {
    var hash_map map[int]int=make(map[int]int) // allocate memory also
    for index,value:=range nums {
        var diff int = target-value
        if val,ok:=hash_map[value]; ok {
            fmt.Println(val,index)
            return []int{val,index}
        } else{
            hash_map[diff]=index
        }
    }
    return []int{}
}
