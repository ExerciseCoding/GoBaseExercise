package main

import (
	"fmt"
	"log"
	"os/exec"
)

func main(){
	var(
		cmd *exec.Cmd
		output []byte
		err error
	)
	//生成cmd
	cmd = exec.Command("/bin/bash", "-c", "ls -l")

	//执行命令，捕获子进程的输出

	if output, err = cmd.CombinedOutput(); err != nil {
		log.Fatalf("exec command fail",err)
	}

	//打印子进程输出

	fmt.Println(string(output))

}
