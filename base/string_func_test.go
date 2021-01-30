package base

import (
	"strconv"
	"strings"
	"testing"
)

func TestStringFn(t  *testing.T){
	s := "A,B,C"
	parts := strings.Split(s,",")
	for _, part := range parts {
		t.Log(part)
	}
	//字符串连接
	t.Log(strings.Join(parts,"-"))

}

func TestStringConv(t *testing.T){
	//数字转成字符串
	s := strconv.Itoa(10)
	t.Log("str" + s)
	//字符串转整型
	if i,err := strconv.Atoi("10"); err == nil {
		t.Log(10 + i)
	}


}