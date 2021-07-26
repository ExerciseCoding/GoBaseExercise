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
// @emthod_type: 递归求解
func findAcountDiGui(coins []int ,account,rest int)int{
	min := account
	for i := 0 ; i < len(coins); i++{
		if rest < coins[i]{
			continue
		}
		rest = rest - coins[i]
		minCount := findAcountDiGui(coins,account,rest)

		if minCount == -1{
			continue
		}
		totalCount := minCount + 1
		if totalCount < min {
			min = totalCount
		}
	}
	if min == account{
		return -1
	}
	return 0
}

// @method_type: 贪心
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
// @method_type 动态规划
// @Summary Get min coins count
// @Param curCount: 当前面币数 coins: 面币的值数组
func GetMinCoinsAmount(coins []int, total int)int{
	mem := make([]int,total+1) //构建备忘录
	mem[0] = 0 //初始化状态
	//初始化状态
	for i := 1; i < total + 1; i ++ {
		mem[i] = total + 1
	}

	for i := 1; i < total +1 ; i++ {
		for _,coin := range coins {
			if i - coin < 0 {
				continue
			}
			//做出决策
			mem[i] = Min(mem[i],mem[i-coin]) //做出决策
		}
	}
	if mem[total] == total+1 {
		return -1
	}

	return mem[total]

}

func Min(x,y int)int{
	if x > y{
		return y
	}
	return x
}
// @method_type 贪心+回溯
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