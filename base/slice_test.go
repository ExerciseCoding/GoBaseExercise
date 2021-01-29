package base

import (
	"testing"
)
//切片
func TestSliceInit(t *testing.T){
	var s0 []int
	t.Log(len(s0),cap(s0))
	s0 = append(s0,1)
	t.Log(len(s0),cap(s0))

	s1 := []int{1,2,3,4}
	t.Log(len(s1),cap(s1))

	//使用make创建切片make(type,len,cap)
	s2 := make([]int,3,5)
	t.Log(len(s2),cap(s2))
	//t.Log(s2[0],s2[1],s2[2],s2[3],s2[4])
	s2 = append(s2,1)
	t.Log(s2[0],s2[1],s2[2],s2[3])
	t.Log(len(s2),cap(s2))
}

//测试切片可变长
func TestSliceGrowing(t *testing.T){
	s := []int{}
	for i := 0; i < 10; i++{
		//append元素时会创建新的可变长存储空间，地址发生变化因此需要重新赋值
		s = append(s,i)
		t.Log(len(s),cap(s))
	}
	/**
	运行结果:
	1 1
	2 2
	3 4
	4 4
	5 8
	6 8
	7 8
	8 8
	9 16
	10 16
	 */
}

func TestSliceShareMemory(t *testing.T){
	year := []string{"Jan","Feb","Mar","Apr","May","Jun","Jul","Aug","Sep","Oct","Nov","Dec"}
	Q2 := year[3:6]
	t.Log(Q2,len(Q2),cap(Q2))
	/**
	运行结果:
	[Apr May Jun] 3 9
	 */
	summer := year[5:8]
	t.Log(summer,len(summer),cap(summer))
	/**
	运行结果
	[Jun Jul Aug] 3 7
	 */
	summer[0] = "Unkown"
	t.Log(Q2)
	/**
	运行结果:
	[Apr May Unkown]
	 */
	t.Log(year)


}

func TestSliceComparing(t *testing.T){
	a := []int{1,2,3,4}
	b := []int{1,2,3,4}
	t.Log(a,b)
	//if a == b {
	//	t.Log("equal")
	//}
}