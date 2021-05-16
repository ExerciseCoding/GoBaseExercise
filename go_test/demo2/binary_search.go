package demo2

import "fmt"

func BinarySearch( arr []int,left,right,target int)int{
	mid := (left + right) / 2 //left = 0 ,right = 5 ,mid = 2
	fmt.Println(mid,left,right,arr[mid])
	if left <= right{
		if arr[mid] == target{
			return mid
		}
		if arr[mid] > target{
			return BinarySearch(arr,left,mid-1,target)
		}

		if arr[mid] < target{
			return BinarySearch(arr,mid+1,right,target)
		}
	}

	return -1

}

