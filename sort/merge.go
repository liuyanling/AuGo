package sort

//import "fmt"

var tmp []int

func MergeSort(arr []int) []int {
	tmp = make([]int, len(arr))
	mergeInnerSort(arr, 0, len(arr)-1)
	return arr
}

/**
 * merge sort from down to up
 */
func MergeDToU(arr []int) []int {
	n := len(arr)
	var sz, lo int
	tmp = make([]int, n)
	for sz=1; sz < n; sz += sz {
		for lo=0; lo<n-sz; lo += sz+sz {
			merge(arr, lo, lo+sz-1, min(lo+sz+sz-1, n-1))
		}
	}
	return arr
}

func min(a, b int) int {
	if a>b {
		return b
	} else {
		return a
	}
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
	for i:=lo; i<=hi; i++ {
		tmp[i] = arr[i]
	}
	j := mid+1
	for k:=lo; k<=hi; k++ {
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