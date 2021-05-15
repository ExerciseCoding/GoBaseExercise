package main

import "fmt"

func main(){

	fmt.Println("============")
	fmt.Println("return:",func1())
	//func1的结果
	//defer1: 1
	//defer2: 2
	//return: 2

	fmt.Println("============")
	fmt.Println("return:",func2())
	//func2
	//defer1: 1
	//defer2: 2
	//return: 0



	fmt.Println("============")
	fmt.Println("return:",func3())
	//func3
	//10
	//return: 5


	fmt.Println("============")
	fmt.Println("return:",func4())
	//func4
	//99
	//return: 19


}

func func1()(i int){
	defer func(){
		i++
		fmt.Println("defer2:",i) // 2
	}()

	//规则二: defer执行顺序为先进后出
	defer func(){
		i++
		fmt.Println("defer1:",i) // 1
	}()

	//规则三 defer可以读取有名返回值(函数指定了返回参数名)
	return 0 //直接返回值为2，怎么写都是直接return的效果
}


func func2()int{
	var i int
	defer func(){
		i++
		fmt.Println("defer2:",i)
	}()
	defer func(){
		i++
		fmt.Println("defer1:",i)
	}()

	return i
}


func func3()(r int){
	t := 5

	defer func(){
		t = t + 5
		fmt.Println(t)
	}()
	return t
}

func func4() int {
	i := 8
	//规则一 当defer被声明时，其参数就会被实时解析
	defer func(i int){
		i = 99
		fmt.Println(i)
	}(i)
	i = 19
	return i
}