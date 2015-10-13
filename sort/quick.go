package sort

func QuickSort(arr []int) []int {
	innerQuickSort(arr, 0, len(arr))
	return arr
}

func innerQuickSort(arr []int, lo, hi int) {
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
	for {
		for i<hi && k > arr[i]{
			i++
		}
		for j>lo && k < arr[j] {
			j--
		}
		if i>=j {
			break
		}
	}
	arr[lo], arr[j] = arr[j], arr[lo]
	return j
}