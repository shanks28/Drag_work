func isValid(s string) bool {
    hash_map:=make(map[string]string)
	hash_map["]"]="["
	hash_map["}"]="{"
	hash_map[")"]="("
	if len(s)<=1{
		return false
	}
	stack:=[]string{}
	for i:=0;i<len(s);i++ {
		n:=len(stack)
		if s[i]=='{' || s[i]=='(' || s[i]=='[' {
			fmt.Println(s[i])
			stack=append(stack,string(s[i]))
			fmt.Println(stack)
		} else {
			if n > 0{
				top:=stack[n-1]
				if hash_map[string(s[i])] == top {
					fmt.Println(top)
					stack=stack[:n-1]
				} else{
					return false
				}
			} else {
				return false
			}
		}
	}
	if len(stack) == 0{
		return true
	} else {
		return false
	}
}
