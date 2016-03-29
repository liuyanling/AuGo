package basic

import (
  "math"
  "strings"
  "bytes"
  "fmt"
)

func init() {
  _ = fmt.Println
}
func Strstr(hay, needle string) int {
  lenHay, lenNeddle := len(hay), len(needle)
  if lenHay == 0 || lenHay < lenNeddle {
    return -1
  }
  if lenNeddle == 0 {
    return 0
  }

  return -1
}

// FastPrimes output all prime numbers smaller than n
func FastPrimes(n int) []int {
  stub := make([]bool, n)
  for i := range stub {
    stub[i] = true
  }
  sq := int(math.Sqrt(float64(n)))
  var i = 2

  for i<sq {
    for j:=i+i; j<n; j += i {
      if stub[j] == true {
        stub[j] = false
      }
    }
    i++
    for stub[i] == false {
      i++
    }
  }

  var res []int
  for i=1; i<n; i++ {
    if stub[i] == true {
      res = append(res, i)
    }
  }

  return res
}

func ReverseWords(s string) string {
  s = strings.Trim(s, " ")
  bstr := []byte(s)
  length := len(s)
  reverseStr(bstr, 0, length-1)
  start := 0
  var res bytes.Buffer
  for i:=0; i<len(bstr); i++ {
    for i<len(bstr) && bstr[i] != ' ' {
      i++
    }
    reverseStr(bstr, start, i-1)
    j:=i
    for j < len(bstr) && bstr[j] == ' ' {
      j++
    }
    //fmt.Println(string(bstr[start:i+1]))
    res.Write(bstr[start:i+1])
    i = j
    start = i
  }
  return res.String()
}

func reverseStr(bstr []byte, start, end int) string {
  for start < end {
    bstr[start], bstr[end] = bstr[end], bstr[start]
    start++
    end--
  }

  return string(bstr)
}

func Atoi(s string) int {
  maxDiv := math.MaxInt32 / 10
  length := len(s)
  i, res := 0, 0
  neg := 1
  if s[0] == '-' {
    neg, i = -1, 1
  }
  for ; i<length && s[i] >= '0' && s[i] <= '9'; i++ {
    if res > maxDiv{
      if neg == 1 {
        return math.MaxInt32
      } else {
        return math.MinInt32
      }
    }
    if res == maxDiv && s[i] >= '8' {
      if (neg == 1) {
        return math.MaxInt32
      } else {
        return math.MinInt32
      }
    }

    res = res * 10 + int(s[i] - '0')
  }
  return neg * res
}

func SubstrNoRepeat(s string) string {
  s = strings.Trim(s, " ")
  var res string = string(s[0])
  i, j, k, lens := 0, 1, 0, len(s)

  for ;j<lens;j++ {
    for k=i; k<j; k++ {
      if s[k] == s[j] {
        if len(res) < j-i {
          res = s[i:j]
        }
        i = j
      }
    }
  }
  return res
}

// SubstrKDistinct Longest Substring with At Most K Distinct Characters
func SubstrKDistinct(s string, k int) string {
  s = strings.Trim(s, " ")
  if k <= 0 || len(s) == 0 {
    return ""
  }
  if (k > len(s)) {
    return s
  }
  var dict [256]int
  for idx := range dict {
    dict[idx] = 0
  }
  dict[s[0]] = 1
  hasChar := 1
  start, cur, length := 0, 1, len(s)
  var res string = string(s[0])

  // eecedddddehbdbdbdbdbdbdbdbdbdbdggddfgtt
  for ;cur < length; cur++ {
    ch := s[cur]
    // fmt.Println(string(ch),dict[ch], res, cur, "-", string(ch), hasChar, dict[int('a'):int('z')])

    // old method
    //if dict[ch] > 0 {
    //  dict[ch]++
    //  continue
    //}
    //if hasChar < k {
    //  dict[ch] = 1
    //  hasChar++
    //  continue
    //}

    if dict[ch] == 0 {
      hasChar++
    }
    dict[ch]++
    // if update result
    if cur-start > len(res) {
      res = s[start:cur]
    }

    for hasChar > k {
      dict[s[start]]--
      if dict[s[start]] == 0 {
        hasChar--
      }
      start++
    }
  }
  if cur-start > len(res) {
    res = s[start:cur]
  }
  return res
}

