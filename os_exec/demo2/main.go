package main

import (
	"context"
	"fmt"
	"os/exec"
	"time"
)
//定义结构体把执行命令的进程的结果赋值给子进程
type result struct{
	output []byte
	err error
}
/**
需求: 执行一个cmd,让它在一个协程里去执行，让它执行2秒，sleep 2;echo hello
在1秒的时候，杀死cmd
 */
func main(){
	var (
		ctx context.Context
		cancelFunc context.CancelFunc
		cmd *exec.Cmd
		resultChan chan * result
		res *result
	)
	//创建结果队列
	resultChan = make(chan *result,1000)
    //context对象有一个: chan byte ,会放在context对象里面

    //cancelFunc: close(chan byte)关闭context里面的chan
	ctx, cancelFunc = context.WithCancel(context.TODO())
	go func(){
		var(
			output []byte
			err error
		)
		cmd = exec.CommandContext(ctx,"/bin/bash", "-c", "ls -l")

		//实时监听context的chan是否关闭
		//select {case <- ctx.Done():}
		//kill pid,进程ID，杀死子进程
		//执行任务捕获输出
		output,err = cmd.CombinedOutput()

		//把任务输出结果，传给main协程
		resultChan <- &result{
			output: output,
			err:    err,
		}
	}()

	//继续往下走
	time.Sleep(1 * time.Second)
	//取消上下文
	cancelFunc()

	//在main协程里，等待子协程的退出，并打印任务执行结果
	res = <- resultChan

	//打印任务执行结果
	fmt.Println(res.err,string(res.output))

}