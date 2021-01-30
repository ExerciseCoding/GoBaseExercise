package base

import "testing"

//map实现工厂模式
func TestMapWithFunValue(t *testing.T){
	m := map[int]func(op int) int{}
	m[1] = func(op int) int { return op }
	m[2] = func(op int) int {return op*op}
	m[3] = func(op int) int {return op*op*op}
	t.Log(m[1](2),m[2](2),m[3](2))

}
//map实现set
func TestMapForSet(t *testing.T){
	mySet := map[int]bool{}
	mySet[3] = true
	n := 3
	if mySet[1]{
		t.Logf("%d is existing",n)
	} else {
		t.Logf("%d is not existing", n)
	}
    //set 的长度
	t.Log(len(mySet))
	//删除元素
	delete(mySet,3)
	if mySet[1]{
		t.Logf("%d is existing",n)
	} else {
		t.Logf("%d is not existing", n)
	}
}

