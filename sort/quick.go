package sort

import "fmt"

func QuickSort(arr []int) []int {
	innerQuickSort(arr, 0, len(arr)-1)
	return arr
}

func innerQuickSort(arr []int, lo, hi int) {
	fmt.Println("lo, hi", lo, hi)
	if lo >= hi {
		return 
	}
	mid := innerParrtition(arr, lo, hi)
	innerQuickSort(arr, lo, mid)
	innerQuickSort(arr, mid+1, hi)
}

func innerParrtition(arr []int, lo, hi int) int {
	i, j := lo+1, hi
	k := arr[lo]
	// fmt.Println("lo, hi", lo, hi)
	for {
		for i<hi && k >= arr[i]{
			i++
		}
		for j>lo && k <= arr[j] {
			j--
		}
		if i>=j {
			break
		}
		arr[i], arr[j] = arr[j], arr[i]
	}
	arr[lo], arr[j] = arr[j], arr[lo]
	// fmt.Println("j", j)
	return j
}