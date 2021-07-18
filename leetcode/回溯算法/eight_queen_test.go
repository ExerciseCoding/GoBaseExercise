package 回溯算法

import (
	"testing"
)

/**
回溯算法：八皇后问题
解题思路分析: 见链接https://studygolang.com/articles/18452
*/

//定义存储皇后位置的切片

func Put(boards []int,col int){
	size := len(boards)
	//截止条件: 当放置皇后到达第8列时，皇后已经全部放置完
	if col == size{
		return
	}
	//分别从每一行第一个位置放置到最后一个，判断是否与前面几个皇后冲突
	for i := 0; i < size; i++{
		boards[col] = i //放置第n个皇后的位置
		//判断放置后该皇后是否安全
		if Safe(boards,col,i){
			Put(boards,col+1)
		}
	}

}

//判断当前放置的棋子是否安全
func Safe(board []int, col,pos int)bool{
	for i := 0; i < col; i++{
		if isAttack(board,i,col,pos){
			return false
		}
	}
	return true
}

func isAttack(board []int,c,col,pos int)bool{
	//垂直方向是否存在攻击
	if board[c] == pos{
		return true
	}
	if pos - board[c] == c - col{
		return true
	}

	if pos - board[c] == col - c{
		return true
	}
	return false
}

func Queen(size int){
	boards := make([]int,size)
	Put(boards,0)
}
func TestEightQueen(t *testing.T){
	Queen(8)//入口，开始放置第一个皇后
}