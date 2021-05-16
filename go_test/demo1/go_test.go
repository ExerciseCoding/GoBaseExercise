package demo1

import "testing"

func TestGO(t *testing.T){
	t.Log("打印常规的测试日志")
	t.Logf("打印常规的测试日志")
	t.Fail()
}

func TestFailNow(t *testing.T){
	for i := 0; i < 10; i++{
		if i == 5{
			//FailNow 执行之后，当前函数会立即终止执行
			t.FailNow()
		}
		t.Log(i)
	}
}

func TestError(t *testing.T){
	for i := 0; i < 10; i++{
		if i == 5{
			//测试失败的同时打印失败测试日志
			t.Error("i equal 5 error")
			t.FailNow()
		}
		t.Log(i)
	}
}

func TestFatal(t *testing.T){
	for i := 0; i < 10; i++{
		if i == 5{
			//Fatal作用是打印是啊比错误日志之后立即终止当前测试函数的执行并宣告测试失败
			t.Fatal()
		}
		t.Log(i)
	}
}