// LongPalinSubstr 13. Longest Palindromic Substring
func LongPalinSubstr(s string) string {

  return ""
}

// LongSubstr 最长公共字串
func LongSubstr(s, p string) int {
  lens, lenp := len(s), len(p)
  var mark [][]int = make([][]int, (lens+1))
  var pixels []int = make([]int, (lens+1)*(lenp+1))

  for i := range mark {
    mark[i],pixels = pixels[:lenp+1], pixels[lenp+1:]
  }

  res := 0
  ri, rj := 0, 0
  for i:=0; i<lens+1; i++ {
    //mark[i] = mark[]
    mark[i][0]= 0
  }
  for i:=0; i<lenp; i++ {
    mark[0][i] = 0
  }
  for i:=1; i<=lens; i++ {
    //sIdx := i-1
    for j :=1; j<= lenp; j++ {
      //pIdex := j-1
      if s[i-1] == p[j-1] {
        mark[i][j] = mark[i-1][j-1]+1
      } else {
        mark[i][j] = 0
      }

      if mark[i][j] > res {
        res = mark[i][j]
        ri, rj = i, j
      }
    }
  }

  fmt.Println("i, j =", ri, rj, "res:", res)
  //for i := range mark {
  //  fmt.Println(mark[i])
  //}
  for i:=0; i<res; i++ {
    fmt.Print(string(s[ri-res+i]))
  }
  fmt.Println("")
  return res
}

func LongSubSequence(s, p string) int {
  lens, lenp := len(s), len(p)
  var mark [][]int = make([][]int, (lens+1))
  var pixels []int = make([]int, (lens+1)*(lenp+1))

  for i := range mark {
    mark[i],pixels = pixels[:lenp+1], pixels[lenp+1:]
  }

  res := 0
  //ri, rj := 0, 0
  for i:=0; i<lens+1; i++ {
    //mark[i] = mark[]
    mark[i][0]= 0
  }
  for i:=0; i<lenp; i++ {
    mark[0][i] = 0
  }
  for i:=1; i<=lens; i++ {
    //sIdx := i-1
    for j :=1; j<= lenp; j++ {
      //pIdex := j-1
      if s[i-1] == p[j-1] {
        mark[i][j] = mark[i-1][j-1]+1
      } else {
        mark[i][j] = mark[i][j-1]
        if mark[i-1][j] > mark[i][j] {
          mark[i][j] = mark[i-1][j]
        }
      }

      if mark[i][j] > res {
        res = mark[i][j]
        //ri, rj = i, j
      }
    }
  }
  //fmt.Println("")
  return res
}

// LSD 低位优先字符串排序
// 字符串长度相同，字符为ascii码,k为低位排序位数
func  LSD(ss []string, k int) []string {
  lens := len(ss[0])
  lenss := len(ss)
  if k > lens {
    k = lens
  }
  ssSwap := make([]string, lenss)
  count := make([]int, 256)
  // k次按键索引排序
  for ik:=0; ik<k; ik++ {
    for i:=0; i<256; i++ {
      count[i] = 0
    }
    for i:=0; i<lenss; i++ {
      // 表示改字符占用了多少个空间，下一个字符应在这空间后索引
      // 这样，在排新数组时才是正确的
      count[ss[i][lens-ik-1]+1]++
    }

    for i:=1; i<256; i++ {
      count[i] += count[i-1];
    }

    // 按索引设置新数组
    for i:=0; i<lenss; i++ {
      ssSwap[count[ss[i][lens-ik-1]]] = ss[i]
      count[ss[i][lens-ik-1]]++
    }

    // move swap to ss
    for i:=0; i<lenss; i++ {
      ss[i] = ssSwap[i]
    }
  }
  return ss
}

