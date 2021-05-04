package chan_chan

import (
	"fmt"
	"testing"
	"time"
)

func OneTask() string{
	fmt.Println("start do oneTask")
	time.Sleep(time.Millisecond * 50)
	return "oneTask back"
}

func OtherTask() {
	fmt.Println("I am do other task")
	time.Sleep(time.Millisecond * 100)
	fmt.Println("OtherTask is done")
}

func TestService(t *testing.T){
	fmt.Println(OneTask())
	OtherTask()
}

func AsycOneTask() chan string {
	c := make(chan string)
	go func(){
		result := OneTask()
		c <- result
		fmt.Println("do OneTask done")
	}()
	return c
}
//无缓冲channel
func TestAsycOneTask(t *testing.T){
	var c chan string = AsycOneTask()
	OtherTask()
	fmt.Println(<-c)
}


func AsycOneTaskBuffer() chan string {
	c := make(chan string,1)
	go func(){
		result := OneTask()
		c <- result
		fmt.Println("do OneTask done")
	}()
	return c
}
//有缓冲channel

func TestAsycOneTaskBuffer(t *testing.T){
	var c chan string = AsycOneTaskBuffer()
	OtherTask()
	fmt.Println(<-c)
}

//无缓冲结果
//I am do other task
//start do oneTask
//OtherTask is done
//oneTask back
//do OneTask done

//有缓冲结果



//I am do other task
//start do oneTask
//do OneTask done
//OtherTask is done
//oneTask back