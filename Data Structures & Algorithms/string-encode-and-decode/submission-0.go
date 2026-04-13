
type Solution struct{}

func (s *Solution) Encode(strs []string) string {
	for index,value:=range strs {
		length:=strconv.Itoa(len(value))
		delimitter:=length+"#"
		strs[index]=delimitter+value
	}
	result:=strings.Join(strs,"")
	return result
}

func (s *Solution) Decode(encoded string) []string {
	i:=0
	result:=[]string{}
	for i < len(encoded) {
		j:=i
		for j < len(encoded) && encoded[j]!='#' {
			j++
		}
		length,_:=strconv.Atoi(encoded[i:j])
		start:=j+1
		end:=start+length
		str:=encoded[start:end]
		result=append(result,str)
		i=end
	}
	fmt.Println(result)
	return result

}
