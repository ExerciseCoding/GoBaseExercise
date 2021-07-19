package sort


func QuickSort(nums []int,start, end int){

	if start < end{
		i ,j := start,end
		key := nums[(start+end)/2]
		for i <= j {
			for nums[i] < key{
				i++
			}
			for nums[i] > key{
				j--
			}

			if i <= j{
				nums[i],nums[j] = nums[j],nums[i]
				i++
				j--
			}
		}

		if start < j{
			QuickSort(nums,start,j)

		}
		if end > i {
			QuickSort(nums,i,end)
		}
	}
}



func QuickSortBackend(nums []int,left,right int){
	if left < right{
		mid := nums[(left+right)/2]

	}
}