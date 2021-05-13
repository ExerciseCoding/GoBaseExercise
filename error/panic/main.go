package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main(){
	Parse("")
}
func Parse(input string)(numbers []int,err error){
	defer func(){
		var(
			ok bool
		)
		if r := recover(); r != nil{
			if err, ok = r.(error); !ok {
				err = fmt.Errorf("pkg: %v",r)
			}

		}
	}()

	fields := strings.Fields(input)
	numbers = fields2numbers(fields)
	return
}

func fields2numbers(fields []string)(numbers []int){
	var (
		num int
		err error
	)
	if len(fields) == 0{
		panic("no words to parse")
	}

	for idx, field := range fields{
		if num, err = strconv.Atoi(field); err != nil{
			panic("parse error ")
		}
		numbers = append(numbers,num)
		fmt.Errorf("%s",idx)

	}
	return
}