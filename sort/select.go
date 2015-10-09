package sort

// import "fmt"

/**
 * in fact this is a bubble sort, time cost N*N/2
 * @page 156
 * @date 2015-10-09
 */
func SelectSort(arr []int) []int {
	length := len(arr)
	for i:=0; i<length; i++ {
		// v := arr[i]
		for j := i+1; j<length; j++ {
			if arr[j] < arr[i] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
	return arr
}