package demo1

import (
	"fmt"
	"testing"
	"unsafe"
)
type Pet interface{

	Name() string
	Category() string

}

type Dog struct {
	name string
}
func (dog *Dog) SetName(name string){
	dog.name = name
}

func (dog *Dog) Name() string{
	return dog.name
}

func (dog *Dog) Category() string {
	return "eat"
}
func TestInterface(t *testing.T){
	dog := Dog{"xiaomao"}
	var pet Pet = &dog
	dog.SetName("monster")
	fmt.Println(pet.Name())

	dog1 := Dog{"little pig"}
	//dog2 := dog1
	//dog1.name = "monster"

	dog2 := &dog1
	dog2.name = "monster"

	fmt.Println(dog1.name,dog2.name)



}

func TestInterfaceNil(t *testing.T){
	var dog1 *Dog
	fmt.Println("the first dog is nil. [wrap1]")
	dog2 := dog1
	fmt.Println("the second dog is nil. [wrap1]")
	var pet Pet = dog2

	if pet == nil {
		fmt.Println("the pet is nil. [wrap1]")
	} else {
		fmt.Println("the pet is  not nil. [wrap1]",pet)
	}

}

type Animal interface{
	ScientificName() string
	Category() string
}

type Cat interface{
	Animal
	Name() string
}

func TestInterfaceCombine(t *testing.T){

}

