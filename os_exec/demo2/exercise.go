package main

import (
	"context"
	"fmt"
	"os/exec"
	"time"
)
type Result struct {
	output []byte
	err error
}
func Excrcise(){
	resultchan := make(chan *Result,1000)
	ctx,cancel := context.WithCancel(context.Background())

	go func(){
		cmd := exec.CommandContext(ctx,"/bin/bash","-c","ls -l|wc -l")

		output,err := cmd.CombinedOutput()

		resultchan <- &Result{
			output: output,
			err:  err,
		}


	}()

	time.Sleep(1 * time.Second)
	cancel()

	res := <- resultchan
	fmt.Println("res",res)

}