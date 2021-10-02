package merge_sort

import (
	"fmt"
	cre "merge_sort/compare"
	"testing"
)

type Int int
func (v Int) Compare(value interface{}) bool{
	return v < value.(Int)
}

func TestCompare(t *testing.T) {
	var text []cre.Compare
	text = []cre.Compare{Int(8), Int(1), Int(5), Int(7), Int(8), Int(10)}
	value, err := MergeSort(text)
	if err != nil{
		t.Fail()
	}
	for _, i := range value {
		fmt.Println(i)
	}
	return
}


