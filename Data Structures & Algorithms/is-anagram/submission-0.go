func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	hash_map_s:=make(map[rune]int)
	i:=0
	for i < len(s) {
		if value,ok:=hash_map_s[rune(s[i])]; ok {
			hash_map_s[rune(s[i])]=value+1
		} else {
			hash_map_s[rune(s[i])]=1
		}
		i++
	}
	hash_map_t:=make(map[rune]int)
	j:=0
	for j < len(t) {
		if value,ok:=hash_map_t[rune(t[j])]; ok {
			hash_map_t[rune(t[j])]=value+1
		} else {
			hash_map_t[rune(t[j])]=1
		}
		j++
	}
	fmt.Println(hash_map_s,hash_map_t)
	for key,value:=range(hash_map_t) {
		if value != hash_map_s[key]{
			return false
		}
	}
	
	return true
}
