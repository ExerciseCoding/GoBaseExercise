package demo2

import (
	"fmt"
	"testing"
)

func TestBinarySearch(t *testing.T) {
	var arr = []int{1,2,3,4,5,6}
	result := BinarySearch(arr,0,len(arr)-1,4)
	fmt.Println(result)
}

func BenchmarkBinarySearch(b *testing.B) {
	var arr = []int{1,2,3,4,5,6}
	for i := 0; i < b.N; i++{
		BinarySearch(arr,0,len(arr)-1,4)
	}

	//BenchmarkBinarySearch-8   	1000000000	         0.000011 ns/op
	//BenchmarkBinarySearch-8:测试函数为BenchmarkBinarySearch  测试时最大P为8
	//1000000000 : 被测函数实际运行次数
	//0.000011 ns/op 表明单次执行GetPrimes函数的平均耗时0.000011纳秒 通过将最后一次执行测试函数时的执行时间，除以（被测函数的）执行次数而得出的
}
