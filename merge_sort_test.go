package merge_sort

import (
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
	result := []int{1, 5, 7, 8, 8, 10}
	value, err := MergeSort(text)
	if err != nil {
		t.Fail()
	}
	for num, i := range value {
        if i != Int(result[num]){
			t.Fail()
		}
	}
	return
}


