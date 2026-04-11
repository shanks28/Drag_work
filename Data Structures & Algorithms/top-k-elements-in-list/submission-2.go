
func topKFrequent(nums []int, k int) []int {
    res:=[]int{}
    hash_map:=make(map[int]int)
    for _,value:= range nums {
        if _,ok:=hash_map[value]; ok {
            new_value:=hash_map[value]+1
            hash_map[value]=new_value
        } else {
            hash_map[value]=1
        }
    }
    fmt.Println("map is :",hash_map)
    keys:=make([]int,0,len(hash_map))
    for k:=range hash_map {
        keys=append(keys,k)
    }
    sort.SliceStable(keys,func(i,j int) bool {
        return hash_map[keys[i]] > hash_map[keys[j]]
    })
    counter:=0
    for i,_:=range keys {
        res=append(res,keys[i])
        counter++
        if counter == k {
            break
        }
    }
    
    return res

}
