package chan_chan

import (
	"fmt"
	"math/rand"
	"testing"
)

func ChanSelect(){
	initChans := [3]chan int{
		make(chan int,1),
		make(chan int, 1),
		make(chan int ,1),
	}
	index := rand.Intn(3)
	fmt.Println(index)
	initChans[index] <- index
	close(initChans[index])
	select{
	case <-initChans[0]:
		fmt.Println("this is first chan")

	case <-initChans[1]:
		fmt.Println("this is second chan")

	case <-initChans[2]:
		fmt.Println("this is third chan")

	//default:
	//	fmt.Println("no chan")
	}

}

func TestChanSelect(t *testing.T){
	ChanSelect()

}