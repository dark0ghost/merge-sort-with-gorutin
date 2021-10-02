package merge_sort

import (
	cre "merge_sort/compare"
	ce "merge_sort/error"
	"sync"
)

func merge(af, ad  []cre.Compare, buffer *[]cre.Compare) {
	var i, j, k int
	saf := len(af)
	sad := len(ad)
	for ; i < saf && j < sad; k++ {
		raf := af[i]
		rad := ad[j]
		if raf.Compare(rad) && raf == rad {
			*buffer = append(*buffer, af[i])
			k++
			continue
		}
		if !raf.Compare(rad) {
			*buffer = append(*buffer, ad[i])
			k++
			continue
		}
		k++
	}
	for i < saf {
		*buffer = append(*buffer, af[i])
		i++
	}
	for j < sad {
		*buffer = append(*buffer, ad[i])
		j++
	}
}

func MergeSort(array* []cre.Compare) error {
	if array != nil {
		sizeArray := len(*array)
		if sizeArray > 1 {
			wg := sync.WaitGroup{}
			var result []cre.Compare
			hsa := sizeArray / 2
			a1 := (*array)[0:hsa]
			a2 := (*array)[hsa:sizeArray]
			go func() {
				defer wg.Done()
				err := MergeSort(&a1)
				if err != nil {

				}
			}()
			go func() {
				defer wg.Done()
				err := MergeSort(&a2)
				if err != nil {

				}
			}()
			wg.Wait()
			go merge(a1, a2, &result)
			array = &result
		}
		return ce.SizeError("array size <= 1")
	}
	return ce.NilPointerError("array is nil")
}
