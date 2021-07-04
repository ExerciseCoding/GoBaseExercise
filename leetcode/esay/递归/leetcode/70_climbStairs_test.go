package leetcode

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

//递归方式
func climbStairs(n int)int{
	if n == 1 {
		return 1
	}

	if n == 2 {
		return 2
	}

	return climbStairs(n-1) + climbStairs(n-2)
}
//递归方式解决重复计算问题

var mp = map[int]int{1:1,2:2}
func climbStairsNoRepeat(n int) int {
	var(
		ok bool
	)
	if n == 1 {
		return 1
	}

	if n == 2 {
		return 2
	}
	if _,ok = mp[n]; !ok{
		 ret := climbStairsFor(n-1)+climbStairsFor(n-2)
		 mp[n] = ret
	}

	return mp[n]

}
//非递归改写
func climbStairsFor(n int) int{

	if n == 1{
		return 1
	}
	if n == 2 {
		return 2
	}
	f1,f2,fe := 1,2,0
	/**
	i = 3; f1 = 2; f2 = 3
	i = 4; f1 = 3; f2 = 5
	 */
	for i := 3; i < n; i++{
		fe = f2
		f2 = f1 + f2
		f1 = fe
	}
	return f1+f2
}

func TestClimbStairs(t *testing.T){
	sumMethod := climbStairs(7)
	sumMethodFor := climbStairsFor(7)
	sumMethodNoRepeat := climbStairsNoRepeat(7)
	fmt.Println(sumMethod,sumMethodFor,sumMethodNoRepeat)

}

