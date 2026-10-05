func reverseString(s []byte) {
	var left int=0
	var right int=len(s)-1
	for left < right {
		temp:=s[left]
		s[left]=s[right]
		s[right]=temp
		left+=1
		right-=1
	}
}
