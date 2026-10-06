func validPalindrome(s string) bool {
	left:=0
	right:=len(s)-1
	for left < right {
		if s[left]!=s[right] {
			//skipl or skipr
			var skipL string=s[left+1:right+1]
			var skipR string=s[left:right]
			return (isPalindrome(skipL) || isPalindrome(skipR))
		} else {
			left+=1
			right-=1
		}
	}
	return true
}
func isPalindrome(word string)(bool) {
	left:=0
	right:=len(word)-1
	for left < right {
		if word[left]!=word[right]{
			return false
		} else {
			left+=1
			right-=1
		}
	}
	return true
}
