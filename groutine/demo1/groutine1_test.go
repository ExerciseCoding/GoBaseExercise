package demo1

import (
	"fmt"
	"testing"
	"time"
)
/**
下面代码的输出结果是没有具体的结果的，因为主协程启动的10个子协程放入P队列的顺序未知
并且在进度G和M的调度的时候也不一定会调度所有的G,因为主协程在启动完10个子协程后
已经执行完了所有的程序并将退出，主协程如果退出其他的子协程也没有执行的意义，也就不会再执行了
 */
func TestGroutine(t *testing.T){
	for i := 0; i < 10; i++{
		go func(){
			fmt.Println(i)
		}()
	}
}

//让主goroutine等待其他goroutine执行完
//1.方式一：主协程睡一会

func TestWaitSonGroutine(t *testing.T){
	for i := 0; i < 10; i++{
		go func(){
			fmt.Println(i)
		}()
	}

	time.Sleep(2 * time.Second)
}

//2.方式2： 通道的方式

func TestWaitSonGroutine2(t *testing.T){
	ch := make(chan int,10)
	for i := 0; i < 10; i++{
		go func(){
			fmt.Println(i)
			ch <- 1
		}()
	}

	for i := 0; i < 10; i++{
		<-ch
	}

}