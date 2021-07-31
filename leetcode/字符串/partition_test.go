package 字符串

/**
题目: 分割回文串
 */

// 回溯法
func partion(s string)[][]string{
	result := [][]string{}
	partionFind(s,[]string{},&result,0)
	return result
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



// 回溯+记忆法

func partionFindMem(s string,temp []string,result *[][]string,start int,mem [][]int){
	if start == len(s){
		tempResult := make([]string,len(temp))
		copy(tempResult,temp)
		*result  = append(*result,tempResult)
		return
	}
	for i := start; i < len(s); i++{
		if mem[start][i] == 2 {
			continue
		}
		if mem[start][i] == 1 && isPail(s,start,i){
			temp = append(temp,s[start:i+1])
			partionFindMem(s,temp,result,i+1,mem)
		}
	}
}


func isPailMem(s string,start,end int,mem [][]int)bool{
	for start < end {
		if s[start] != s[end]{
			mem[start][end] = 2
			return false
		}
		start++
		end--
	}
	mem[start][end] = 1
	return true
}
// 回溯+动态规划
