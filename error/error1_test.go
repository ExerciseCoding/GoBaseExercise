package error

import (
	"errors"
	"fmt"
	"testing"
)

func echo(request string)(response string,err error){
	if request == ""{
		err = errors.New("empty content")
		return
	}

	response = fmt.Sprintf("echo: %s",request)
	return
}

func TestEcho(t *testing.T){
	var(
		resp string
		err error
	)
	for _,req := range []string{"","hello"}{
		fmt.Printf("request:%s \n",req)

		if resp,err = echo(req); err != nil{
			fmt.Printf("error: %s \n",err)
			continue
		}

		fmt.Printf("response:%s \n",resp)
	}

}