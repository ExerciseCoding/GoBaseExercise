package base

import "testing"

//map集合
/**
m := map[string]int{"one":1,"two":2,"three":3}
m1 := map[string]int
m1["one"] = 1
m2 := make(map[string]int,10 (Initial Capacity))
//为什么不初始化len?
map不能赋值为0，而slice可以赋值为默认值0
*/

func TestInitMap(t *testing.T){
	m1 := map[int]int {1:1,2:2,3:3}
	t.Log(m1[2])
	t.Logf("len m1=%d", len(m1))
	m2 := map[int]int{}
	m2[4] = 16
	t.Logf("len m2=%d",len(m2))
	m3 := make(map[int]int,10)
	t.Logf("len m3=%d",len(m3))

}
func TestAccessNotExistingKey(t *testing.T){
	m1 := map[int]int{}
	//没有赋值时，从mao中取值是0
	t.Log(m1[1])
	//赋值以后取出的值依然是零
	m1[2] = 0
	t.Log(m1[2])
	//解决方案
	//m1[3]会返回两个值，前面一个是值，后面ok是布尔值
	m1[3] = 3
	if v,ok := m1[3]; ok {
		t.Logf("key 3 is  existing value is %d.",v)
	} else {
		t.Log("key 3 is not existing.")
	}
}

//map遍历
func TestTraveMap(t *testing.T){
	m1 := map[int]int {1:1,2:2,3:3}
	for k,v := range m1 {
		t.Log(k,v)
	}
}