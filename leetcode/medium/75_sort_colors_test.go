package medium

import (
	"fmt"
	"testing"
)

/**
Leetcode: 75题：颜色分类
给定一个包含红色、白色和蓝色，一共 n 个元素的数组，原地对它们进行排序，使得相同颜色的元素相邻，并按照红色、白色、蓝色顺序排列。
此题中，我们使用整数 0、 1 和 2 分别表示红色、白色和蓝色。

示例 1：
输入：nums = [2,0,2,1,1,0]
输出：[0,0,1,1,2,2]

示例 2：
输入：nums = [2,0,1]
输出：[0,1,2]
 */
func sortColors(nums []int){
	p0,p2 := 0, len(nums)-1

	for i := 0; i < len(nums); i++{
		for ; i < p2 && nums[i] == 2; p2--{
			nums[i],nums[p2] = nums[p2],nums[i]
		}
		if nums[i] == 0{
			nums[i],nums[p0] = nums[p0],nums[i]
			p0++
		}
	}
}

func TestSortColors(t *testing.T){
	nums := []int{1,2,1,0}
	sortColors(nums)
	fmt.Println(nums)
}

