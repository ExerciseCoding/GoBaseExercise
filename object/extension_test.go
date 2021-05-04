package object

import (
	"fmt"
	"testing"
)

/**
第一种方式: 复合形式
 */

//type Pet struct {
//
//}
//func (p *Pet) Speak(){
//	fmt.Println("...")
//}
//
//func (p *Pet) SpeakTwo(host string){
//	fmt.Println(" ",host)
//}
//
//type Dog struct {
//	p *Pet
//}
//
//func (d *Dog) Speak(){
//	//fmt.Println("Dog...")
//	d.p.Speak()
//}
//
//func (d *Dog) SpeakTwo(host string){
//	d.p.SpeakTwo(host)
//	//fmt.Println(" Dog",host)
//}
//func TestExtendsion(t *testing.T){
//	dog := new(Dog)
//    dog.Speak()
//	dog.SpeakTwo("localhost")
//
//}

//复合第二种形式,匿名嵌套类型,（这不是继承）
type Pet struct {

}
func (p *Pet) Speak(){
	fmt.Println("...")
}

func (p *Pet) SpeakTwo(host string){
	p.Speak()
	fmt.Println(" ",host)
}

type Dog struct {
	Pet
}

//func (d *Dog) Speak(){
//	fmt.Println("Dog...")
//	//d.Speak()
//}

//func (d *Dog) SpeakTwo(host string){
//	//d.SpeakTwo(host)
//	fmt.Println(" Dog",host)
//}
func TestExtendsion(t *testing.T){
	dog := new(Dog)
	//dog.Speak()
	dog.SpeakTwo("localhost")

	//var dog Pet = new(Dog) //声明不通过
	//dog.SpeakTwo("chao")

	//强制类型转换也不行
	//var dog1 Dog = *new(Dog)
	//var p = (*Pet)(dog1)


}

//