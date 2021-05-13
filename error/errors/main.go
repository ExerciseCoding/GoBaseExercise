package main

import "errors"

func main(){
	Sqrt(0.8)

}

func Sqrt(f float64)(float64,error){
	err := errors.New("math square root of negative number")
	if f < 0{
		return 0,err
	}
	return 0,nil
}
