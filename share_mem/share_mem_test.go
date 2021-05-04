package share_mem

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
	"time"
)

func TestCounter(t *testing.T) {
	counter := 0
	for i := 0; i < 5000; i++ {
		go func(){
			counter++
		}()
	}
	time.Sleep(1 * time.Second)
	t.Logf("counter = %d", counter)
}


func TestCounterThreadSafe(t *testing.T) {
	var mut sync.Mutex
	counter := 0
	for i := 0; i < 5000; i++ {
		go func(){
			defer func(){
				mut.Unlock()
			}()
			mut.Lock()
			counter++
		}()
	}
	time.Sleep(1 * time.Second)
	t.Logf("counter = %d", counter)
}


func TestCounterThreadSafeWait(t *testing.T) {
	var mut sync.Mutex
	var wg sync.WaitGroup
	counter := 0
	for i := 0; i < 5000; i++ {
		wg.Add(1)
		go func(){
			defer func(){
				mut.Unlock()
			}()
			mut.Lock()
			counter++
			wg.Done()
		}()
	}
	wg.Wait()
	t.Logf("counter = %d", counter)
}


func TestRuntime(t *testing.T){
	fmt.Println(runtime.NumCPU())
}


func TestGorutine(t *testing.T){
	runtime.GOMAXPROCS(1) //指定最大p为1，从而管理协程最多的线程为1个
	wg := sync.WaitGroup{}
	wg.Add(2)
	//第一个协程
	go func(){
		fmt.Println(1)
		fmt.Println(2)
		fmt.Println(3)
		wg.Done()
	}()



	//第二个协程
	go func(){
		fmt.Println(65)
		fmt.Println(66)
		time.Sleep(1 * time.Second)
		fmt.Println(67)
		wg.Done()
	}()

	wg.Wait()

	//意味着在执行协程A的过程中，可以随时中断，去执协程行B，协程B也可能在执行过程中中断再去执行协程A。
	//看起来协程A 和 协程B 的运行像是线程的切换，但是请注意，这里的 A 和 B 都运行在同一个线程里面。它们的调度不是线程的切换，而是纯应用态的协程调度。
	//runtime.GOMAXPROCS(1)
	//time.Sleep(time.Second)
	// 如果不设置 runtime.GOMAXPROCS(1)，那么程序将会根据操作系统的 CPU 核数而启动对应数量的 P，导致多个 M，即线程的启动。那么我们程序中的协程，就会被分配到不同的线程里面去了。为了演示，故设置数量 1，使得它们都被分配到了同一个线程里面，存于线程的协程队列里面，等待被执行或调度。
}