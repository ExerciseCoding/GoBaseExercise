package error

import (
	"errors"
	"fmt"
	"testing"
)

func TestPanicVxExit(t *testing.T){
	//defer func(){
	//	fmt.Println("finally")
	//}()
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("recover from", err)
		}
	}()
	fmt.Println("start")
	panic(errors.New("something wrong"))
}