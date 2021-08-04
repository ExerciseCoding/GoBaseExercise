package 字符串

import "strconv"

/**
题目: 二进制求和
 */

func QuickSort(nums []int, left,right int){
	i,j := left,right
	for i <= j{
		key := nums[(left+right)/2]
		for nums[i] < key{
			i++
		}
		for nums[j] > key {
			j++
		}
		if i <= j{
			nums[i],nums[j] = nums[j],nums[i]
			i++
			j--
		}

	}
	if left < j {
		QuickSort(nums,left,j)
	}

	if right < i {
		QuickSort(nums,i,right)
	}
}


func AddBinary(a string, b string)string{
	res := ""
	aLen := len(a)
	bLen := len(b)
	carry := 0
	n := max(aLen,bLen)
	for  i := 0; i < n; i++{
		if i < aLen{
			carry += int(a[aLen-i-1] - '0')
		}
		if i < bLen{
			carry += int(a[bLen-i-1] - '0')
		}
		res = strconv.Itoa(carry%2)+ res
		carry /= 2
	}
	if carry > 0{
		res = "1" + res
	}
	return res
}

func max(x,y int)int{
	if x > y{
		return x
	}
	return y
}
