package demo1

import (
	"fmt"
	"sync/atomic"
	"testing"
	"time"
)


func TestAtomic(t *testing.T){
	//用原子加法实现原子减法
	num := uint32(18)
	fmt.Printf("The number: %d\n",num)
	delvalue := int32(-3)
	atomic.AddUint32(&num,uint32(delvalue))
	fmt.Printf("The number: %d\n",num)

	atomic.AddUint32(&num,^uint32(-(-3)-1))

	fmt.Println(^uint32(-(-3)-1))
	fmt.Printf("The number: %d\n",num)

	fmt.Printf("The two's complement of %d: %b\n",
		delvalue, uint32(delvalue)) // -3的补码。
	fmt.Printf("The equivalent: %b\n", ^uint32(-(-3)-1)) // 与-3的补码相同。
	fmt.Println()


	//forAndCAS1()
	//实现自旋锁

	fmt.Println()
	forAndCAS2()
}

func forAndCAS1(){
	ch := make(chan struct{},2)
	num := int32(0)
	fmt.Printf("%d \n",num)
	go func(){
		defer func(){
			ch <- struct{}{}
		}()

		for{
			time.Sleep(time.Millisecond*500)
			newNum := atomic.AddInt32(&num,2)
			fmt.Printf("%d \n",newNum)
			if newNum == 10{
				break
			}
		}
	}()

	go func(){
		defer func() {
			ch <- struct{}{}
		}()
		for{
			if atomic.CompareAndSwapInt32(&num,10,0){
				fmt.Printf("the number has gone to zero")
				break
			}
			time.Sleep(time.Millisecond*500)
		}
	}()
	<- ch
	<- ch
}


func forAndCAS2(){
	num := int32(0)
	ch := make(chan struct{},2)
	fmt.Printf("The number: %d \n",num)
	max := int32(10)
	go func(id int,max int32){
		defer func() {
			ch <- struct{}{}
		}()
		for i := 0;; i++{
			curNum := atomic.LoadInt32(&num)
			if curNum > max{
				break
			}

			newNum := curNum + 2
			time.Sleep(time.Millisecond * 200)
			if atomic.CompareAndSwapInt32(&num,curNum,newNum){
				fmt.Printf("The number: %d [%d-%d]\n",newNum,id,i)
			}else{
				fmt.Printf("The CAS operator failed. [%d-%d]\n",id,i)
			}

		}
	}(1,max)

	go func(id int,max int32) {
		defer func() {
			ch <- struct{}{}
		}()

		for j := 0;; j++{
			curNum := atomic.LoadInt32(&num)
			if curNum >= max {
				break
			}
			newNum := curNum + 2
			time.Sleep(time.Millisecond * 200)
			if atomic.CompareAndSwapInt32(&num,curNum,newNum){
				fmt.Printf("The number: %d-[%d-%d]\n",newNum,id,j)
			} else{
				fmt.Printf("The CAS operator failed.[%d-%d]\n",id,j)
			}
		}
	}(2,max)
	<- ch
	<- ch
}