package sort

/**
 * shell sort algorithm based on insert sort method
 * 希尔排序是一种无法准确描述其性能的排序方法
 * shell sort is much faster than insert and select sort medthod
 * pressure test of then, for a 100000 random slice
 * insert sort cost:  10419917900 nano second
 * select sort cost:  10192766600 nano second
 * shell sort cost:  1001100 nano second
 */
func ShellSort(arr []int) []int {
	length := len(arr)
	// 确定h
	h := 1
	for ;h<length/3; {
		h = h*3+1
	}
	for h>=1 {
		for i:=h; i<length; i++ {
			for j:=i; j>=h && arr[j]<arr[j-h]; j -= h {
				arr[j], arr[j-h] = arr[j-h], arr[j]
			}
		}
		h /= 3
	}
	return arr
}