package error

import (
	"errors"
	"fmt"
	"strconv"
	"testing"
)
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
func TestGetFibonacci(t *testing.T){
	if v,err := GetFibonacci(-10); err != nil {
		//t.Error(err)
		if err == LessThanTwoError {
			fmt.Println("It is less 2")
		}
		if err == LargerThenHundredError {
			fmt.Println("It is large 100")
		}
	}else{
		t.Log(v)
	}
}


func GetFibonaccil(str string) {
	var (
		i int
		err error
		list []int
	)
	if i, err = strconv.Atoi(str); err == nil {
		if list, err = GetFibonacci(i); err == nil {
			fmt.Println(list)
		} else {
			fmt.Println("Error",err)
		}
	} else {
		fmt.Println("Error",err)
	}
}

func GetFibonacci2(str string) {
	var (
		i int
		err error
		list []int
	)
	if i, err = strconv.Atoi(str); err != nil {
		fmt.Println("Error",err)
		return
	}

	if list, err = GetFibonacci(i); err != nil {
		fmt.Println("Error",err)
		return
	}
	fmt.Println(list)
}