package switch_exercise

import (
	"fmt"
	"testing"
)

func TestSwitch(t *testing.T){
	value1 := [...]int{0,1,2,3,4,5,6}

	switch 1+4 {
	case value1[0],value1[1]:
		fmt.Println("0 or 1")
	case value1[2],value1[3]:
		fmt.Println("2 or 3")
	case value1[4],value1[5],value1[6]:
		fmt.Println("4 or 5 or 6")
	}

	switch value1[3] {
	case 0,1:
		fmt.Println("0 or 1")
	case 2,3:
		fmt.Println("2 or 3")
	case 4,5,6:
		fmt.Println("4 or 5 or 6")
	}



	switch value1[2] {
	case value1[0],value1[1],value1[2]:
		fmt.Println("0 or 1")
	case value1[2],value1[3]:
		fmt.Println("2 or 3")
	case value1[4],value1[5],value1[6]:
		fmt.Println("4 or 5 or 6")
	}



}