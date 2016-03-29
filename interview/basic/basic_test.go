package basic

import (
  "testing"
  "fmt"
)
var arr []int

func init() {
  arr = []int{1,2,3,4,5}
}

func TestRotate(t *testing.T) {
  brr := Rotate(arr, 2)
  fmt.Println(brr)
}

func TestRotateReverse(t *testing.T) {
  arr = []int{1,2,3,4,5,6,7,8,9}
  brr := RotateWithReverse(arr, 4)
  fmt.Println(brr)
}

func TestWordLadder(t *testing.T) {
  start, end := "hit", "cog"
  dict := []string{"hot" , "dot" , "dog" , "lot" , "log"}
  len := WordLadder(start, end, dict)
  fmt.Println("len:", len)

  fmt.Println("newLen:", WordLadder("a", "c", []string{"a", "b", "c"}))
}

func TestTwoSum(t *testing.T) {
  arr := []int{1,2,3,7,8, 11, 22, 23, 24}
  s, e := TwoSum(arr, 13)
  fmt.Println("twosum:", s, e)
}

func TestPrimes(t *testing.T) {
  res := FastPrimes(1000000)
  fmt.Println("primes: ", res)
}

func TestReverseWords(t *testing.T) {
  str := "  the             sky is blue a-b "
  res := ReverseWords(str)
  fmt.Println(res)
}

func TestAtoi(t *testing.T) {
  str := "-214748365923fdsa"
  fmt.Println(Atoi(str))
}

func TestSubstrNoRepeat(t *testing.T) {
  str := "abcabcbbc"
  fmt.Println(SubstrNoRepeat(str))
  fmt.Println(SubstrNoRepeat("fdasfeghytttbnmhgsdfrczufjkt"))
  fmt.Println(SubstrNoRepeat("bbbbbb"))
}

func TestSubstrKDistinct(t *testing.T) {
  fmt.Println(SubstrKDistinct("eceba", 3))
  fmt.Println(SubstrKDistinct("ecfdsafeba", 3))
  fmt.Println(SubstrKDistinct("eecedddddehbdbdbdbdbdbdbdbdbdbdgg", 3))
}

func TestQsort(t *testing.T) {
  arr := []int{11,43,23,543,62,99,1,543,6,23,4,643,234,22,4,7,1,88, 77}
  Qsort(arr, 0, len(arr)-1)
  fmt.Println(arr)
  arr = []int{1,2,3,4,5,4,3,2,1}
  Qsort(arr, 0, 8)
  fmt.Println(arr)
}

func TestLongSubstr(t *testing.T) {
  fmt.Println(LongSubstr("ababc", "babca"))
  fmt.Println(LongSubstr("blog.iderzheng.com", "ider.cs@gmail.com"))
}

func TestLongSubSequence(t *testing.T) {
  fmt.Println(LongSubSequence("ababc", "babca"))
  fmt.Println(LongSubSequence("abdaeadba", "#a#b#a#a#b#a#"))
}

func TestLSD(t *testing.T) {
  fmt.Println(LSD([]string{"8ijkd","fdafw","fdafdw","ffafw", "234sc"}, 5))
}