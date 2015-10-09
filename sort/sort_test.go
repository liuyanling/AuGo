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

func (st *OrderTestCase) RunTest(t *testing.T) {
	var result = SelectSort(st.input)
	if !reflect.DeepEqual(st.expect, result) {
		t.Error("Unexpected result, expected: ", st.expect, " result:", result)
	}
}

var testCases = []OrderTestCase {
	{ []int{2,1,9,11,5,6}, []int{1,2,5,6,9,11} },
	{ []int{100,9,8,6,5,4}, []int{4,5,6,8,9,100} },
	{ []int{100,9,8,8,8,4}, []int{4,8,8,8,9,100} },
	{ []int{100,90,118,6000,5,4}, []int{4,5,90,100,118,6000} },
}

func TestSelect(t *testing.T) {
	for _, testcase := range testCases {
		testcase.RunTest(t)
	}
}

func (ot *OrderTestCase) RunInsertTest(t *testing.T) {
	ot.input = InsertSort(ot.input)
	if !reflect.DeepEqual(ot.input, ot.expect) {
		t.Error("Unexpected result, expect: ", ot.expect, " result: ", ot.input)
	}
}

func TestInsert(t *testing.T) {
	for _, testcase := range testCases {
		testcase.RunInsertTest(t)
	}
}

func TestPressure(t *testing.T) {
		startTime := time.Now().UnixNano()
	var input [100000]int
	randSource := rand.NewSource(time.Now().Unix())
	r := rand.New(randSource)
	for i:=0; i<100000; i++ {
		input[i] = r.Intn(5000)
	}
	inputSlice := input[:len(input)-100]

	fmt.Println("start at ", startTime)
	_ = InsertSort(inputSlice)
	endTime := time.Now().UnixNano()
	fmt.Println("insert sort end at ", endTime, " cost: ", endTime-startTime, "nano second")
	_ = SelectSort(inputSlice)
	endTime2 := time.Now().UnixNano()
	fmt.Println("select sort end at ", endTime2, " cost: ", endTime2-endTime, "nano second")
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