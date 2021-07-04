package 数组

import (
	"fmt"
	"testing"
)

/**
类别: Leetcode
题号: 283题
题目描述: 移动零
给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，
同时保持非零元素的相对顺序。给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，
同时保持非零元素的相对顺序。
示例 1：
输入： [0,1,0,3,12]
输出： [1,3,12,0,0]
*/

func moveZeros(nums []int){
	index := -1
	numsLength := len(nums)
	for i := 0; i < numsLength ; i++ {
		//fmt.Println(numsLength,"--")
		if nums[i] != 0 {
			//fmt.Println(i,"-",nums[i],"index",index)
			index++
			fmt.Println("after index++",index)
			nums[index] = nums[i]
			fmt.Println("change",nums[index])
		}
	}

	for i := index+1 ; i < len(nums); i++{
		nums[i] = 0
	}
}

func TestMoveZeros(t *testing.T){
	nums := []int{0,1,0,3,12}
	moveZeros(nums)
	for _,value := range nums{
		fmt.Println(value)
	}
}
