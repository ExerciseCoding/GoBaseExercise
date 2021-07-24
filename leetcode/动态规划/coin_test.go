package 动态规划

import (
	"fmt"
	"sort"
	"testing"
)

/**
动态规划
题目: 最少硬币数
 */

func TestGetMinCoinCountOfValue(t *testing.T){
	nums := []int{5,3}
	total := 11
	minCoin  := GetMinCoinCountLoop1(nums ,total,0)
	fmt.Println(minCoin)
}


func GetMinCoinCountLoop1(nums []int,total,k int,)int{
	count := 0
	for i := 0 ; i < k; i++{
		curCount := total / nums[i]
		rest := total - curCount * nums[i]
		count += curCount
		if rest == 0 {
			return count
		}
	}
	return -1
}



type Coins struct {
	coins []int //面值列表
	acount int //需要凑出的面硬值
	min int //最小面币数量
}
// @Summary Get min coins count
// @Param curCount: 当前面币数 rest: 剩余面币值  index: 当前面币值得个数
func (c *Coins) findAmount(curCount, rest, index int){
	if rest == 0 {
		if c.min > curCount {
			c.min = curCount
		}
		return
	}
	if index < 0 || rest < 0 {
		return
	}
	for i := rest/c.coins[index]; i >= 0 ;i-- {
		if i + curCount > c.min && c.min > 0{
			break
		}
		c.findAmount(i+curCount,rest - i * c.coins[index],index-1)
	}
	return
}
func coinChange(coins []int, acount int)int{
	// 先排序面值列表
	sort.Ints(coins)
	coin := &Coins{
		coins: coins,
		acount: acount,
		min: 0,
	}
	coin.findAmount(0,acount,len(coins)-1)
	return coin.min
}

func TestCoinChange(t *testing.T){

}