package sort

import (
	"testing"
	"reflect"
	"fmt"
	"time"
	"math/rand"
	)

type OrderTestCase struct {
	input, expect []int
}

var testCases = []OrderTestCase {
	{ []int{2,1,9,11,5,6}, []int{1,2,5,6,9,11} },
	{ []int{100,9,8,6,5,4}, []int{4,5,6,8,9,100} },
	{ []int{100,9,8,8,8,4}, []int{4,8,8,8,9,100} },
	{ []int{100,90,118,6000,5,4}, []int{4,5,90,100,118,6000} },
}

func (st *OrderTestCase) RunSelectTest(t *testing.T) {
	var in = append(st.input, 0)[:len(st.input)]
	fmt.Println("input s", in)
	var result = SelectSort(in)
	if !reflect.DeepEqual(result, st.expect) {
		t.Error("Unexpected result, expected: ", st.expect, " result:", result)
	}
}

func TestSelect(t *testing.T) {
	for _, testcase := range testCases {
		testcase.RunSelectTest(t)
	}
}

func (ot *OrderTestCase) RunInsertTest(t *testing.T) {
	var in = append(ot.input, 0)[:len(ot.input)]
	fmt.Println("input i", ot.input)
	res := InsertSort(in)
	if !reflect.DeepEqual(res, ot.expect) {
		t.Error("Unexpected result, expect: ", ot.expect, " result: ", res)
	}
}

func TestInsert(t *testing.T) {
	for _, testcase := range testCases {
		testcase.RunInsertTest(t)
	}
}

func (ot *OrderTestCase) RunShellTest(t *testing.T) {
	in := append(ot.input, 0)[:len(ot.input)]
	fmt.Println("input shell: ", in)
	res := ShellSort(in)
	if !reflect.DeepEqual(res, ot.expect) {
		t.Error("Unexpected result, expect: ", ot.expect, " result: ", res)
	}
}
func TestShell(t *testing.T) {
	for _, testcase := range testCases {
		testcase.RunShellTest(t);
	}
}

func (ot *OrderTestCase) RunMergeTest(t *testing.T) {
	in := append(ot.input, 0)[:len(ot.input)]
	fmt.Println("input m", in)
	res := MergeSort(in)
	if !reflect.DeepEqual(res, ot.expect) {
		t.Error("Unexpected result, expect: ", ot.expect, " result: ", res)
	}
}

func TestMerge(t *testing.T) {
	for _, testcase := range testCases {
		testcase.RunMergeTest(t)
	}
}

func (ot *OrderTestCase) RunMergeDToUTest(t *testing.T) {
	in := append(ot.input, 0)[:len(ot.input)]
	res := MergeDToU(in)
	if !reflect.DeepEqual(res, ot.expect) {
		t.Error("Unexpected result, expect: ", ot.expect, " result: ", res)
	}
}
func TestMergeDToU(t *testing.T) {
	for _, testcase := range testCases {
		testcase.RunMergeDToUTest(t)
	}
}

func (ot *OrderTestCase) RunQuickTest(t *testing.T) {
	in := append(ot.input, 0)[:len(ot.input)]
	res := QuickSort(in)
	if !reflect.DeepEqual(res, ot.expect) {
		t.Error("Unexpected result, expect: ", ot.expect, " result: ", res)
	}
	// t.Log("pass one")
}

func TestQuickSort(t *testing.T) {
	for _, testcase := range testCases {
		testcase.RunQuickTest(t)
	}
}

func TestPressure(t *testing.T) {
//	startTime := time.Now().UnixNano()
	var input [10000000]int
	randSource := rand.NewSource(time.Now().Unix())
	r := rand.New(randSource)
	for i:=0; i<10000000; i++ {
		input[i] = r.Intn(500000)
	}
//	inputSlice := input[:]
	fmt.Println("start at ", time.Now())
//	_ = InsertSort(inputSlice)
//	endTime := time.Now().UnixNano()
//	fmt.Println("insert sort cost: ", endTime-startTime, "nano second")
//	_ = SelectSort(inputSlice)
	endTime2 := time.Now().UnixNano()
//	fmt.Println("select sort cost: ", endTime2-endTime, "nano second")
	_ = ShellSort(input[:])
	endTime3 := time.Now().UnixNano()
	fmt.Println("shell", len(input),"sort cost: ", endTime3-endTime2, "nano second")
	
	var input2 [10000000]int
	for i:=0; i<10000000; i++ {
		input2[i] = r.Intn(500000)
	}
	endTime3_ := time.Now().UnixNano()
	MergeSort(append(input2[:], 11))
	endTime4 := time.Now().UnixNano()
	fmt.Println("merge", len(input2),"sort cost: ", endTime4-endTime3_, "nano second")
	endTime4_ := time.Now().UnixNano()
	MergeDToU(append(input2[:], 12222))
	endTime5 := time.Now().UnixNano()
	fmt.Println("dtoup", len(input2),"sort cost: ", endTime5-endTime4_, "nano second")
	endTime5_ := time.Now().UnixNano()
	QuickSort(append(input2[:], 12))
	endTime6 := time.Now().UnixNano()
	fmt.Println("quick", len(input2),"sort cost: ", endTime6-endTime5_, "nano second")
	// fmt.Println("merge result part: ", mResult[1000:1020])
}

/**
 * @deprecated use reflect.DeepEqual instead
 */
func SliceEqual(a, b []int) bool {
	if a==nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}
	
	if len(a) != len(b) {
		return false
	}
	
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}