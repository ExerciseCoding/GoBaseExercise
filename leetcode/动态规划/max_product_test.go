package 动态规划

import (
	"math"
)

/**
题目: 找出数组中乘积最大的连续子数组
 */

func Max(x,y int)int{
	if x > y{
		return x
	}
	return y
}



func maxProduct(nums []int)int{
	max, imax,imin := int(math.MinInt32),1,1
	for i := 0; i < len(nums); i++{
		if nums[i] < 0{
			imax,imin = imin,imax
		}
		imax = Max(imax * nums[i],nums[i])
		imin = Min(imin * nums[i], nums[i])
		max = Max(imax,max)
	}
	return max
}
