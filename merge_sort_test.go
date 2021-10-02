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
	text = []cre.Compare{Int(8), Int(1), Int(5), Int(7)}
	err := MergeSort(&text)
	if err != nil {
		for _,i := range text {
			fmt.Println(i)
		}
		return
	}
	t.Fail()
}


