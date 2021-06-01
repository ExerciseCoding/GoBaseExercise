package 递归

import (
	"fmt"
	"testing"
)

/**
Leetcode: 1306题 跳跃游戏2
题目描述:
	这里有一个非负整数数组 arr，你最开始位于该数组的起始下标 start 处。当你位于下标 i 处时，你可以跳到 i + arr[i] 或者 i - arr[i]。
	请你判断自己是否能够跳到对应元素值为 0 的 任一 下标处。
	注意，不管是什么情况下，你都无法跳到数组之外。

示例:
	输入：arr = [4,2,3,0,3,1,2], start = 5
	输出：true
	解释：
	到达值为 0 的下标 3 有以下可能方案：
	下标 5 -> 下标 4 -> 下标 1 -> 下标 3
	下标 5 -> 下标 6 -> 下标 4 -> 下标 1 -> 下标 3
 */


func canReach(arr []int, start int)bool{
	book := make([]int,len(arr))
	return dfs(arr,start,book)
}

func dfs(arr []int, start int,book []int)bool{
	if start < 0 || start >= len(arr) || book[start] == 1 {
		return false
	}

	if arr[start] == 0{
		return true
	}
	book[start] = 1

	ret1 := dfs(arr,start-arr[start],book)
	ret2 := dfs(arr,start+arr[start],book)

	return ret1 || ret2
}

func TestCanReach(t *testing.T){
	arr := []int{4,2,3,0,3,1,2}
	start := 5
	result := canReach(arr,start)
	fmt.Println(result)
}