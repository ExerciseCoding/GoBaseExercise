package 字符串

func longestCommonPrefix(nums []string)string{
	if len(nums) == 0{
		return ""
	}
	prefix := nums[0]
	for i := 1 ; i < len(nums); i++{
		prefix = lcp(prefix,nums[i])
		if len(prefix) == 0{
			break
		}
	}
	return prefix

}
func lcp(str1,str2 string)string{
	len := min(len(str1),len(str2))
	index := 0
	for index < len && str1[index] == str2[index]{
		index++
	}
	return str1[:index]
}
func min(x , y int)int{
	if x < y{
		return x
	}
	return y
}