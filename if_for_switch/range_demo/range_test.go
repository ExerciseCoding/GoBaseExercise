package range_demo

import (
	"fmt"
	"testing"
)
/**
此处迭代的是索引值，因此i是表示从0-5
 */
/**
当只有一个迭代变量的时候，数组、数组的指针、切片和字符串的元素值都是无处安放的，我们只能拿到按照从小到大顺序给出的一个个索引值。
 */
func TestRange(t *testing.T){
	sl := []int{1,2,3,4,5,6}
	for i := range sl{
		if i == 3{
			sl[i] |= i
		}
	}
	fmt.Println(sl)
}
//结果: [1 2 3 7 5 6]

/**
1.range表达式只会在for语句开始执行时被求值一次，无论后边会有多少次迭代；
2.range表达式的求值结果会被复制，也就是说，被迭代的对象是range表达式结果值的副本而不是原值。
 */

/**
此处迭代对象sl由于range原因是range表达式的副本不是原值
 */
func TestRange2(t *testing.T){
	sl := [...]int{1,2,3,4,5,6}
	maxLen := len(sl) - 1
	for i,e := range sl{
		if i == maxLen{
			sl[0] += e
		} else{
			sl[i+1] += e
		}
		fmt.Println(i,e)
	}
	fmt.Println(sl)
}
//结果: [7 3 5 7 9 11]
/**
将方法2中的数组变成切片，由于range对象遍历的切片，而切片是引用类型，因此改变的是原值
 */
func TestRange3(t *testing.T){
	sl := []int{1,2,3,4,5,6}
	maxLen := len(sl) - 1 //5
	for i,e := range sl{ //0:1,1:2,2:3,3:4,4:5,5:6
		if i == maxLen{  //i == 5
			sl[0] += e
		} else{
			sl[i+1] += e
		}
		fmt.Println(i,e)
	}
	fmt.Println(sl)
}

//结果: [22 3 6 10 15 21]
// 0:1 1:(1+2)3 2:3 3:4 4:5 5:6
// 0:1 1:3 2:(3+3)6 3:4 4:5 5:6
// 0:1 1:3 2:6 3:(6+4)10 4:5 5:6
// 0:1 1:3 2:6 3:10 4:(10+5)15 5:6
// 0:1 1:3 2:6 3:10 4:15 5:(15+6)21
// 0:(21+1)22 1:3 2:6 3:10 4:15 5:21
// 22 3 6 10 15 21