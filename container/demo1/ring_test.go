package demo1

import (
	"container/ring"
	"fmt"
	"testing"
)

func TestRing(t *testing.T){
	YueSefuhuan(3,1)

}

func YueSefuhuan(n , m int) int{
	if n == 1{
		return 0
	}
	list := ring.New(n)

	for i := 1; i <= n; i++ {
		list.Value = i
		list = list.Next()
	}

	fmt.Println(list.Value)
	return 0
}