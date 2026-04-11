func majorityElement(nums []int) int {
    hash_map:=make(map[int]int)
    for _,value:=range nums {
        if val,ok:=hash_map[value]; ok {
            hash_map[value]=val+1
        } else {
            hash_map[value] = 0 
        }
    }
    keys:=make([]int,0,len(nums))
    for key,_:=range hash_map {
        keys=append(keys,key)
    }
    sort.SliceStable(keys,func(i,j int) bool {
        return hash_map[keys[i]] > hash_map[keys[j]]
    })
    return keys[0]
}
