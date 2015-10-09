package sort

import (
	"fmt"
	"time"
	"math/rand"
)

func main() {
	startTime := time.Now().UnixNano()
	var input [10000]int
	randSource := rand.NewSource(time.Now().Unix())
	r := rand.New(randSource)
	for i:=0; i<10000; i++ {
		input[i] = r.Intn(500000)
	}
	inputSlice := input[:]
	fmt.Println("start at ", startTime)
	_ = InsertSort(inputSlice)
	endTime := time.Now().UnixNano()
	fmt.Println("insert sort end at ", endTime, " cost: ", endTime-startTime, "nano second")
	_ = SelectSort(inputSlice)
	endTime2 := time.Now().UnixNano()
	fmt.Println("select sort end at ", endTime2, " cost: ", endTime-endTime2, "nano second")
}