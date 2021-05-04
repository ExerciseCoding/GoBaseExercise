package object

import (
	"fmt"
	"testing"
)

func DoSomething(p interface{}) {
	if i, ok := p.(int); ok {
		fmt.Println("Integer",i)
		return
	}
	if i,ok := p.(string); ok {
		fmt.Println("String",i)
		return
	}
	fmt.Println("Unkown")
}

func TestDoSomething(t *testing.T){
	DoSomething(10)
	DoSomething("aa")
}