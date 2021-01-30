package base

import "testing"

/**
数组的声明
var name [len]type //声明并初始化默认值
b := [3]int{1,2,3} //声明同时初始化
c := [2][2]int{{1,2},{3,4}} //多维数组初始化
 */
func TestArrayInit(t *testing.T) {
	var arr [3]int
	arr1 := [4]int{1,2,3,4}
	arr2 := [...]int{1,2,3,5,6}
	arr1[1] = 5
	t.Log(arr1[1],arr[2])
	t.Log(arr1,arr2)
}

func TestArrayTravel(t *testing.T){
	arr3 := [...]int{1,3,4,5}
	//第一种遍历方式
	for i := 0; i < len(arr3); i++ {
		t.Log(arr3[i])
	}
	//第二种遍历方式
	for idx, val := range arr3 {
		t.Log(idx,val)
	}
}
//数组的截取
func TestArraySection(t *testing.T){
	arr3 := [...]int{1,2,3,4,5,67,8}
	arr3_sec := arr3[:3]
	t.Log(arr3_sec)
}

//数组可以比较，比如var arr1 [...]int(1,2,3} 和var arr2 [...]int{1,2,3}相同

func TestArrayComparing(t *testing.T){
	a := [...]int{1,2,3}
	b := [...]int{1,2,3}
	if a == b {
		t.Log("equal")
	}
	//c := [...]int{1,2,3}
	//d := [...]int{4,5,6}
	//if d > c {
	//	t.Log("equal")
	//}
	//数组不能比大小
}

