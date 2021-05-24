package esay

import (
	"fmt"
	"testing"
)

/**
类别: Leetcode
题号: 70题
题目描述: 爬楼梯问题
假设你正在爬楼梯。需要 n 阶你才能到达楼顶。
每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？
注意：给定 n 是一个正整数。
示例 1：
输入： 2
输出： 2
解释： 有两种方法可以爬到楼顶。
1.  1 阶 + 1 阶
2.  2 阶
*/

func climbStairs(n int)int{
	if n == 1 {
		return 1
	}

	if n == 2 {
		return 2
	}

	return climbStairs(n-1) + climbStairs(n-2)
}

func TestClimbStairs(t *testing.T){
	sumMethod := climbStairs(6)
	fmt.Println(sumMethod)
}