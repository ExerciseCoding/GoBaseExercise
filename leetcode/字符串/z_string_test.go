package 字符串

import "strings"

/**
题目: Z字形变换
 */

func converts(s string, numRows int)string{
	res := make([]string,numRows)
	n := 2 * numRows - 2
	for index,value := range s {
		x := index % n
		res[min(x,n-x)] += string(value)
	}
	return strings.Join(res,"")
}
