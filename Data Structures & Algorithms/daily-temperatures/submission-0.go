func dailyTemperatures(temperatures []int) []int {
	n:=len(temperatures)	
	result:=make([]int,n)
	stack:=[]int{}
	for i :=0 ; i<n ;i ++ {
		
		for len(stack) > 0 && temperatures[stack[len(stack)-1]] < temperatures[i] {
			
			topIndex:=stack[len(stack)-1]
			stack=stack[:len(stack)-1]
			result[topIndex]=i-topIndex
			
		}
		stack=append(stack,i)
	}
	return result
}
