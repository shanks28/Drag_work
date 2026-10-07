type MyHashMap struct {
	nestedArray [][]int
}

func Constructor() MyHashMap {
	var obj MyHashMap=MyHashMap{
		nestedArray:[][]int{},
	}
	return obj
    
}

func (this *MyHashMap) Put(key int, value int) {
	this.Remove(key)
    innerArray:=[]int{key,value}
	this.nestedArray=append(this.nestedArray,innerArray)
}

func (this *MyHashMap) Get(key int) int {
	for _,value:=range this.nestedArray {
		if value[0]==key{
			return value[1]
		} else{
			continue
		}
	}
	return -1
    
}

func (this *MyHashMap) Remove(key int) {
	for index,value:=range this.nestedArray {
		if value[0] == key {
			this.nestedArray=append(this.nestedArray[:index],this.nestedArray[index+1:]...)
		} else{
			continue
		}
	}
    
}

/**
 * Your MyHashMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Put(key,value);
 * param_2 := obj.Get(key);
 * obj.Remove(key);
 */