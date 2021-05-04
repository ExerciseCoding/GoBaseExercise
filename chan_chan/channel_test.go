package chan_chan

import (
	"fmt"
	"testing"
	"time"
)

//向管道里面放消息
func send(c chan<- int){
	for i := 0; i< 10 ; i++ {
		fmt.Println("开始放消息",i)
		c <- i
		fmt.Println("放消息结束")
	}

}

//取消息
func recv(c <-chan int) {
	for i := range c {
		fmt.Println("取到的消息",i)
	}
}

//无缓冲channel
func TestNoBufferSendRecv(t *testing.T){
	c := make(chan int)
	go send(c)
	go recv(c)
	time.Sleep(3 * time.Second)
	//关闭channel
	close(c)
}


//有缓冲channel
func TestBufferSendRecv(t *testing.T){
	c := make(chan int,20)
	go send(c)
	go recv(c)
	time.Sleep(3 * time.Second)
	close(c)
}