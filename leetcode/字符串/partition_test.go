package 字符串
/**
题目: 分割回文串
 */

func partion(s string)[][]string{

}


func partionFind(s string,temp []string,result *[][]string,start int){
	if start == len(s){
		tempStr := make([]string,len(temp))
		copy(tempStr,temp)
		*result = append(*result,tempStr)
		return
	}

	for i := start; i < len(s); i++{
		if isPail(s,start,i){
			temp = append(temp,s[start:i])
			partionFind(s,temp,result,i+1)
			temp = temp[:len(temp)-1]
		}
	}
}

func isPail(s string,start,end int)bool{
	for start < end {
		if s[start] != s[end]{
			return false
		}
		start++
		end--

	}
	return true
}