package sort

//import "fmt"
/**
 * insert sort algorithm is like play bridge
 * insert current card into sorted cards
 * insert sort can be very fast when items have been almost ordered 
 */
func InsertSort(arr []int) []int{
	length := len(arr)
	for i:=0; i<length; i++ {
		for j:=i; j>0 && arr[j-1] > arr[j]; j-- {
			arr[j-1], arr[j] = arr[j], arr[j-1]  
		}
	}
	return arr
}