package base

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func BinarySearch(arrs []int,left , right , key int)int{
	mid := (left + right) / 2
	if arrs[mid] == key{
		return mid
	}

	if arrs[mid] < key {
		return BinarySearch(arrs,mid+1,right,key)
	}

	if arrs[mid] > key {
		return BinarySearch(arrs,left,mid-1,key)
	}
	return -1
}


func TestBinarySearch(t *testing.T){
	arr := []int{1,2,3,4,5,7}
	index := BinarySearch(arr,0,len(arr)-1,4)
	fmt.Println(index)
}


func TestRand(t *testing.T){
	result := GenerateRandnum()
	fmt.Println(result)
}

func GenerateRandnum() int {
	rand.Seed(time.Now().Unix())
	randNum := rand.Intn(100)
	return randNum
}