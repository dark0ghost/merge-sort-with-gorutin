package merge_sort

import (
	cre "merge_sort/compare"
	"testing"
)

type Int int
func (v Int) Compare(value interface{}) bool{
	return v < value.(Int)
}

func toCompareArray(array []int) (buffer []cre.Compare){
	for _, i := range array {
		buffer = append(buffer, Int(i))
	}
	return
}

func TestCompare(t *testing.T) {
	var test []cre.Compare
	test = []cre.Compare{Int(8), Int(1), Int(5), Int(7), Int(8), Int(10)}
	result := []int{1, 5, 7, 8, 8, 10}
	value, err := MergeSort(test)
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


func BenchmarkMergeSort(b *testing.B) {
	testData := toCompareArray([]int{1,3,9,6,34,21,12,3,4,5,63,23,4,234,2,121,2,3,4,6,5,990,65,652,2,3,423,4})
	for i := 0; i < b.N; i++ {
		MergeSort(testData)
	}
}

