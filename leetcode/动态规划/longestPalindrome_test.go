package 动态规划
/**
题目：最长回文子串
 */
func LongestPalindrome(s string)string{
	if len(s) < 2{
		return s
	}
	startSub := 0
	maxLen := 0
	for i := 0; i < len(s); i++{
		left := i-1
		right := i+1
		curMaxLen := 1
		for left >= 0 && s[left] == s[i]{
			left--
			curMaxLen++
		}
		for right < len(s) && s[right] == s[i]{
			right++
			curMaxLen++
		}
		for left >= 0 && right < len(s) && s[left] == s[right]{
			left--
			right++
			curMaxLen += 2
		}
		if curMaxLen > maxLen{
			maxLen = curMaxLen
			startSub = left+1
		}
	}

	return s[startSub:startSub+maxLen]
}