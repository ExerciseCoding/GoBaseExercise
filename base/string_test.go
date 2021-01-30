package base

import "testing"

//字符串
func TestString(t *testing.T){
	var s string
	t.Log(s) //初始化为默认零值
	// s[1] = '3' //string是不可变的byte slice
	s = "\xE4\xB8\xA5" //可以存储任何二进制数据
	//s = "\xE4\xBA\xB5\*FF"
	t.Log(s)
	t.Log(len(s)) //是byte数

	//字符串转化成rune的切片，取出unicode值
	c := []rune(s)
	t.Log("c=",c,"长度为",len(c))
	//t.Log("rune size:",unsafe.Sizeof(c[0))
	t.Logf("中 unicode %x",c[0])
	t.Logf("中 UTF8 %x",s)
	// s[1] = '3' str是不可变切片不能赋值
	s = "hello"
	t.Log(s)
}
//字符串遍历
func TestStringToRune(t *testing.T){
	s := "中华人民共和国"
	for _, c := range s{
		t.Logf("%[1]c %[1]x",c)
	}
}