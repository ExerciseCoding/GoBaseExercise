package demo2

import (
	"fmt"
	"sync/atomic"
	"testing"
	"time"
)

/**
多个goroutine按照既定的顺序运行
 */
//下面方法只能让保证i按照顺序给每个groutine但是不能保证groutine按照顺序执行
func TestGroutineSort(t *testing.T){
	for i := 0; i < 10; i++{
		go func(i int) {
			fmt.Println(i)
		}(i)
	}
}

func TestGroutineSequence(t *testing.T){
	var count uint32
	trigger := func(i uint32,fn func()){
		for{
			if n := atomic.LoadUint32(&count); n == i{
				fn()
				atomic.AddUint32(&count,1)
				break
			}
			time.Sleep(time.Nanosecond)
		}
	}
	for i := uint32(0); i < 10; i++{
		go func(i uint32){
			fn := func(){
				fmt.Println(i)
			}
			trigger(i,fn)
		}(i)
	}
	trigger(10, func() {

	})
}