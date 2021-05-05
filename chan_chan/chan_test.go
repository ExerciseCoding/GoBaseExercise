package chan_chan

import "testing"

func getIntChan() <- chan int {
	num := 5
	ch := make(chan int ,num)
	for i := 0; i < num; i++{
		//向通道里面放入数据
		ch <- i
	}
	//关闭通道
	close(ch)
	return ch
}

func TestGetIntChan(t *testing.T) {
	ch := getIntChan()
	for i := range ch {
		t.Log(i)
	}

	for i := 0 ; i < 5; i++{
		n := <- ch
		t.Log(n)
	}

}