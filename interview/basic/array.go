package basic

import "fmt"

func m() {
  _ = fmt.Println
}
// Rotate rotate an array to right by k
// use slice append function
func Rotate(arr []int, k int) (brr []int) {
  len := len(arr)
  brr = append(arr[len-k:],arr[:len-k]...)
  return
}

// Rotate use reverse
func RotateWithReverse(arr []int, k int) ( []int) {
  var len = len(arr)
  reverse(arr[:len-k])
  reverse(arr)
  reverse(arr[:k])
  return arr
}

type candi struct {
  item string
  deep int
  dict []string
}
// WordLadder Q4/Q127
// @see https://leetcode.com/problems/word-ladder/
func WordLadder(start, end string, dict []string) int {
  if start == end {
    return 0
  }
  var candis []candi
  candis = append(candis, candi{item: start, deep: 0, dict: dict})
  // fmt.Println(start, end, dict, candis)
  // dictLen := len(dict)

  var item candi
  for len(candis) != 0{
    // fmt.Println("candis", candis)
    item, candis = candis[0], candis[1:]
    if wordDist(end, item.item) == 1 {
      return item.deep
    }
    dict = item.dict
    for i:=0; i < len(dict); i++{
      // fmt.Println("item:", i, item.item, dict[i], wordDist(item.item, dict[i]))
      if wordDist(item.item, dict[i]) == 1 {
        item_dict := append(dict[:i], dict[i+1:]...)
        newItem := candi{item: dict[i], deep: item.deep+1, dict: item_dict}
        candis = append(candis, newItem)
      }
    }
  }
  return -1
}

func reverse(arr []int) {
  len := len(arr)
  var i, j = 0, len-1
  for i <= j {
    arr[i], arr[j] = arr[j], arr[i]
    i++
    j--
  }
}

func wordDist(a, b string) (dist int) {
  lena, lenb := len(a), len(b)
  if lena != lenb {
    // fmt.Println("dis:len", lena, lenb)
    return -1
  }

  for i:=0; i< lena; i++ {
    if a[i] != b[i] {
      dist++
    }
  }
  // fmt.Println(a, b, "dist", dist)
  return dist
}

// TwoSum of sorted array Q2
func TwoSum(arr []int, target int) (int, int) {
  i, j := 0, len(arr)-1
  var left int
  for i<j {
    left = target-arr[i]
    for arr[j] > left {
      j--
    }
    if arr[j] == left {
      return i, j
    }
    i++
  }
  return -1,-1
}

// 11,43,23,543,62,99,1,543,6,23,4,643,234,22,4,7,1,88, 77
func Qsort(arr []int, start, end int) {
  if start >= end {
    return
  }
  if start+1 == end && arr[start] > arr[end] {
    arr[start], arr[end] = arr[end], arr[start]
    return
  }
  cur := arr[start]
  var i, j int
  for i,j =start+1, end; i<j;  {
    for i<end && arr[i] <= cur {
      i++
    }
    for j>start && arr[j] >= cur {
      j--
    }
    if i < j {
      arr[i], arr[j] = arr[j], arr[i]
    }
  }
  arr[start], arr[j] = arr[j], arr[start]
  //fmt.Println(start, end, arr[start],  j, arr)
  Qsort(arr, start, j)
  Qsort(arr, j+1, end)
}
