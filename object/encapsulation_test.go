package object

import (
	"fmt"
	"testing"
	"unsafe"
)

type Employee struct {
	Id string
	Name string
	Age int
}

func TestCreateEmployee(t *testing.T) {
	e := Employee{"0","Bob",23}
	e1 := Employee{Name:"Mike",Id:"2",Age:22}
	e2 := new(Employee)
	e2.Id = "3"
	t.Logf("e is %T", e)
	t.Logf("e1 is %T", e1)
	t.Logf("e2 is %T ",e2)
}
//第⼀一种定义⽅方式在实例例对应⽅方法被调⽤用时，实例例的成员会进⾏行行值复制
//func (e Employee) String() string {
//	fmt.Printf("String address is %x", unsafe.Pointer(&e.Name))
//	return fmt.Sprintf("ID:%s-Name:%s-Age:%d",e.Id,e.Name,e.Age)
//}
//通常情况下为了了避免内存拷⻉贝我们使⽤用第⼆二种定义⽅方式
func (e *Employee) StringTwo() string {
	fmt.Printf("StringTwo address is %x", unsafe.Pointer(&e.Name))
	return fmt.Sprintf("ID:%s-Name:%s-Age:%d",e.Id,e.Name,e.Age)
}
func TestString(t *testing.T){
	e := Employee{"2","zhangsan",23}
	e1 := Employee{"23","zhangsan",23}
	fmt.Printf("String address is %x \n", unsafe.Pointer(&e.Name))
	fmt.Printf("StringTwo address is %x \n", unsafe.Pointer(&e1.Name))
	//fmt.Println("String原方式",e.String())
	//fmt.Println("StringTwo指针方式",e1.String())
	fmt.Println("String原方式",e.StringTwo())
	fmt.Println("StringTwo指针方式",e1.StringTwo())
}