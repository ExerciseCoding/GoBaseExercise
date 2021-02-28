package base

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func returnMultiValues() (int, int) {
	return rand.Intn(10), rand.Intn(20)
}
func timeSpend(inner func(op int) int) func (op int) int {

	return func(n int) int {
		start := time.Now()
		ret := inner(n)
		fmt.Println("time spend:",time.Since(start).Seconds())
		return ret
	}
}
func slowFuc(op int) int {
	time.Sleep(time.Second*1)
	return op
}
func TestFn(t *testing.T) {
	a, _ := returnMultiValues()
	t.Log(a)
	tsSF := timeSpend(slowFuc)
	t.Log(tsSF(10))
}

//可变长参数
func Sum(ops ...int) int {
	sum := 0
	for _,v := range ops {
		sum += v
	}
	return sum
}

func TestSum(t *testing.T){
	t.Log("求和为",Sum(1,2,3))
}