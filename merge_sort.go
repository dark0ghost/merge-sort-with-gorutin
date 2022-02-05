package merge_sort

import (
	"fmt"
	cre "merge_sort/compare"
	ce "merge_sort/error"
	"sync"
)

func merge(left, right  []cre.Compare) (buffer []cre.Compare) {
	var i int
	buffer = make([]cre.Compare, len(left) + len(right))
	for len(left) > 0 && len(right) > 0 {
		if left[0].Compare(right[0]) {
			buffer[i] = left[0]
			left = left[1:]
			i++
			continue
		}
			buffer[i] = right[0]
			right = right[1:]
		i++
	}
	for j := 0; j < len(left); j++ {
		buffer[i] = left[j]
		i++
	}
	for j := 0; j < len(right); j++ {
		buffer[i] = right[j]
		i++
	}
	return
}

func MergeSort(array []cre.Compare) ([]cre.Compare,error) {
	sizeArray := len(array)
	if array == nil{
		return array, ce.NilPointerError("array is nil")
	}
	if sizeArray <= 1 {
		return array, ce.SizeError(fmt.Sprintf("array size %d", sizeArray))
	}
	wg := sync.WaitGroup{}
	middle := sizeArray / 2
	var (
		left  = array[0:middle]
		right = array[middle:sizeArray]
	)
	var ln, rn []cre.Compare
	wg.Add(2)
	go func() {
		defer wg.Done()
		ln, _ = MergeSort(left)
	}()
	go func() {
		defer wg.Done()
		rn, _ = MergeSort(right)
	}()
	wg.Wait()
	return merge(ln, rn), nil

}
