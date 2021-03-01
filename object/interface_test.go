package object

import "testing"

type Programmer interface {
	WriteHelloWorild() string
}

type GoProgrammer struct{

}

func (g *GoProgrammer) WriteHelloWorild() string {
	return "fmt.Println(\"hello world\")"
}
func TestClien(t *testing.T){
	var p Programmer
	p = new(GoProgrammer)
	t.Log(p.WriteHelloWorild())
}
