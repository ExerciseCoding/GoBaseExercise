package utils

import (
	"errors"
	"fmt"
)
func init(){
	fmt.Println("I am init")
}
func Add(a int, b int) int{
	return a + b
}
var LessThanTwoError = errors.New("n should be not less than 2")
var LargerThenHundredError = errors.New("n should be not larger than 100")
func GetFibonacci(n int)([]int , error){
	//if n < 2 || n > 100 {
	//	return nil, errors.New("n must be in [2,100]")
	//}
	if n < 2 {
		return nil, LessThanTwoError
	}
	if n > 100 {
		return nil, LargerThenHundredError
	}
	fiblist := []int{1,1}
	for i := 2; i < n; i++ {
		fiblist = append(fiblist,fiblist[i-2]+fiblist[i-1])
	}
	return fiblist, nil
}
