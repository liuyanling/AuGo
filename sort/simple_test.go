package sort

import (
//	"fmt"
	"testing"
)

func TestMPLess(t *testing.T) {
	a := MP{100}
	b := MP{101}
	if !a.lessThan(b) {
		t.Error("wrong", a, b)
	}
	a.value = 102
	if !a.lessThan(b) {
		t.Error("wrong", a, b)
	}
}