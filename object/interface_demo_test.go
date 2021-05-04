package object

import (
	"fmt"
	"testing"
)

type Code string
type ProgrammerNew interface{
	WriteHelloWorld() Code
}
type GoProgrammerNew struct{

}
func (p *GoProgrammerNew) WriteHelloWorld() Code {
	return "fmt.Println(\"Go hello\")"
}
type JavaProgrammer struct {

}
func (p *JavaProgrammer) WriteHelloWorld() Code {
	return "fmt.Println(\"Java hello\")"
}

func writeFirstProgram(p ProgrammerNew){
	fmt.Printf("%T, %v\n",p,p.WriteHelloWorld())
}

func TestInterfaceDemo(t *testing.T){
	gp := new (GoProgrammerNew)
	jp := &JavaProgrammer{}
	writeFirstProgram(gp)
	writeFirstProgram(jp)
}