
func isPalindrome(s string) bool {
	var lower_s string=""
	for _,value:= range s{
		if unicode.IsDigit(value) || unicode.IsLetter(value){
			lower_s+=string(strings.ToLower(string(value)))
		}
	}
	start:=0
	end:=len(lower_s)-1
	for start < end {
		if lower_s[start]==lower_s[end]{
			start++
			end--
		} else {
			return false
		}
	}
	return true
}
