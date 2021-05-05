package chan_chan

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func ChanSelectPlus(){
	initChans := [3]chan int{
		make(chan int,1),
		make(chan int, 1),
		make(chan int ,1),
	}
	//rd := rand.New(rand.NewSource(time.Now().UnixNano()))
	rand.Seed(time.Now().UnixNano())
	index := rand.Intn(3)
	fmt.Println(index)
	initChans[index] <- index
	close(initChans[2])
	select{
	case _,ok := <-initChans[0]:
		if !ok{
			fmt.Println("通道关闭了")
		}else{
			fmt.Println("this is first chan")
		}


	case _,ok := <-initChans[1]:
		if !ok{
			fmt.Println("通道关闭了")
		}else{
			fmt.Println("this is second chan")
		}


	case _,ok := <-initChans[2]:
		if !ok{
			fmt.Println("通道关闭了")
		}else{
			fmt.Println("this is third chan")
		}

		//default:
		//	fmt.Println("no chan")
	}

}

func TestChanSelectPlus(t *testing.T){
	ChanSelect()

}