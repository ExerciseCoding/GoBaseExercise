package demo1

import (
	"fmt"
	"testing"
	"unsafe"
)

/**
Go 语言中的哪些值是不可寻址的吗？
	常量的值。
	基本类型值的字面量。
	算术操作的结果值。
	对各种字面量的索引表达式和切片表达式的结果值。
	不过有一个例外，对切片字面量的索引结果值却是可寻址的。
	对字符串变量的索引表达式和切片表达式的结果值。
	对字典变量的索引表达式的结果值。
	函数字面量和方法字面量，以及对它们的调用表达式的结果值。
	结构体字面量的字段值，也就是对结构体字面量的选择表达式的结果值。
	类型转换表达式的结果值。
	类型断言表达式的结果值。
	接收表达式的结果值。
 */
type Named interface{
	Named() string
}

type  Cat struct {
	name string
}
func (cat *Cat) Name(){
	return
}
func TestPointer(t *testing.T){
	//1.不可变的
	const num = 123
	//_ = &num  常量不可寻址
	// _ = &(123) 基本类型值的字面量不可寻址

	var str = "abc"
	_ = str
	_ = &(str)
	// _ = &(str[0]) 对字符串变量的索引结果不可寻址
	// _ = &(str[0:2] 对字符串变量的切片结果值不可寻址

	str2 := str[0]
	_ = &(str2)

	// _ = &(123+456) 算术操作的结果值不可寻址

	num2 := 456
	_ = num2
	// _ = &(num1 + num2) //算法操作的结果值不可寻址

	// _ = &([3]int{1,2,3}[0]) 对数组字面量的索引结果值不可寻址
	// _ = &([3]int{1,2,3,4}[0:2]) 对数组字面量的切片结果值不可寻址

	_ = &([]int{1,2,3}[0]) //对切片字面量的索引结果值是可寻址的
	// _ = &([]int{1,2,3}[0:2]) //对切面字面量的切片结果值不可寻址
	// _ = &(map[int]string{1:"a"}[0]) 对字典字面量的索引结果值不可寻址

	var map1 = map[int]string{1:"a",2:"b",3:"c"}
	_ = map1
	// _ = &(map1[2]) 对字典变量的索引结果值不可寻址

	//_ = &(func(x int, y int))int{
	//	return x + y
	//}  //字面量代表的函数不可寻址

	//_ = &(func(cat Cat) Cat{
	//	return Cat{"ni"}
	//}) //不可寻址

	// _ = &(fmt.Sprintf) 标识符代表的函数不可寻址
	//_ = &(fmt.Sprintf("abc")) //对函数调用结果值不可寻址

	cat := Cat{"little cat"}
	_ = cat

	_ = &(cat.name) //标识符代表的属性值可以寻址

	// _ = &(cat.Name) 标识符代表的函数不可寻址
	// _ = &(cat.Name()) 对方法的调用结果值不可寻址

	// _ = &(Cat{"cat"}.name) 结构体字面量的字段不可寻址

	// _ = &(interface{}(cat)) 类型转换表达式的结果值不可寻址

	cat1 := interface{}(cat)
	_ = cat1

	// _ = &(cat1.(Named)) 类型断言表达式的结果只不可寻址

	named := cat1.(Named)
	_ = named

	// _ = &(named.(Cat)) 类型断言表达式结果值不可寻址

	var chan1 = make(chan int, 1)
	chan1 <- 1
	//_ = &(<-chan1) // 接收表达式的结果值不可寻址。
}

func TestMap(t *testing.T){
	value := map[int]string{1:"a",2:"b"}[2]
	fmt.Println(value)

	cat := Cat{"little cat"}
	_ = cat

	address := &(cat.name) //标识符代表的属性值可以寻址

	//address := &(Cat{"cat"}.name)
	fmt.Println(address)
}

func TestUintptr(t *testing.T){
	cat := Cat{"cat"}
	catP := &cat
	catPtr := uintptr(unsafe.Pointer(catP))

	//unsafe.Offsetof函数用于获取两个值在内存中的起始存储地址之间的偏移量，以字节为单位。
	//这两个值一个是某个字段的值，另一个是该字段值所属的那个结构体值。我们在调用这个函
	//数的时候，需要把针对字段的选择表达式传给它，比如dogP.name。
	//有了这个偏移量，又有了结构体值在内存中的起始存储地址（这里由dogPtr变量代表），
	//把它们相加我们就可以得到dogP的name字段值的起始存储地址了。这个地址由变量
	//namePtr代表。
	namePtr := catPtr + unsafe.Offsetof(cat.name)

	//再通过两次类型转换把namePtr的值转换成一个*string类型的值，这样就得到了指向dogP的name字段值的指针值。
	nameP := (*string)(unsafe.Pointer(namePtr))
	fmt.Println(nameP)
}