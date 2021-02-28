package main

import "fmt"

func main(){
	var sli1 = []int{1,2,3,4}
	update(sli1)
	fmt.Println(sli1)
}
func update(sli []int) []int{
	sli[1] = 10
	return sli
}