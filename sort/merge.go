package sort

import "fmt"

var tmp []int

func MergeSort(arr []int) []int {
	tmp = make([]int, len(arr)+1)
	copy(tmp, arr)
	mergeInnerSort(arr, 0, len(arr)-1)
	return arr
}

func mergeInnerSort(arr []int, lo, hi int) {
	if lo >= hi {
		return
	}
	if hi-lo == 1 {
		if arr[lo] > arr[hi] {
			arr[lo], arr[hi] = arr[hi], arr[lo]
		}
		return
	}
	mid := (lo + hi) / 2
	mergeInnerSort(arr, lo, mid)
	mergeInnerSort(arr, mid+1, hi)
	merge(arr, lo, mid, hi)
}

/**
 * lo<=mid<=hi
 */
func merge(arr []int, lo, mid, hi int) {
	fmt.Println("merge:", lo, mid, hi)
	j := mid+1
	for k:=0; k<=hi; k++ {
		if lo>mid {
			arr[k] = tmp[j]
			j++
		} else if j > hi {
			arr[k] = tmp[lo]
			lo++
		} else if tmp[lo] > tmp[j] {
			arr[k] = tmp[j]
			j++
		} else {
			arr[k] = tmp[lo]
			lo++
		}
	}
}