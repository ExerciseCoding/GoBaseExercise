package base

func subarraySum(arrs []int,k int)int{
	//统计总共有多少个子数组和为k
	count := 0
	sum := 0
	mp := map[int]int{}
	mp[0] = 1
	for i := 0; i < len(arrs); i++{
		sum += arrs[i]
		if _,ok := mp[sum-k];ok {
			count += mp[sum-k]
		}
		mp[sum] += 1
	}
	return count
}